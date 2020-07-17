/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package statemachines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"magma/fbinternal/cloud/go/services/testcontroller/obsidian/models"
	"magma/fbinternal/cloud/go/services/testcontroller/storage"
	"magma/fbinternal/cloud/go/services/testcontroller/utils"
	"magma/lte/cloud/go/lte"
	ltemodels "magma/lte/cloud/go/services/lte/obsidian/models"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/magmad"
	models2 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/cloud/go/services/state"
	"magma/orc8r/lib/go/protos"

	"github.com/golang/glog"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
)

const (
	checkForUpgradeState = "check_for_upgrade"

	verifyUpgradeStateFmt      = "verify_upgrade_%d"
	maxVerifyUpgradeStateCount = 3
	verifyUpgrade1State        = "verify_upgrade_1"
	verifyUpgrade2State        = "verify_upgrade_2"
	verifyUpgrade3State        = "verify_upgrade_3"

	maxTrafficStateCount = 3
	trafficTestStateFmt  = "traffic_test%d_%d"
	trafficTest1State1   = "traffic_test1_1"
	trafficTest1State2   = "traffic_test1_2"
	trafficTest1State3   = "traffic_test1_3"

	rebootEnodebStateFmt      = "reboot_enodeb_%d"
	maxRebootEnodebStateCount = 3
	rebootEnodeb1State        = "reboot_enodeb_1"
	rebootEnodeb2State        = "reboot_enodeb_2"
	rebootEnodeb3State        = "reboot_enodeb_3"

	verifyConnectivityState = "verify_conn"

	trafficTest2State1 = "traffic_test2_1"
	trafficTest2State2 = "traffic_test2_2"
	trafficTest2State3 = "traffic_test2_3"

	maxConfigStateCount    = 3
	reconfigEnodebStateFmt = "reconfig_enodeb%d"
	reconfigEnodebState1   = "reconfig_enodeb1"
	reconfigEnodebState2   = "reconfig_enodeb2"
	reconfigEnodebState3   = "reconfig_enodeb3"
	verifyConfig1State     = "verify_config1"

	trafficTest3State1 = "traffic_test3_1"
	trafficTest3State2 = "traffic_test3_2"
	trafficTest3State3 = "traffic_test3_3"

	restoreEnodebConfigStateFmt = "restore_enodeb%d"
	restoreEnodebConfigState1   = "restore_enodeb1"
	restoreEnodebConfigState2   = "restore_enodeb2"
	restoreEnodebConfigState3   = "restore_enodeb3"
	verifyConfig2State          = "verify_config2"

	trafficTest4State1 = "traffic_test4_1"
	trafficTest4State2 = "traffic_test4_2"
	trafficTest4State3 = "traffic_test4_3"
)

// GatewayClient defines an interface which is used to switch between
// implementations of its methods between a real implementation and a mock up for unit testing
type GatewayClient interface {
	GenerateTraffic(networdId string, gatewayId string, ssid string, pw string) (*protos.GenericCommandResponse, error)
	RebootEnodeb(networkdId string, gatewayId string, enodebSerial string) (*protos.GenericCommandResponse, error)
}

type MagmadClient struct{}

func (m *MagmadClient) GenerateTraffic(networkId string, gatewayId string, ssid string, pw string) (*protos.GenericCommandResponse, error) {
	stringVal := fmt.Sprintf("-c 'python3 /usr/local/bin/traffic_cli.py gen_traffic %s %s http://www.google.com'", ssid, pw)
	params := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"shell_params": &structpb.Value{Kind: &structpb.Value_StringValue{
				StringValue: stringVal,
			}},
		},
	}
	trafficScriptCmd := &protos.GenericCommandParams{
		Command: "bash",
		Params:  params,
	}
	resp, err := magmad.GatewayGenericCommand(networkId, gatewayId, trafficScriptCmd)
	return resp, err
}

func (m *MagmadClient) RebootEnodeb(networkId string, gatewayId string, enodebSerial string) (*protos.GenericCommandResponse, error) {
	params := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"shell_params": &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: enodebSerial}},
		},
	}
	rebootEndCmd := &protos.GenericCommandParams{
		Command: "reboot_enodeb",
		Params:  params,
	}
	resp, err := magmad.GatewayGenericCommand(networkId, gatewayId, rebootEndCmd)
	return resp, err
}

func getEnodebStatus(networkID string, enodebSN string) (*ltemodels.EnodebState, error) {
	st, err := state.GetState(networkID, lte.EnodebStateType, enodebSN)
	if err != nil {
		return nil, err
	}
	enodebState := st.ReportedState.(*ltemodels.EnodebState)
	enodebState.TimeReported = st.TimeMs
	ent, err := configurator.LoadEntityForPhysicalID(st.ReporterID, configurator.EntityLoadCriteria{})
	if err == nil {
		enodebState.ReportingGatewayID = ent.Key
	}
	return enodebState, err
}

var (
	magmaPackageVersionRegex = regexp.MustCompile(`(?:Version:\s)(.*)`)
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

func NewEnodebdE2ETestStateMachine(store storage.TestControllerStorage, client HttpClient, gatewayClient GatewayClient) TestMachine {
	return &enodebdE2ETestStateMachine{
		store: store,
		stateHandlers: map[string]handlerFunc{
			storage.CommonStartState: startNewTest,
			checkForUpgradeState:     checkForUpgrade,

			verifyUpgrade1State: makeVerifyUpgradeStateHandler(1, trafficTest1State1),
			verifyUpgrade2State: makeVerifyUpgradeStateHandler(2, trafficTest1State1),
			verifyUpgrade3State: makeVerifyUpgradeStateHandler(3, trafficTest1State1),

			trafficTest1State1: makeTrafficTestStateHandler(1, 1, gatewayClient, rebootEnodeb1State),
			trafficTest1State2: makeTrafficTestStateHandler(1, 2, gatewayClient, rebootEnodeb1State),
			trafficTest1State3: makeTrafficTestStateHandler(1, 3, gatewayClient, rebootEnodeb1State),

			rebootEnodeb1State:      makeRebootEnodebStateHandler(1, gatewayClient),
			rebootEnodeb2State:      makeRebootEnodebStateHandler(2, gatewayClient),
			rebootEnodeb3State:      makeRebootEnodebStateHandler(3, gatewayClient),
			verifyConnectivityState: makeVerifyConnectivityHandler(trafficTest2State1),

			trafficTest2State1: makeTrafficTestStateHandler(2, 1, gatewayClient, reconfigEnodebState1),
			trafficTest2State2: makeTrafficTestStateHandler(2, 2, gatewayClient, reconfigEnodebState1),
			trafficTest2State3: makeTrafficTestStateHandler(2, 3, gatewayClient, reconfigEnodebState1),

			reconfigEnodebState1: makeConfigEnodebStateHandler(1, verifyConfig1State),
			reconfigEnodebState2: makeConfigEnodebStateHandler(2, verifyConfig1State),
			reconfigEnodebState3: makeConfigEnodebStateHandler(3, verifyConfig1State),
			verifyConfig1State:   makeVerifyConfigStateHandler(trafficTest3State1),

			trafficTest3State1: makeTrafficTestStateHandler(3, 1, gatewayClient, restoreEnodebConfigState1),
			trafficTest3State2: makeTrafficTestStateHandler(3, 2, gatewayClient, restoreEnodebConfigState1),
			trafficTest3State3: makeTrafficTestStateHandler(3, 3, gatewayClient, restoreEnodebConfigState1),

			restoreEnodebConfigState1: makeConfigEnodebStateHandler(1, verifyConfig2State),
			restoreEnodebConfigState2: makeConfigEnodebStateHandler(2, verifyConfig2State),
			restoreEnodebConfigState3: makeConfigEnodebStateHandler(3, verifyConfig2State),
			verifyConfig2State:        makeVerifyConfigStateHandler(trafficTest4State1),

			trafficTest4State1: makeTrafficTestStateHandler(4, 1, gatewayClient, checkForUpgradeState),
			trafficTest4State2: makeTrafficTestStateHandler(4, 2, gatewayClient, checkForUpgradeState),
			trafficTest4State3: makeTrafficTestStateHandler(4, 3, gatewayClient, checkForUpgradeState),
		},
		client: client,
	}
}

func NewEnodebdE2ETestStateMachineNoTraffic(store storage.TestControllerStorage, client HttpClient, gatewayClient GatewayClient) TestMachine {
	return &enodebdE2ETestStateMachine{
		store: store,
		stateHandlers: map[string]handlerFunc{
			storage.CommonStartState: startNewTest,
			checkForUpgradeState:     checkForUpgrade,

			verifyUpgrade1State: makeVerifyUpgradeStateHandler(1, rebootEnodeb1State),
			verifyUpgrade2State: makeVerifyUpgradeStateHandler(2, rebootEnodeb1State),
			verifyUpgrade3State: makeVerifyUpgradeStateHandler(3, rebootEnodeb1State),

			rebootEnodeb1State:      makeRebootEnodebStateHandler(1, gatewayClient),
			rebootEnodeb2State:      makeRebootEnodebStateHandler(2, gatewayClient),
			rebootEnodeb3State:      makeRebootEnodebStateHandler(3, gatewayClient),
			verifyConnectivityState: makeVerifyConnectivityHandler(reconfigEnodebState1),

			reconfigEnodebState1: makeConfigEnodebStateHandler(1, verifyConfig1State),
			reconfigEnodebState2: makeConfigEnodebStateHandler(2, verifyConfig1State),
			reconfigEnodebState3: makeConfigEnodebStateHandler(3, verifyConfig1State),
			verifyConfig1State:   makeVerifyConfigStateHandler(restoreEnodebConfigState1),

			restoreEnodebConfigState1: makeConfigEnodebStateHandler(1, verifyConfig2State),
			restoreEnodebConfigState2: makeConfigEnodebStateHandler(2, verifyConfig2State),
			restoreEnodebConfigState3: makeConfigEnodebStateHandler(3, verifyConfig2State),
			verifyConfig2State:        makeVerifyConfigStateHandler(checkForUpgradeState),
		},
		client: client,
	}
}

type handlerFunc func(*enodebdE2ETestStateMachine, *models.EnodebdTestConfig) (string, time.Duration, error)

type enodebdE2ETestStateMachine struct {
	store         storage.TestControllerStorage
	stateHandlers map[string]handlerFunc
	client        HttpClient
}

func (e *enodebdE2ETestStateMachine) Run(state string, config interface{}, previousErr error) (string, time.Duration, error) {
	// TODO: notify slack if previousErr is non-nil?
	configCasted, ok := config.(*models.EnodebdTestConfig)
	if !ok {
		return "", 1 * time.Hour, errors.Errorf("expected config *models.EnodebdTestConfig, got %T", config)
	}
	handler, found := e.stateHandlers[state]
	if !found {
		return "", 1 * time.Hour, errors.Errorf("no handler registered for test case state %s", state)
	}
	return handler(e, configCasted)
}

// Handlers (consider making these instance methods)
// TODO: ASCII art diagram of the state machine
// TODO: enodebd reconfiguration; right now we just do gateway autoupgrades
// TODO: refactor out the AGW autoupgrade handlers if/when we make more test cases

/*
States:
	- Check for upgrade: compare repo version and gateway reported version; change tier config if different
		> Epsilon if same, 20 minutes
		> "Verify upgrade 1" if different, 10 minutes
	- Upgrade gateway: change version of target tier
		> "Verify upgrade 1", 10 minutes
	- Verify upgrade N: check gateway's reported version against tier, ping slack if max attempts reached and unsuccessful; there are 3 of these states
		> "Reboot enodeb 1" if equal, 20 minutes
		> "Verify upgrade N+1" if not equal and N < 3, 20 minutes
		> "Check for upgrade" if not equal and N >= 3 (after pinging slack), 20 minutes
	- Reboot enodeb N: reboot a gateway's enodeb, ping slack if max attempts reached and unsuccessful; there are 3 of these states
		> "Reboot Enodeb N+1" if reboot fails and N < 3, 15 minutes
		> "Check for upgrade" if N >=3, 15 minutes (ping slack)
		> "Verify Connectivity", 15 minutes
	- Verify Connectivity: after reboot, we check that the enodeb has successfully reconnected, then ping slack whether successful or unsuccessful
		> "Check for upgrade" if cannot get enodeb status or hwID, 5 minutes
		> "Check for upgrade", 15 minutes
*/

func startNewTest(*enodebdE2ETestStateMachine, *models.EnodebdTestConfig) (string, time.Duration, error) {
	return checkForUpgradeState, time.Minute, nil
}

func makeConfigEnodebStateHandler(stateNumber int, successState string) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return configEnodeb(stateNumber, successState, machine, config)
	}
}

func configEnodeb(stateNumber int, successState string, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	return successState, 10 * time.Minute, nil
}

func makeVerifyConfigStateHandler(successState string) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return verifyConfig(successState, machine, config)
	}
}

func verifyConfig(successState string, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	return successState, 10 * time.Minute, nil
}

func makeTrafficTestStateHandler(trafficTestNumber int, stateNumber int, gatewayClient GatewayClient, successState string) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return trafficTest(trafficTestNumber, stateNumber, gatewayClient, successState, machine, config)
	}
}

func trafficTest(trafficTestNumber int, stateNumber int, gatewayClient GatewayClient, successState string, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	trafficGWID := *config.TrafficGwID
	pretext := fmt.Sprintf(trafficPretextFmt, trafficTestNumber, *config.EnodebSN, *config.AgwConfig.TargetGatewayID, "SUCCEEDED")
	fallback := "Generate traffic notification"

	resp, err := gatewayClient.GenerateTraffic(*config.NetworkID, trafficGWID, config.Ssid, config.SsidPw)
	if resp == nil || err != nil {
		if stateNumber >= maxTrafficStateCount {
			pretext = fmt.Sprintf(trafficPretextFmt, trafficTestNumber, *config.EnodebSN, *config.AgwConfig.TargetGatewayID, "FAILED")
			postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, false, pretext, fallback, "", "")
			return checkForUpgradeState, 1 * time.Minute, errors.Errorf("Traffic test number %d failed on gwID %s after %d tries", trafficTestNumber, trafficGWID, maxTrafficStateCount)
		}
		return fmt.Sprintf(trafficTestStateFmt, trafficTestNumber, stateNumber+1), 1 * time.Minute, err
	}
	postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, true, pretext, fallback, "", "")
	return successState, 1 * time.Minute, nil
}

func makeRebootEnodebStateHandler(stateNumber int, gatewayClient GatewayClient) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return rebootEnodebStateHandler(stateNumber, gatewayClient, machine, config)
	}
}

func rebootEnodebStateHandler(stateNumber int, gatewayClient GatewayClient, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	targetGWID := *config.AgwConfig.TargetGatewayID
	enodebSN := *config.EnodebSN

	resp, err := gatewayClient.RebootEnodeb(*config.NetworkID, targetGWID, enodebSN)
	if resp == nil || err != nil {
		if stateNumber >= maxRebootEnodebStateCount {
			pretext := fmt.Sprintf(rebootPretextFmt, enodebSN, targetGWID, "FAILED")
			fallback := "Reboot enodeb notification"
			postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, false, pretext, fallback, "", "")
			return checkForUpgradeState, 15 * time.Minute, errors.Errorf("enodeb %s did not reboot within %d tries", enodebSN, maxRebootEnodebStateCount)
		}
		return fmt.Sprintf(rebootEnodebStateFmt, stateNumber+1), 5 * time.Minute, err
	}
	return verifyConnectivityState, 15 * time.Minute, nil
}

func makeVerifyConnectivityHandler(successState string) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return verifyConnectivity(successState, machine, config)
	}
}

func verifyConnectivity(successState string, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	targetGWID := *config.AgwConfig.TargetGatewayID
	enodebSN := *config.EnodebSN
	pretext := fmt.Sprintf(rebootPretextFmt, enodebSN, targetGWID, "FAILED")
	fallback := "Reboot enodeb notification"

	resp, err := getEnodebStatus(*config.NetworkID, enodebSN)
	if resp == nil || err != nil {
		postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, false, pretext, fallback, "", "")
		return checkForUpgradeState, 5 * time.Minute, errors.Wrap(err, "error getting enodeb status")
	}

	pretext = fmt.Sprintf(rebootPretextFmt, enodebSN, targetGWID, "SUCCEEDED")
	postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, true, pretext, fallback, "", "")
	return successState, 15 * time.Minute, nil
}

func checkForUpgrade(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	repoVersion, err := getLatestRepoMagmaVersion(machine.client, *config.AgwConfig.PackageRepo, *config.AgwConfig.ReleaseChannel)
	if err != nil {
		return checkForUpgradeState, 20 * time.Minute, errors.Wrap(err, "error getting latest package version from repo")
	}

	tierCfg, err := getTargetTierConfig(config)
	if err != nil {
		return checkForUpgradeState, 20 * time.Minute, err
	}
	existingVersion := tierCfg.Version.ToString()
	if existingVersion == "" {
		existingVersion = "0.0.0"
	}

	newer, err := utils.IsNewerVersion(existingVersion, repoVersion)
	if err != nil {
		return checkForUpgradeState, 20 * time.Minute, errors.Wrapf(err, "bad versions encountered: %s, %s", repoVersion, existingVersion)
	}
	if !newer {
		return checkForUpgradeState, 20 * time.Minute, nil
	}

	// Update the tier config
	newTierCfg := tierCfg
	newTierCfg.Version = models2.TierVersion(repoVersion)
	_, err = configurator.UpdateEntity(*config.NetworkID, configurator.EntityUpdateCriteria{
		Key:       *config.AgwConfig.TargetTier,
		Type:      orc8r.UpgradeTierEntityType,
		NewConfig: newTierCfg,
	})
	if err != nil {
		return checkForUpgradeState, 20 * time.Minute, errors.Wrap(err, "error updating target tier")
	}
	return verifyUpgrade1State, 10 * time.Minute, nil
}

func makeVerifyUpgradeStateHandler(stateNumber int, successState string) handlerFunc {
	return func(machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
		return verifyUpgrade(stateNumber, successState, machine, config)
	}
}

func verifyUpgrade(stateNumber int, successState string, machine *enodebdE2ETestStateMachine, config *models.EnodebdTestConfig) (string, time.Duration, error) {
	targetGWID := *config.AgwConfig.TargetGatewayID
	fallback := "Gateway auto-upgrade notification"

	// Load target version
	tierCfg, err := getTargetTierConfig(config)
	if err != nil {
		return fmt.Sprintf(verifyUpgradeStateFmt, stateNumber+1), 10 * time.Minute, err
	}
	currentVersion, err := getCurrentAGWPackageVersion(config)
	if err != nil {
		return fmt.Sprintf(verifyUpgradeStateFmt, stateNumber+1), 10 * time.Minute, err
	}

	// If equal, transition to reboot enodeb state
	if string(tierCfg.Version) == currentVersion {
		pretext := fmt.Sprintf(autoupgradePretextFmt, targetGWID, "SUCCEEDED", "")
		postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, true, pretext, fallback, string(tierCfg.Version), "")
		return successState, 20 * time.Minute, nil
	}

	if stateNumber >= maxVerifyUpgradeStateCount {
		pretext := fmt.Sprintf(autoupgradePretextFmt, targetGWID, "FAILED", "")
		postToSlack(machine.client, *config.AgwConfig.SLACKWebhook, false, pretext, fallback, string(tierCfg.Version), "")
		return checkForUpgradeState, 20 * time.Minute, errors.Errorf("gateway %s did not upgrade within %d tries", targetGWID, maxVerifyUpgradeStateCount)
	} else {
		return fmt.Sprintf(verifyUpgradeStateFmt, stateNumber+1), 10 * time.Minute, nil
	}
}

// State handler helpers

func getLatestRepoMagmaVersion(client HttpClient, url string, releaseChannel string) (string, error) {
	url = fmt.Sprintf("%s/dists/%s/main/binary-amd64/Packages", url, releaseChannel)
	resp, err := client.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "unable to retrieve packages from url %s", url)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "unable to read body from http get response at url %s", url)
	}

	if len(responseData) == 0 {
		return "", errors.Errorf("no packages found at url %s", url)
	}

	packages := strings.Split(string(responseData), "\n\n")

	latest := "0.0.0-0-0"
	for _, pkg := range packages {
		if !strings.Contains(pkg, "Package: magma\n") {
			continue
		}

		version := magmaPackageVersionRegex.FindStringSubmatch(pkg)
		if len(version) != 2 {
			glog.Warningf("incorrect regex match on package version. "+
				"should have one capturing and one non-capturing group: %s", version)
			continue
		}
		newer, err := utils.IsNewerVersion(latest, version[1])
		if err != nil {
			return "", errors.Wrap(err, "unable to compare package versions")
		}
		if newer {
			latest = version[1]
		}
	}
	if latest == "0.0.0-0-0" {
		return "", errors.Errorf("no latest magma version found")
	}
	return latest, nil
}

func getTargetTierConfig(config *models.EnodebdTestConfig) (*models2.Tier, error) {
	tierEnt, err := configurator.LoadEntity(*config.NetworkID, orc8r.UpgradeTierEntityType, *config.AgwConfig.TargetTier, configurator.EntityLoadCriteria{LoadConfig: true})
	if err != nil {
		return nil, errors.Wrap(err, "failed to load target upgrade tier")
	}

	tierCfg, ok := tierEnt.Config.(*models2.Tier)
	if !ok {
		return nil, errors.Wrapf(err, "expected tier of type *models.Tier, got %T", tierEnt.Config)
	}
	return tierCfg, nil
}

func getCurrentAGWPackageVersion(config *models.EnodebdTestConfig) (string, error) {
	targetGWID := *config.AgwConfig.TargetGatewayID

	hwID, err := configurator.GetPhysicalIDOfEntity(*config.NetworkID, orc8r.MagmadGatewayType, *config.AgwConfig.TargetGatewayID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to load hwID for target gateway %s", targetGWID)
	}
	agwState, err := state.GetGatewayStatus(*config.NetworkID, hwID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to load gateway status for %s", targetGWID)
	}
	if agwState == nil || agwState.PlatformInfo == nil {
		return "", errors.Wrapf(err, "gateway status not fully reported for %s", targetGWID)
	}
	magmaPackage := funk.Find(agwState.PlatformInfo.Packages, func(p *models2.Package) bool { return p.Name == "magma" })
	if magmaPackage == nil {
		return "", errors.Errorf("no magma package version reported for %s", targetGWID)
	}
	return magmaPackage.(*models2.Package).Version, nil
}

// swallow errors
func postToSlack(client HttpClient, slackURL string, success bool, pretext string, fallback string, targetVersion string, extraErrorText string) {
	payload, err := getSlackPayload(success, pretext, fallback, targetVersion, extraErrorText)
	if err != nil {
		glog.Errorf("failed to construct slack payload: %s", err)
		return
	}
	postPayload(client, slackURL, payload)
}

func postPayload(client HttpClient, slackURL string, payload io.Reader) {
	resp, err := client.Post(slackURL, "application/json", payload)
	if err != nil {
		glog.Errorf("slack webhook post failure: %s", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Errorf("failed to read non-200 response body: %s", err)
		}
		glog.Errorf("non-200 response %d from slack: %s", resp.StatusCode, string(respBody))
	}
}

const (
	colorRed              = "#8b0902"
	colorGreen            = "#36a64f"
	autoupgradePretextFmt = "Auto-upgrade of gateway %s %s. %s"
	rebootPretextFmt      = "Enodeb reboot of enodeb %s of gateway %s %s"
	trafficPretextFmt     = "Generate traffic test %d for enodeb %s of gateway %s %s"
)

type slackPayload struct {
	Attachments []slackAttachment `json:"attachments"`
}

type slackAttachment struct {
	Color    string                 `json:"color"`
	Pretext  string                 `json:"pretext"`
	Fallback string                 `json:"fallback"`
	Fields   []slackAttachmentField `json:"fields"`
}

type slackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func getSlackPayload(success bool, pretext string, fallback string, targetVersion string, extraErrorText string) (io.Reader, error) {
	var color string
	if success {
		color = colorGreen
	} else {
		color = colorRed
	}

	if extraErrorText != "" {
		extraErrorText = fmt.Sprintf("Additional Error: %s", extraErrorText)
	}

	var fields []slackAttachmentField
	// targetVersion only not empty when producing autoupgrade payload
	if targetVersion != "" {
		fields = []slackAttachmentField{
			{
				Title: "Target Package Version",
				Value: targetVersion,
				Short: false,
			},
		}
	}

	payload := slackPayload{
		Attachments: []slackAttachment{
			{
				Color:    color,
				Pretext:  fmt.Sprintf("%s%s", pretext, extraErrorText),
				Fallback: fallback,
				Fields:   fields,
			},
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal slack payload")
	}
	return bytes.NewReader(jsonPayload), nil
}
