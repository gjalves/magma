#!/usr/bin/env python3

from distutils.version import LooseVersion
from typing import Any, Dict, Optional, Tuple

from colorama import Fore
from gql.gql.graphql_client import GraphqlClient
from gql.gql.reporter import DUMMY_REPORTER, Reporter
from requests import Session
from requests.auth import HTTPBasicAuth

from .consts import (
    INVENTORY_ENDPOINT,
    INVENTORY_GRAPHQL_ENDPOINT,
    INVENTORY_STORE_DELETE_ENDPOINT,
    INVENTORY_STORE_PUT_ENDPOINT,
    LOCALHOST_INVENTORY_ENDPOINT,
    __version__,
)
from .graphql.latest_python_package_query import LatestPythonPackageQuery


class SymphonyClient(GraphqlClient):
    locationTypes: Dict[str, Any] = {}
    equipmentTypes: Dict[str, Any] = {}
    serviceTypes: Dict[str, Any] = {}
    portTypes: Dict[str, Any] = {}

    def __init__(
        self,
        email: str,
        password: str,
        tenant: str = "fb-test",
        is_local_host: bool = False,
        is_dev_mode: bool = False,
        reporter: Reporter = DUMMY_REPORTER,
    ):
        """This is the class to use for working with symphony server.

            The __init__ method uses the credentials to establish session with
            the inventory website. It also consumes graphql schema for
            validations, and validates the client version is compatible with server.

            Args:
                email (str): The email of the user to connect with.
                password (str): The password of the user to connect with.
                tenant (str, optional): The tenant to connect to -
                            should be the beginning of "{}.purpleheadband.cloud"
                            The default is "fb-test" for QA environment
                is_local_host (bool, optional): Used for developers to connect to
                            local inventory. This changes the address and also
                            disable verification of ssl certificate
                is_dev_mode (bool, optional): Used for developers to connect to
                            local inventory from a container. This changes the
                            address and also disable verification of ssl
                            certificate
                reporter (object, optional): Use reporter.InventoryReporter to
                            store reports on all successful and failed mutations
                            in inventory. The default is DummyReporter that
                            discards reports

        """

        address = (
            LOCALHOST_INVENTORY_ENDPOINT.format(tenant)
            if is_local_host
            else INVENTORY_ENDPOINT.format(tenant)
        )
        graphql_endpoint_address = address + INVENTORY_GRAPHQL_ENDPOINT

        self.session: Session = Session()
        self.session.auth = HTTPBasicAuth(email, password)
        verify_ssl = not is_local_host and not is_dev_mode
        self.session.verify = verify_ssl
        if not verify_ssl:
            import urllib3

            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

        self.put_endpoint: str = address + INVENTORY_STORE_PUT_ENDPOINT
        self.delete_endpoint: str = address + INVENTORY_STORE_DELETE_ENDPOINT

        super().__init__(graphql_endpoint_address, self.session, reporter)
        self._verify_version_is_not_broken()

    def _verify_version_is_not_broken(self):
        package = self._get_latest_python_package_version()

        latest_version, latest_breaking_version = (
            package if package is not None else (None, None)
        )

        if latest_breaking_version is not None and LooseVersion(
            latest_breaking_version
        ) > LooseVersion(__version__):
            raise Exception(
                "This version of pyinventory is not supported anymore. \
                Please download and install the latest version ({})".format(
                    latest_version
                )
            )

        if latest_version is not None and LooseVersion(latest_version) > LooseVersion(
            __version__
        ):
            print(
                str(Fore.RED)
                + "A newer version of pyinventory exists ({}). \
            It is recommended to download and install it".format(
                    latest_version
                )
            )

    def _get_latest_python_package_version(self) -> Optional[Tuple[str, str]]:

        package = LatestPythonPackageQuery.execute(self).latestPythonPackage
        if package is not None:
            return (
                package.lastPythonPackage.version,
                package.lastBreakingPythonPackage.version,
            )
        return None

    def store_file(self, file_path: str, file_type: str, is_global: bool) -> str:
        sign_response = self.session.get(
            self.put_endpoint,
            params={"contentType": file_type},
            headers={"Is-Global": str(is_global)},
        )
        sign_response = sign_response.json()
        signed_url = sign_response["URL"]
        with open(file_path, "rb") as f:
            file_data = f.read()
        response = self.session.put(
            signed_url, data=file_data, headers={"Content-Type": file_type}
        )
        response.raise_for_status()
        return sign_response["key"]

    def delete_file(self, key: str, is_global: bool) -> None:
        sign_response = self.session.delete(
            self.delete_endpoint.format(key),
            headers={"Is-Global": str(is_global)},
            allow_redirects=False,
        )
        sign_response.raise_for_status()
        assert sign_response.status_code == 307
        signed_url = sign_response.headers["location"]
        response = self.session.delete(signed_url)
        response.raise_for_status()
