// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AppRequest17 app request 17
// swagger:model app_request_17
type AppRequest17 struct {

	// drain apps
	DrainApps bool `json:"drain_apps,omitempty"`

	// drain databases
	DrainDatabases bool `json:"drain_databases,omitempty"`

	// drain ephemeral sessions
	DrainEphemeralSessions bool `json:"drain_ephemeral_sessions,omitempty"`

	// drain proxies
	DrainProxies bool `json:"drain_proxies,omitempty"`
}

// Validate validates this app request 17
func (m *AppRequest17) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest17) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest17) UnmarshalBinary(b []byte) error {
	var res AppRequest17
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
