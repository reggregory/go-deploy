// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20016Links inline response 200 16 links
// swagger:model inline_response_200_16__links
type InlineResponse20016Links struct {

	// database
	Database *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"database,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`
}

// Validate validates this inline response 200 16 links
func (m *InlineResponse20016Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20016Links) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20016Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20016Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20016Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20016Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
