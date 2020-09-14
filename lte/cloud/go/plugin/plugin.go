/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package plugin

import (
	"magma/lte/cloud/go/lte"
	"magma/lte/cloud/go/plugin/stream_provider"
	lteHandlers "magma/lte/cloud/go/services/lte/obsidian/handlers"
	lteModels "magma/lte/cloud/go/services/lte/obsidian/models"
	policydbHandlers "magma/lte/cloud/go/services/policydb/obsidian/handlers"
	policyModels "magma/lte/cloud/go/services/policydb/obsidian/models"
	"magma/lte/cloud/go/services/subscriberdb"
	subscriberdbHandlers "magma/lte/cloud/go/services/subscriberdb/obsidian/handlers"
	subscriberModels "magma/lte/cloud/go/services/subscriberdb/obsidian/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/pluginimpl/legacy_stream_providers"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/metricsd"
	"magma/orc8r/cloud/go/services/state"
	"magma/orc8r/cloud/go/services/state/indexer"
	"magma/orc8r/cloud/go/services/streamer/providers"
	"magma/orc8r/lib/go/registry"
	"magma/orc8r/lib/go/service/config"
	"magma/orc8r/lib/go/service/serviceregistry"
)

// LteOrchestratorPlugin implements OrchestratorPlugin for the LTE module
type LteOrchestratorPlugin struct{}

func (*LteOrchestratorPlugin) GetName() string {
	return lte.ModuleName
}

func (*LteOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := serviceregistry.LoadServiceRegistryConfig(lte.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

func (*LteOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		state.NewStateSerde(lte.EnodebStateType, &lteModels.EnodebState{}),
		state.NewStateSerde(lte.ICMPStateType, &subscriberModels.IcmpStatus{}),
		// MME state messages which use arbitrary untyped JSON serdes because
		// they're defined/used as protos in the MME codebase
		state.NewStateSerde(lte.MMEStateType, &state.ArbitaryJSON{}),
		state.NewStateSerde(lte.SPGWStateType, &state.ArbitaryJSON{}),
		state.NewStateSerde(lte.S1APStateType, &state.ArbitaryJSON{}),

		// Configurator serdes
		configurator.NewNetworkConfigSerde(lte.CellularNetworkType, &lteModels.NetworkCellularConfigs{}),
		configurator.NewNetworkConfigSerde(lte.NetworkSubscriberConfigType, &policyModels.NetworkSubscriberConfig{}),
		configurator.NewNetworkEntityConfigSerde(lte.CellularGatewayType, &lteModels.GatewayCellularConfigs{}),
		configurator.NewNetworkEntityConfigSerde(lte.CellularEnodebType, &lteModels.EnodebConfiguration{}),

		configurator.NewNetworkEntityConfigSerde(lte.PolicyRuleEntityType, &policyModels.PolicyRuleConfig{}),
		configurator.NewNetworkEntityConfigSerde(lte.BaseNameEntityType, &policyModels.BaseNameRecord{}),
		configurator.NewNetworkEntityConfigSerde(subscriberdb.EntityType, &subscriberModels.LteSubscription{}),

		configurator.NewNetworkEntityConfigSerde(lte.RatingGroupEntityType, &policyModels.RatingGroup{}),

		configurator.NewNetworkEntityConfigSerde(lte.ApnEntityType, &lteModels.ApnConfiguration{}),
	}
}

func (*LteOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{
		&Builder{},
	}
}

func (*LteOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return []metricsd.MetricsProfile{}
}

func (*LteOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return plugin.FlattenHandlerLists(
		lteHandlers.GetHandlers(),
		policydbHandlers.GetHandlers(),
		subscriberdbHandlers.GetHandlers(),
	)
}

func (*LteOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	factory := legacy_stream_providers.LegacyProviderFactory{}
	return []providers.StreamProvider{
		factory.CreateLegacyProvider(lte.SubscriberStreamName, &stream_provider.LteStreamProviderServicer{}),
		factory.CreateLegacyProvider(lte.PolicyStreamName, &stream_provider.LteStreamProviderServicer{}),
		factory.CreateLegacyProvider(lte.BaseNameStreamName, &stream_provider.LteStreamProviderServicer{}),
		factory.CreateLegacyProvider(lte.MappingsStreamName, &stream_provider.LteStreamProviderServicer{}),
		factory.CreateLegacyProvider(lte.NetworkWideRules, &stream_provider.LteStreamProviderServicer{}),
	}
}

func (*LteOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	return []indexer.Indexer{}
}
