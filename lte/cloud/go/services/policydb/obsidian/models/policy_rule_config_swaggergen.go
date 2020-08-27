// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyRuleConfig policy rule config
// swagger:model policy_rule_config
type PolicyRuleConfig struct {

	// app name
	// Enum: [NO_APP_NAME FACEBOOK FACEBOOK_MESSENGER INSTAGRAM YOUTUBE GOOGLE GMAIL GOOGLE_DOCS NETFLIX APPLE MICROSOFT REDDIT WHATSAPP GOOGLE_PLAY APPSTORE AMAZON WECHAT TIKTOK TWITTER WIKIPEDIA GOOGLE_MAPS YAHOO IMO]
	AppName string `json:"app_name,omitempty"`

	// app service type
	// Enum: [NO_SERVICE_TYPE CHAT AUDIO VIDEO]
	AppServiceType string `json:"app_service_type,omitempty"`

	// flow list
	// Required: true
	FlowList []*FlowDescription `json:"flow_list"`

	// monitoring key
	MonitoringKey string `json:"monitoring_key,omitempty"`

	// priority
	// Required: true
	Priority *uint32 `json:"priority"`

	// rating group
	RatingGroup uint32 `json:"rating_group,omitempty"`

	// redirect
	Redirect *RedirectInformation `json:"redirect,omitempty"`

	// tracking type
	// Enum: [ONLY_OCS ONLY_PCRF OCS_AND_PCRF NO_TRACKING]
	TrackingType string `json:"tracking_type,omitempty"`
}

// Validate validates this policy rule config
func (m *PolicyRuleConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var policyRuleConfigTypeAppNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_APP_NAME","FACEBOOK","FACEBOOK_MESSENGER","INSTAGRAM","YOUTUBE","GOOGLE","GMAIL","GOOGLE_DOCS","NETFLIX","APPLE","MICROSOFT","REDDIT","WHATSAPP","GOOGLE_PLAY","APPSTORE","AMAZON","WECHAT","TIKTOK","TWITTER","WIKIPEDIA","GOOGLE_MAPS","YAHOO","IMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleConfigTypeAppNamePropEnum = append(policyRuleConfigTypeAppNamePropEnum, v)
	}
}

const (

	// PolicyRuleConfigAppNameNOAPPNAME captures enum value "NO_APP_NAME"
	PolicyRuleConfigAppNameNOAPPNAME string = "NO_APP_NAME"

	// PolicyRuleConfigAppNameFACEBOOK captures enum value "FACEBOOK"
	PolicyRuleConfigAppNameFACEBOOK string = "FACEBOOK"

	// PolicyRuleConfigAppNameFACEBOOKMESSENGER captures enum value "FACEBOOK_MESSENGER"
	PolicyRuleConfigAppNameFACEBOOKMESSENGER string = "FACEBOOK_MESSENGER"

	// PolicyRuleConfigAppNameINSTAGRAM captures enum value "INSTAGRAM"
	PolicyRuleConfigAppNameINSTAGRAM string = "INSTAGRAM"

	// PolicyRuleConfigAppNameYOUTUBE captures enum value "YOUTUBE"
	PolicyRuleConfigAppNameYOUTUBE string = "YOUTUBE"

	// PolicyRuleConfigAppNameGOOGLE captures enum value "GOOGLE"
	PolicyRuleConfigAppNameGOOGLE string = "GOOGLE"

	// PolicyRuleConfigAppNameGMAIL captures enum value "GMAIL"
	PolicyRuleConfigAppNameGMAIL string = "GMAIL"

	// PolicyRuleConfigAppNameGOOGLEDOCS captures enum value "GOOGLE_DOCS"
	PolicyRuleConfigAppNameGOOGLEDOCS string = "GOOGLE_DOCS"

	// PolicyRuleConfigAppNameNETFLIX captures enum value "NETFLIX"
	PolicyRuleConfigAppNameNETFLIX string = "NETFLIX"

	// PolicyRuleConfigAppNameAPPLE captures enum value "APPLE"
	PolicyRuleConfigAppNameAPPLE string = "APPLE"

	// PolicyRuleConfigAppNameMICROSOFT captures enum value "MICROSOFT"
	PolicyRuleConfigAppNameMICROSOFT string = "MICROSOFT"

	// PolicyRuleConfigAppNameREDDIT captures enum value "REDDIT"
	PolicyRuleConfigAppNameREDDIT string = "REDDIT"

	// PolicyRuleConfigAppNameWHATSAPP captures enum value "WHATSAPP"
	PolicyRuleConfigAppNameWHATSAPP string = "WHATSAPP"

	// PolicyRuleConfigAppNameGOOGLEPLAY captures enum value "GOOGLE_PLAY"
	PolicyRuleConfigAppNameGOOGLEPLAY string = "GOOGLE_PLAY"

	// PolicyRuleConfigAppNameAPPSTORE captures enum value "APPSTORE"
	PolicyRuleConfigAppNameAPPSTORE string = "APPSTORE"

	// PolicyRuleConfigAppNameAMAZON captures enum value "AMAZON"
	PolicyRuleConfigAppNameAMAZON string = "AMAZON"

	// PolicyRuleConfigAppNameWECHAT captures enum value "WECHAT"
	PolicyRuleConfigAppNameWECHAT string = "WECHAT"

	// PolicyRuleConfigAppNameTIKTOK captures enum value "TIKTOK"
	PolicyRuleConfigAppNameTIKTOK string = "TIKTOK"

	// PolicyRuleConfigAppNameTWITTER captures enum value "TWITTER"
	PolicyRuleConfigAppNameTWITTER string = "TWITTER"

	// PolicyRuleConfigAppNameWIKIPEDIA captures enum value "WIKIPEDIA"
	PolicyRuleConfigAppNameWIKIPEDIA string = "WIKIPEDIA"

	// PolicyRuleConfigAppNameGOOGLEMAPS captures enum value "GOOGLE_MAPS"
	PolicyRuleConfigAppNameGOOGLEMAPS string = "GOOGLE_MAPS"

	// PolicyRuleConfigAppNameYAHOO captures enum value "YAHOO"
	PolicyRuleConfigAppNameYAHOO string = "YAHOO"

	// PolicyRuleConfigAppNameIMO captures enum value "IMO"
	PolicyRuleConfigAppNameIMO string = "IMO"
)

// prop value enum
func (m *PolicyRuleConfig) validateAppNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, policyRuleConfigTypeAppNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRuleConfig) validateAppName(formats strfmt.Registry) error {

	if swag.IsZero(m.AppName) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppNameEnum("app_name", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

var policyRuleConfigTypeAppServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_SERVICE_TYPE","CHAT","AUDIO","VIDEO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleConfigTypeAppServiceTypePropEnum = append(policyRuleConfigTypeAppServiceTypePropEnum, v)
	}
}

const (

	// PolicyRuleConfigAppServiceTypeNOSERVICETYPE captures enum value "NO_SERVICE_TYPE"
	PolicyRuleConfigAppServiceTypeNOSERVICETYPE string = "NO_SERVICE_TYPE"

	// PolicyRuleConfigAppServiceTypeCHAT captures enum value "CHAT"
	PolicyRuleConfigAppServiceTypeCHAT string = "CHAT"

	// PolicyRuleConfigAppServiceTypeAUDIO captures enum value "AUDIO"
	PolicyRuleConfigAppServiceTypeAUDIO string = "AUDIO"

	// PolicyRuleConfigAppServiceTypeVIDEO captures enum value "VIDEO"
	PolicyRuleConfigAppServiceTypeVIDEO string = "VIDEO"
)

// prop value enum
func (m *PolicyRuleConfig) validateAppServiceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, policyRuleConfigTypeAppServiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRuleConfig) validateAppServiceType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppServiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppServiceTypeEnum("app_service_type", "body", m.AppServiceType); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleConfig) validateFlowList(formats strfmt.Registry) error {

	if err := validate.Required("flow_list", "body", m.FlowList); err != nil {
		return err
	}

	for i := 0; i < len(m.FlowList); i++ {
		if swag.IsZero(m.FlowList[i]) { // not required
			continue
		}

		if m.FlowList[i] != nil {
			if err := m.FlowList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flow_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRuleConfig) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleConfig) validateRedirect(formats strfmt.Registry) error {

	if swag.IsZero(m.Redirect) { // not required
		return nil
	}

	if m.Redirect != nil {
		if err := m.Redirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect")
			}
			return err
		}
	}

	return nil
}

var policyRuleConfigTypeTrackingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONLY_OCS","ONLY_PCRF","OCS_AND_PCRF","NO_TRACKING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleConfigTypeTrackingTypePropEnum = append(policyRuleConfigTypeTrackingTypePropEnum, v)
	}
}

const (

	// PolicyRuleConfigTrackingTypeONLYOCS captures enum value "ONLY_OCS"
	PolicyRuleConfigTrackingTypeONLYOCS string = "ONLY_OCS"

	// PolicyRuleConfigTrackingTypeONLYPCRF captures enum value "ONLY_PCRF"
	PolicyRuleConfigTrackingTypeONLYPCRF string = "ONLY_PCRF"

	// PolicyRuleConfigTrackingTypeOCSANDPCRF captures enum value "OCS_AND_PCRF"
	PolicyRuleConfigTrackingTypeOCSANDPCRF string = "OCS_AND_PCRF"

	// PolicyRuleConfigTrackingTypeNOTRACKING captures enum value "NO_TRACKING"
	PolicyRuleConfigTrackingTypeNOTRACKING string = "NO_TRACKING"
)

// prop value enum
func (m *PolicyRuleConfig) validateTrackingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, policyRuleConfigTypeTrackingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRuleConfig) validateTrackingType(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrackingTypeEnum("tracking_type", "body", m.TrackingType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleConfig) UnmarshalBinary(b []byte) error {
	var res PolicyRuleConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
