// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AppRequest5 app request 5
// swagger:model app_request_5
type AppRequest5 struct {

	// certificate arn
	CertificateArn string `json:"certificate_arn,omitempty"`
}

// Validate validates this app request 5
func (m *AppRequest5) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest5) UnmarshalBinary(b []byte) error {
	var res AppRequest5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
