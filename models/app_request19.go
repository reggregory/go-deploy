// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AppRequest19 app request 19
// swagger:model app_request_19
type AppRequest19 struct {

	// cancelled
	Cancelled bool `json:"cancelled,omitempty"`
}

// Validate validates this app request 19
func (m *AppRequest19) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest19) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest19) UnmarshalBinary(b []byte) error {
	var res AppRequest19
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
