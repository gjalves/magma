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

// AgwTestConfig Shared AGW test configuration for auto-upgrades
// swagger:model agw_test_config
type AgwTestConfig struct {

	// URL of debian package repo
	// Required: true
	// Min Length: 1
	PackageRepo *string `json:"package_repo"`

	// Release channel for package repo (stretch-beta, stretch-dev, etc)
	// Required: true
	// Min Length: 1
	ReleaseChannel *string `json:"release_channel"`

	// Slack webhook for test notifications
	// Required: true
	// Min Length: 1
	SLACKWebhook *string `json:"slack_webhook"`

	// Gateway ID of the target gateway
	// Required: true
	// Min Length: 1
	TargetGatewayID *string `json:"target_gateway_id"`

	// ID of upgrade tier to bump when new version is found in package repo
	// Required: true
	// Min Length: 1
	TargetTier *string `json:"target_tier"`
}

// Validate validates this agw test config
func (m *AgwTestConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSLACKWebhook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgwTestConfig) validatePackageRepo(formats strfmt.Registry) error {

	if err := validate.Required("package_repo", "body", m.PackageRepo); err != nil {
		return err
	}

	if err := validate.MinLength("package_repo", "body", string(*m.PackageRepo), 1); err != nil {
		return err
	}

	return nil
}

func (m *AgwTestConfig) validateReleaseChannel(formats strfmt.Registry) error {

	if err := validate.Required("release_channel", "body", m.ReleaseChannel); err != nil {
		return err
	}

	if err := validate.MinLength("release_channel", "body", string(*m.ReleaseChannel), 1); err != nil {
		return err
	}

	return nil
}

func (m *AgwTestConfig) validateSLACKWebhook(formats strfmt.Registry) error {

	if err := validate.Required("slack_webhook", "body", m.SLACKWebhook); err != nil {
		return err
	}

	if err := validate.MinLength("slack_webhook", "body", string(*m.SLACKWebhook), 1); err != nil {
		return err
	}

	return nil
}

func (m *AgwTestConfig) validateTargetGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("target_gateway_id", "body", m.TargetGatewayID); err != nil {
		return err
	}

	if err := validate.MinLength("target_gateway_id", "body", string(*m.TargetGatewayID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AgwTestConfig) validateTargetTier(formats strfmt.Registry) error {

	if err := validate.Required("target_tier", "body", m.TargetTier); err != nil {
		return err
	}

	if err := validate.MinLength("target_tier", "body", string(*m.TargetTier), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgwTestConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgwTestConfig) UnmarshalBinary(b []byte) error {
	var res AgwTestConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
