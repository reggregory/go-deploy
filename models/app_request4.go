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

// AppRequest4 app request 4
// swagger:model app_request_4
type AppRequest4 struct {

	// acme
	Acme bool `json:"acme,omitempty"`

	// certificate body
	// Required: true
	CertificateBody *string `json:"certificate_body"`

	// private key
	// Required: true
	PrivateKey *string `json:"private_key"`
}

// Validate validates this app request 4
func (m *AppRequest4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest4) validateCertificateBody(formats strfmt.Registry) error {

	if err := validate.Required("certificate_body", "body", m.CertificateBody); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest4) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest4) UnmarshalBinary(b []byte) error {
	var res AppRequest4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
