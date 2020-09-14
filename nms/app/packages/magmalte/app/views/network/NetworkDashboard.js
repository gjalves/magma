/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */
import type {
  network,
  network_epc_configs,
  network_ran_configs,
} from '@fbcnms/magma-api';

import AddEditNetworkButton from './NetworkEdit';
import Button from '@material-ui/core/Button';
import CardTitleRow from '../../components/layout/CardTitleRow';
import Grid from '@material-ui/core/Grid';
import JsonEditor from '../../components/JsonEditor';
import LoadingFiller from '@fbcnms/ui/components/LoadingFiller';
import MagmaV1API from '@fbcnms/magma-api/client/WebClient';
import NetworkEpc from './NetworkEpc';
import NetworkInfo from './NetworkInfo';
import NetworkKPI from './NetworkKPIs';
import NetworkRanConfig from './NetworkRanConfig';
import React from 'react';
import TopBar from '../../components/TopBar';
import nullthrows from '@fbcnms/util/nullthrows';
import useMagmaAPI from '@fbcnms/ui/magma/useMagmaAPI';

import {NetworkCheck} from '@material-ui/icons';
import {Redirect, Route, Switch} from 'react-router-dom';
import {colors, typography} from '../../theme/default';
import {makeStyles} from '@material-ui/styles';
import {useCallback, useState} from 'react';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';
import {useRouter} from '@fbcnms/ui/hooks';

const useStyles = makeStyles(theme => ({
  dashboardRoot: {
    margin: theme.spacing(5),
  },
  appBarBtn: {
    color: colors.primary.white,
    background: colors.primary.comet,
    fontFamily: typography.button.fontFamily,
    fontWeight: typography.button.fontWeight,
    fontSize: typography.button.fontSize,
    lineHeight: typography.button.lineHeight,
    letterSpacing: typography.button.letterSpacing,

    '&:hover': {
      background: colors.primary.mirage,
    },
  },
}));

export default function NetworkDashboard() {
  const classes = useStyles();

  const {history, match, relativePath, relativeUrl} = useRouter();
  const networkId: string = nullthrows(match.params.networkId);

  const [networkInfo, setNetworkInfo] = useState<network>({});

  const {isLoading: isInfoLoading} = useMagmaAPI(
    MagmaV1API.getNetworksByNetworkId,
    {
      networkId: networkId,
    },
    useCallback(networkInfo => {
      setNetworkInfo(networkInfo);
    }, []),
  );
  if (isInfoLoading) {
    return <LoadingFiller />;
  }
  return (
    <>
      <TopBar
        header="Network"
        tabs={[
          {
            label: 'Network',
            to: '/network',
            icon: NetworkCheck,
            filters: (
              <Grid
                container
                justify="flex-end"
                alignItems="center"
                spacing={2}>
                <Grid item>
                  <AddEditNetworkButton title={'Add Network'} isLink={false} />
                </Grid>
                <Grid item>
                  <Button
                    className={classes.appBarBtn}
                    onClick={() => {
                      history.push(relativeUrl('/json'));
                    }}>
                    Edit JSON
                  </Button>
                </Grid>
              </Grid>
            ),
          },
        ]}
      />

      <Switch>
        <Route
          path={relativePath('/json')}
          render={() => (
            <NetworkJsonConfig
              network={networkInfo}
              onSave={network => setNetworkInfo(network)}
            />
          )}
        />
        <Route
          path={relativePath('/network')}
          render={() => (
            <NetworkDashboardInternal
              network={networkInfo}
              onSave={network => {
                setNetworkInfo(network);
              }}
            />
          )}
        />
        <Redirect to={relativeUrl('/network')} />
      </Switch>
    </>
  );
}

type Props = {
  network: network,
  onSave?: network => void,
};

export function NetworkJsonConfig(props: Props) {
  const {match} = useRouter();
  const [error, setError] = useState('');
  const networkId: string = nullthrows(match.params.networkId);
  const enqueueSnackbar = useEnqueueSnackbar();

  if (!(Object.keys(props.network).length > 0)) {
    return <LoadingFiller />;
  }
  return (
    <JsonEditor
      content={props.network}
      error={error}
      onSave={async networkInfo => {
        try {
          await MagmaV1API.putNetworksByNetworkId({
            networkId: networkId,
            network: networkInfo,
          });
          enqueueSnackbar('Network saved successfully', {
            variant: 'success',
          });
          setError('');
          props.onSave?.(networkInfo);
        } catch (e) {
          setError(e.response?.data?.message ?? e.message);
        }
      }}
    />
  );
}

export function NetworkDashboardInternal(props: Props) {
  const {match} = useRouter();
  const classes = useStyles();
  const networkId: string = nullthrows(match.params.networkId);
  const [epcConfigs, setEpcConfigs] = useState<network_epc_configs>({});
  const [lteRanConfigs, setLteRanConfigs] = useState<network_ran_configs>({});
  const {isLoading: isEpcLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdCellularEpc,
    {
      networkId: networkId,
    },
    useCallback(epc => setEpcConfigs(epc), []),
  );

  const {isLoading: isRanLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdCellularRan,
    {
      networkId: networkId,
    },
    useCallback(lteRanConfigs => setLteRanConfigs(lteRanConfigs), []),
  );

  const {response: lteGatwayResp, isLoading: isLteRespLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdGateways,
    {
      networkId: networkId,
    },
  );

  const {response: enb, isLoading: isEnbRespLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdEnodebs,
    {
      networkId: networkId,
    },
  );

  const {response: policyRules, isLoading: isPolicyLoading} = useMagmaAPI(
    MagmaV1API.getNetworksByNetworkIdPoliciesRules,
    {
      networkId: networkId,
    },
  );

  const {response: subscriber, isLoading: isSubscriberLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdSubscribers,
    {
      networkId: networkId,
    },
  );

  const {response: apns, isLoading: isAPNsLoading} = useMagmaAPI(
    MagmaV1API.getLteByNetworkIdApns,
    {
      networkId: networkId,
    },
  );

  if (
    isEpcLoading ||
    isRanLoading ||
    isLteRespLoading ||
    isEnbRespLoading ||
    isPolicyLoading ||
    isSubscriberLoading ||
    isAPNsLoading
  ) {
    return <LoadingFiller />;
  }
  const editProps = {
    networkInfo: props.network,
    lteRanConfigs: lteRanConfigs,
    epcConfigs: epcConfigs,
    onSaveNetworkInfo: networkInfo => {
      props.onSave?.(networkInfo);
    },
    onSaveEpcConfigs: setEpcConfigs,
    onSaveLteRanConfigs: setLteRanConfigs,
  };

  function editNetwork() {
    return (
      <AddEditNetworkButton
        title={'Edit'}
        isLink={true}
        editProps={{
          editTable: 'info',
          ...editProps,
        }}
      />
    );
  }

  function editRAN() {
    return (
      <AddEditNetworkButton
        title={'Edit'}
        isLink={true}
        editProps={{
          editTable: 'ran',
          ...editProps,
        }}
      />
    );
  }

  function editEPC() {
    return (
      <AddEditNetworkButton
        title={'Edit'}
        isLink={true}
        editProps={{
          editTable: 'epc',
          ...editProps,
        }}
      />
    );
  }
  return (
    <div className={classes.dashboardRoot}>
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <CardTitleRow label="Overview" />
          <NetworkKPI
            apns={apns}
            enb={enb}
            lteGatwayResp={lteGatwayResp}
            policyRules={policyRules}
            subscriber={subscriber}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <CardTitleRow label="Network" filter={editNetwork} />
              <NetworkInfo networkInfo={props.network} />
            </Grid>
            <Grid item xs={12}>
              <CardTitleRow label="RAN" filter={editRAN} />
              <NetworkRanConfig lteRanConfigs={lteRanConfigs} />
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={12} md={6}>
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <CardTitleRow label="EPC" filter={editEPC} />
              <NetworkEpc epcConfigs={epcConfigs} />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}
