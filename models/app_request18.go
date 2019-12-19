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

// AppRequest18 app request 18
// swagger:model app_request_18
type AppRequest18 struct {

	// For Deploy-hosted InfluxDB instances
	// Format: uri
	Database strfmt.URI `json:"database,omitempty"`

	// For Deploy-hosted InfluxDB instances
	DatabaseID int64 `json:"database_id,omitempty"`

	// drain configuration
	DrainConfiguration *AccountsaccountIdmetricDrainsDrainConfiguration `json:"drain_configuration,omitempty"`

	// drain type
	// Required: true
	DrainType *string `json:"drain_type"`

	// handle
	// Required: true
	Handle *string `json:"handle"`
}

// Validate validates this app request 18
func (m *AppRequest18) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrainConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrainType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest18) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if err := validate.FormatOf("database", "body", "uri", m.Database.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest18) validateDrainConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.DrainConfiguration) { // not required
		return nil
	}

	if m.DrainConfiguration != nil {
		if err := m.DrainConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drain_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *AppRequest18) validateDrainType(formats strfmt.Registry) error {

	if err := validate.Required("drain_type", "body", m.DrainType); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest18) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest18) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest18) UnmarshalBinary(b []byte) error {
	var res AppRequest18
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
