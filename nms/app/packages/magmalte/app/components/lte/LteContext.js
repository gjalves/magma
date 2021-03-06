/**
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
import type {EnodebInfo} from '../lte/EnodebUtils';
import type {
  apn,
  feg_lte_network,
  feg_network,
  lte_gateway,
  lte_network,
  mutable_subscriber,
  network_id,
  network_ran_configs,
  policy_rule,
  subscriber_id,
  tier,
} from '@fbcnms/magma-api';

import * as React from 'react';
import ApnContext from '../context/ApnContext';
import EnodebContext from '../context/EnodebContext';
import GatewayContext from '../context/GatewayContext';
import GatewayTierContext from '../context/GatewayTierContext';
import InitSubscriberState from '../../state/lte/SubscriberState';
import LoadingFiller from '@fbcnms/ui/components/LoadingFiller';
import LteNetworkContext from '../context/LteNetworkContext';
import MagmaV1API from '@fbcnms/magma-api/client/WebClient';
import NetworkContext from '../../components/context/NetworkContext';
import PolicyContext from '../context/PolicyContext';
import SubscriberContext from '../context/SubscriberContext';

import {FEG_LTE} from '@fbcnms/types/network';
import {
  InitEnodeState,
  InitTierState,
  SetEnodebState,
  SetGatewayState,
  SetTierState,
  UpdateGateway,
} from '../../state/lte/EquipmentState';
import {SetApnState} from '../../state/lte/ApnState';
import {SetPolicyState} from '../../state/PolicyState';
import {UpdateNetworkState as UpdateFegLteNetworkState} from '../../state/feg_lte/NetworkState';
import {UpdateNetworkState as UpdateFegNetworkState} from '../../state/feg/NetworkState';
import {UpdateNetworkState as UpdateLteNetworkState} from '../../state/lte/NetworkState';
import {
  getSubscriberGatewayMap,
  setSubscriberState,
} from '../../state/lte/SubscriberState';
import {useContext, useEffect, useState} from 'react';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';

type Props = {
  networkId: network_id,
  children: React.Node,
};

export function GatewayContextProvider(props: Props) {
  const {networkId} = props;
  const [lteGateways, setLteGateways] = useState<{[string]: lte_gateway}>({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchState = async () => {
      try {
        const lteGateways = await MagmaV1API.getLteByNetworkIdGateways({
          networkId,
        });
        setLteGateways(lteGateways);
      } catch (e) {
        enqueueSnackbar?.('failed fetching gateway information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <GatewayContext.Provider
      value={{
        state: lteGateways,
        setState: (key, value?) => {
          return SetGatewayState({
            lteGateways,
            setLteGateways,
            networkId,
            key,
            value,
          });
        },
        updateGateway: props =>
          UpdateGateway({networkId, setLteGateways, ...props}),
      }}>
      {props.children}
    </GatewayContext.Provider>
  );
}

export function EnodebContextProvider(props: Props) {
  const {networkId} = props;
  const [enbInfo, setEnbInfo] = useState<{[string]: EnodebInfo}>({});
  const [lteRanConfigs, setLteRanConfigs] = useState<network_ran_configs>({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchState = async () => {
      try {
        if (networkId == null) {
          return;
        }
        const [lteRanConfigsResp] = await Promise.allSettled([
          MagmaV1API.getLteByNetworkIdCellularRan({networkId}),
          InitEnodeState({networkId, setEnbInfo, enqueueSnackbar}),
        ]);
        if (lteRanConfigsResp.value) {
          setLteRanConfigs(lteRanConfigsResp.value);
        }
      } catch (e) {
        enqueueSnackbar?.('failed fetching enodeb information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }
  return (
    <EnodebContext.Provider
      value={{
        state: {enbInfo},
        lteRanConfigs: lteRanConfigs,
        setState: (key, value?) =>
          SetEnodebState({enbInfo, setEnbInfo, networkId, key, value}),
        setLteRanConfigs: lteRanConfigs => setLteRanConfigs(lteRanConfigs),
      }}>
      {props.children}
    </EnodebContext.Provider>
  );
}

export function SubscriberContextProvider(props: Props) {
  const {networkId} = props;
  const [subscriberMap, setSubscriberMap] = useState({});
  const [subscriberMetrics, setSubscriberMetrics] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchLteState = async () => {
      if (networkId == null) {
        return;
      }
      await InitSubscriberState({
        networkId,
        setSubscriberMap,
        setSubscriberMetrics,
        enqueueSnackbar,
      }),
        setIsLoading(false);
    };
    fetchLteState();
  }, [networkId, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <SubscriberContext.Provider
      value={{
        state: subscriberMap,
        metrics: subscriberMetrics,
        gwSubscriberMap: getSubscriberGatewayMap(subscriberMap),
        setState: (key: subscriber_id, value?: mutable_subscriber) =>
          setSubscriberState({
            networkId,
            subscriberMap,
            setSubscriberMap,
            key,
            value,
          }),
      }}>
      {props.children}
    </SubscriberContext.Provider>
  );
}

export function GatewayTierContextProvider(props: Props) {
  const {networkId} = props;
  const [tiers, setTiers] = useState<{[string]: tier}>({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();
  const [supportedVersions, setSupportedVersions] = useState<Array<string>>([]);

  useEffect(() => {
    const fetchState = async () => {
      try {
        if (networkId == null) {
          return;
        }
        const [stableChannelResp] = await Promise.allSettled([
          MagmaV1API.getChannelsByChannelId({channelId: 'stable'}),
          InitTierState({networkId, setTiers, enqueueSnackbar}),
        ]);
        if (stableChannelResp.value) {
          setSupportedVersions(
            stableChannelResp.value.supported_versions.reverse(),
          );
        }
      } catch (e) {
        enqueueSnackbar?.('failed fetching tier information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <GatewayTierContext.Provider
      value={{
        state: {supportedVersions, tiers},
        setState: (key, value?) =>
          SetTierState({tiers, setTiers, networkId, key, value}),
      }}>
      {props.children}
    </GatewayTierContext.Provider>
  );
}

export function PolicyProvider(props: Props) {
  const {networkId} = props;
  const networkCtx = useContext(NetworkContext);
  const lteNetworkCtx = useContext(LteNetworkContext);
  const [policies, setPolicies] = useState<{[string]: policy_rule}>({});
  const [fegNetwork, setFegNetwork] = useState<feg_network>({});
  const [fegPolicies, setFegPolicies] = useState<{[string]: policy_rule}>({});
  const [isLoading, setIsLoading] = useState(true);
  const networkType = networkCtx.networkType;
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchState = async () => {
      try {
        setPolicies(
          await MagmaV1API.getNetworksByNetworkIdPoliciesRulesViewFull({
            networkId,
          }),
        );
        if (networkType === FEG_LTE) {
          const fegNetworkId = lteNetworkCtx.state?.federation.feg_network_id;
          if (fegNetworkId != null && fegNetworkId !== '') {
            setFegNetwork(
              await MagmaV1API.getFegByNetworkId({networkId: fegNetworkId}),
            );
            setFegPolicies(
              await MagmaV1API.getNetworksByNetworkIdPoliciesRulesViewFull({
                networkId: fegNetworkId,
              }),
            );
          }
        }
      } catch (e) {
        enqueueSnackbar?.('failed fetching policy information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, networkType, lteNetworkCtx, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <PolicyContext.Provider
      value={{
        state: policies,
        setState: async (key, value?, isNetworkWide?) => {
          if (networkType === FEG_LTE) {
            const fegNetworkID = lteNetworkCtx.state?.federation.feg_network_id;
            await SetPolicyState({
              policies,
              setPolicies,
              networkId,
              key,
              value,
            });

            // duplicate the policy on feg_network as well
            if (fegNetworkID != null) {
              await SetPolicyState({
                policies: fegPolicies,
                setPolicies: setFegPolicies,
                networkId: fegNetworkID,
                key,
                value,
              });
            }
          } else {
            await SetPolicyState({
              policies,
              setPolicies,
              networkId,
              key,
              value,
            });
          }
          if (isNetworkWide === true) {
            // we only support isNetworkWide rules now(and not basenames)
            let ruleNames = [];
            let fegRuleNames = [];

            if (value != null) {
              ruleNames =
                lteNetworkCtx.state?.subscriber_config
                  ?.network_wide_rule_names ?? [];
              if (!ruleNames.includes(key)) {
                ruleNames.push(key);
              }
              fegRuleNames =
                fegNetwork.subscriber_config?.network_wide_rule_names ?? [];
              if (!fegRuleNames.includes(key)) {
                fegRuleNames.push(key);
              }
            } else {
              // this is a delete operation
              const oldRuleNames =
                lteNetworkCtx.state?.subscriber_config
                  ?.network_wide_rule_names ?? [];
              const oldFegRuleNames =
                fegNetwork.subscriber_config?.network_wide_rule_names ?? [];

              ruleNames = oldRuleNames.filter(function (
                ruleId,
                _unused0,
                _unused1,
              ) {
                return ruleId !== key;
              });
              fegRuleNames = oldFegRuleNames.filter(function (
                ruleId,
                _unused0,
                _unused1,
              ) {
                return ruleId !== key;
              });
            }
            lteNetworkCtx.updateNetworks({
              networkId,
              subscriberConfig: {
                network_wide_base_names:
                  lteNetworkCtx.state?.subscriber_config
                    ?.network_wide_base_names,
                network_wide_rule_names: ruleNames,
              },
            });
            UpdateFegNetworkState({
              networkId: fegNetwork.id,
              subscriberConfig: {
                network_wide_base_names:
                  fegNetwork.subscriber_config?.network_wide_base_names,
                network_wide_rule_names: fegRuleNames,
              },
              setFegNetwork,
              refreshState: true,
            });
          } else {
            // delete network wide rules for the key
            console.log('DELETING NETWORK WIDE RULES FOR ', key);
            let ruleNames = [];
            let fegRuleNames = [];
            const oldRuleNames =
              lteNetworkCtx.state?.subscriber_config?.network_wide_rule_names ??
              [];
            const oldFegRuleNames =
              fegNetwork.subscriber_config?.network_wide_rule_names ?? [];

            ruleNames = oldRuleNames.filter(function (
              ruleId,
              _unused0,
              _unused1,
            ) {
              return ruleId !== key;
            });
            fegRuleNames = oldFegRuleNames.filter(function (
              ruleId,
              _unused0,
              _unused1,
            ) {
              return ruleId !== key;
            });
            lteNetworkCtx.updateNetworks({
              networkId,
              subscriberConfig: {
                network_wide_base_names:
                  lteNetworkCtx.state?.subscriber_config
                    ?.network_wide_base_names,
                network_wide_rule_names: ruleNames,
              },
            });
            UpdateFegNetworkState({
              networkId: fegNetwork.id,
              subscriberConfig: {
                network_wide_base_names:
                  fegNetwork.subscriber_config?.network_wide_base_names,
                network_wide_rule_names: fegRuleNames,
              },
              setFegNetwork,
              refreshState: true,
            });
          }
        },
      }}>
      {props.children}
    </PolicyContext.Provider>
  );
}

export function ApnProvider(props: Props) {
  const {networkId} = props;
  const [apns, setApns] = useState<{[string]: apn}>({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchState = async () => {
      try {
        setApns(
          await MagmaV1API.getLteByNetworkIdApns({
            networkId,
          }),
        );
      } catch (e) {
        enqueueSnackbar?.('failed fetching APN information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <ApnContext.Provider
      value={{
        state: apns,
        setState: (key, value?) => {
          return SetApnState({
            apns,
            setApns,
            networkId,
            key,
            value,
          });
        },
      }}>
      {props.children}
    </ApnContext.Provider>
  );
}

export function LteNetworkContextProvider(props: Props) {
  const {networkId} = props;
  const networkCtx = useContext(NetworkContext);
  const [lteNetwork, setLteNetwork] = useState<
    $Shape<lte_network & feg_lte_network>,
  >({});
  const [isLoading, setIsLoading] = useState(true);
  const enqueueSnackbar = useEnqueueSnackbar();

  useEffect(() => {
    const fetchState = async () => {
      try {
        if (networkCtx.networkType === FEG_LTE) {
          const [
            fegLteResp,
            fegLteSubscriberConfigResp,
          ] = await Promise.allSettled([
            MagmaV1API.getFegLteByNetworkId({networkId}),
            MagmaV1API.getFegLteByNetworkIdSubscriberConfig({networkId}),
          ]);
          if (fegLteResp.value) {
            let subscriber_config = {};
            if (fegLteSubscriberConfigResp.value) {
              subscriber_config = fegLteSubscriberConfigResp.value;
            }
            setLteNetwork({...fegLteResp.value, subscriber_config});
          }
        } else {
          setLteNetwork(await MagmaV1API.getLteByNetworkId({networkId}));
        }
      } catch (e) {
        enqueueSnackbar?.('failed fetching network information', {
          variant: 'error',
        });
      }
      setIsLoading(false);
    };
    fetchState();
  }, [networkId, networkCtx, enqueueSnackbar]);

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <LteNetworkContext.Provider
      value={{
        state: lteNetwork,
        updateNetworks: props => {
          let refreshState = true;
          if (networkId !== props.networkId) {
            refreshState = false;
          }
          if (networkCtx.networkType === FEG_LTE) {
            return UpdateFegLteNetworkState({
              networkId,
              setLteNetwork,
              refreshState,
              ...props,
            });
          } else {
            return UpdateLteNetworkState({
              networkId,
              setLteNetwork,
              refreshState,
              ...props,
            });
          }
        },
      }}>
      {props.children}
    </LteNetworkContext.Provider>
  );
}

export function LteContextProvider(props: Props) {
  const {networkId} = props;
  return (
    <LteNetworkContextProvider networkId={networkId}>
      <PolicyProvider networkId={networkId}>
        <ApnProvider networkId={networkId}>
          <SubscriberContextProvider networkId={networkId}>
            <GatewayTierContextProvider networkId={networkId}>
              <EnodebContextProvider networkId={networkId}>
                <GatewayContextProvider networkId={networkId}>
                  {props.children}
                </GatewayContextProvider>
              </EnodebContextProvider>
            </GatewayTierContextProvider>
          </SubscriberContextProvider>
        </ApnProvider>
      </PolicyProvider>
    </LteNetworkContextProvider>
  );
}
