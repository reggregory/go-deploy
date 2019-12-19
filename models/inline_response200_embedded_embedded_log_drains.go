// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse200EmbeddedEmbeddedLogDrains inline response 200 embedded embedded log drains
// swagger:model inline_response_200__embedded__embedded_log_drains
type InlineResponse200EmbeddedEmbeddedLogDrains struct {

	// links
	Links *InlineResponse200EmbeddedEmbeddedLinks `json:"_links,omitempty"`
}

// Validate validates this inline response 200 embedded embedded log drains
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLogDrains) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) UnmarshalBinary(b []byte) error {
	var res InlineResponse200EmbeddedEmbeddedLogDrains
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}