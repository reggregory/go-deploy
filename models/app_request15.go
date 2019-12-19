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

// AppRequest15 app request 15
// swagger:model app_request_15
type AppRequest15 struct {

	// For Deploy-hosted Elasticsearch drains
	// Format: uri
	Database strfmt.URI `json:"database,omitempty"`

	// For Deploy-hosted Elasticsearch drains
	DatabaseID int64 `json:"database_id,omitempty"`

	// drain apps
	DrainApps bool `json:"drain_apps,omitempty"`

	// drain databases
	DrainDatabases bool `json:"drain_databases,omitempty"`

	// drain ephemeral sessions
	DrainEphemeralSessions bool `json:"drain_ephemeral_sessions,omitempty"`

	// For syslog-tls-tcp drains
	// Format: uri
	DrainHost strfmt.URI `json:"drain_host,omitempty"`

	// For Tail drains
	DrainPassword string `json:"drain_password,omitempty"`

	// For syslog-tls-tcp drains
	DrainPort int64 `json:"drain_port,omitempty"`

	// drain proxies
	DrainProxies bool `json:"drain_proxies,omitempty"`

	// drain type
	// Required: true
	DrainType *string `json:"drain_type"`

	// DEPRECATED: For legacy Elasticsearch drains
	DrainUsername string `json:"drain_username,omitempty"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// logging token
	LoggingToken string `json:"logging_token,omitempty"`

	// For HTTPs Post drains
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this app request 15
func (m *AppRequest15) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrainHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrainType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest15) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if err := validate.FormatOf("database", "body", "uri", m.Database.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest15) validateDrainHost(formats strfmt.Registry) error {

	if swag.IsZero(m.DrainHost) { // not required
		return nil
	}

	if err := validate.FormatOf("drain_host", "body", "uri", m.DrainHost.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest15) validateDrainType(formats strfmt.Registry) error {

	if err := validate.Required("drain_type", "body", m.DrainType); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest15) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest15) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest15) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest15) UnmarshalBinary(b []byte) error {
	var res AppRequest15
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
