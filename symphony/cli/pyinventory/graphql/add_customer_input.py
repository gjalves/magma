#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from dataclasses import dataclass
from datetime import datetime
from functools import partial
from gql.gql.datetime_utils import DATETIME_FIELD
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

@dataclass
class AddCustomerInput(DataClassJsonMixin):
    name: str
    externalId: Optional[str] = None

