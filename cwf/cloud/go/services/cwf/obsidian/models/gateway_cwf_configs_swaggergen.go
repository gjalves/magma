// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayCwfConfigs CWF configuration for a gateway
// swagger:model gateway_cwf_configs
type GatewayCwfConfigs struct {

	// allowed gre peers
	// Required: true
	AllowedGrePeers AllowedGrePeers `json:"allowed_gre_peers"`

	// gateway health configs
	GatewayHealthConfigs *GatewayHealthConfigs `json:"gateway_health_configs,omitempty"`

	// ipdr export dst
	IPDRExportDst *IPDRExportDst `json:"ipdr_export_dst,omitempty"`
}

// Validate validates this gateway cwf configs
func (m *GatewayCwfConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedGrePeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayHealthConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPDRExportDst(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayCwfConfigs) validateAllowedGrePeers(formats strfmt.Registry) error {

	if err := validate.Required("allowed_gre_peers", "body", m.AllowedGrePeers); err != nil {
		return err
	}

	if err := m.AllowedGrePeers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("allowed_gre_peers")
		}
		return err
	}

	return nil
}

func (m *GatewayCwfConfigs) validateGatewayHealthConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayHealthConfigs) { // not required
		return nil
	}

	if m.GatewayHealthConfigs != nil {
		if err := m.GatewayHealthConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_health_configs")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayCwfConfigs) validateIPDRExportDst(formats strfmt.Registry) error {

	if swag.IsZero(m.IPDRExportDst) { // not required
		return nil
	}

	if m.IPDRExportDst != nil {
		if err := m.IPDRExportDst.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipdr_export_dst")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayCwfConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayCwfConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayCwfConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
