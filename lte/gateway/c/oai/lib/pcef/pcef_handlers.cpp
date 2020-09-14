/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include <grpcpp/impl/codegen/status.h>
#include <string.h>
#include <string>
#include <conversions.h>

#include "pcef_handlers.h"
#include "PCEFClient.h"
#include "MobilityClientAPI.h"
#include "intertask_interface.h"
#include "intertask_interface_types.h"
#include "itti_types.h"
#include "lte/protos/session_manager.pb.h"
#include "lte/protos/subscriberdb.pb.h"
#include "spgw_types.h"

extern "C" {
}

#define ULI_DATA_SIZE 13

static void create_session_response(
  spgw_state_t* state,
  const std::string& imsi,
  const std::string& apn,
  itti_sgi_create_end_point_response_t sgi_response,
  s5_create_session_request_t session_request,
  const grpc::Status& status,
  s_plus_p_gw_eps_bearer_context_information_t* ctx_p)
{
  s5_create_session_response_t s5_response = {0};
  s5_response.context_teid = session_request.context_teid;
  s5_response.eps_bearer_id = session_request.eps_bearer_id;
  s5_response.sgi_create_endpoint_resp = sgi_response;
  s5_response.failure_cause = S5_OK;

  if (!status.ok()) {
    //BUFFER_TO_IN_ADDR (sgi_response.paa.ipv4_address, addr);
    release_ipv4_address(imsi.c_str(), apn.c_str(),
                         &sgi_response.paa.ipv4_address);
    s5_response.failure_cause = PCEF_FAILURE;
  }
  handle_s5_create_session_response(state, ctx_p, s5_response);
}

static void pcef_fill_create_session_req(
  const struct pcef_create_session_data *session_data,
  magma::LocalCreateSessionRequest *sreq)
{
  sreq->set_apn(session_data->apn);
  sreq->set_msisdn(session_data->msisdn, session_data->msisdn_len);
  sreq->set_spgw_ipv4(session_data->sgw_ip);
  sreq->set_plmn_id(session_data->mcc_mnc, session_data->mcc_mnc_len);
  sreq->set_imsi_plmn_id(
    session_data->imsi_mcc_mnc, session_data->imsi_mcc_mnc_len);

  if (session_data->imeisv_exists) {
    sreq->set_imei(session_data->imeisv, IMEISV_DIGITS_MAX);
  }
  if (session_data->uli_exists) {
    sreq->set_user_location(session_data->uli, ULI_DATA_SIZE);
  }

  // QoS Info
  sreq->mutable_qos_info()->set_apn_ambr_dl(session_data->ambr_dl);
  sreq->mutable_qos_info()->set_apn_ambr_ul(session_data->ambr_ul);
  sreq->mutable_qos_info()->set_priority_level(session_data->pl);
  sreq->mutable_qos_info()->set_preemption_capability(session_data->pci);
  sreq->mutable_qos_info()->set_preemption_vulnerability(session_data->pvi);
  sreq->mutable_qos_info()->set_qos_class_id(session_data->qci);
}

void pcef_create_session(
  spgw_state_t* state,
  char* imsi,
  char* ip,
  const pcef_create_session_data* session_data,
  itti_sgi_create_end_point_response_t sgi_response,
  s5_create_session_request_t session_request,
  s_plus_p_gw_eps_bearer_context_information_t* ctx_p)
{
  auto imsi_str = std::string(imsi);
  auto ip_str = std::string(ip);
  // Change ip to spgw_ip. Get it from sgw_app_t sgw_app;
  magma::LocalCreateSessionRequest sreq;

  sreq.mutable_sid()->set_id("IMSI" + imsi_str);
  sreq.set_rat_type(magma::RATType::TGPP_LTE);
  sreq.set_ue_ipv4(ip_str);
  sreq.set_bearer_id(session_request.eps_bearer_id);
  pcef_fill_create_session_req(session_data, &sreq);

  auto apn = std::string(session_data->apn);
  // call the `CreateSession` gRPC method and execute the inline function
  magma::PCEFClient::create_session(
    sreq,
    [imsi_str, apn, sgi_response, session_request, ctx_p, state](
      grpc::Status status, magma::LocalCreateSessionResponse response) {
      create_session_response(
        state, imsi_str, apn, sgi_response, session_request, status, ctx_p);
    });
}

bool pcef_end_session(char *imsi, char *apn)
{
  magma::LocalEndSessionRequest request;
  request.mutable_sid()->set_id("IMSI" + std::string(imsi));
  request.set_apn(apn);
  magma::PCEFClient::end_session(
    request, [&](grpc::Status status, magma::LocalEndSessionResponse response) {
      return; // For now, do nothing. TODO: handle errors asynchronously
    });
  return true;
}
