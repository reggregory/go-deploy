// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20031Embedded inline response 200 31 embedded
// swagger:model inline_response_200_31__embedded
type InlineResponse20031Embedded struct {

	// log drains
	LogDrains []*InlineResponse20031EmbeddedLogDrains `json:"log_drains"`
}

// Validate validates this inline response 200 31 embedded
func (m *InlineResponse20031Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogDrains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20031Embedded) validateLogDrains(formats strfmt.Registry) error {

	if swag.IsZero(m.LogDrains) { // not required
		return nil
	}

	for i := 0; i < len(m.LogDrains); i++ {
		if swag.IsZero(m.LogDrains[i]) { // not required
			continue
		}

		if m.LogDrains[i] != nil {
			if err := m.LogDrains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_drains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20031Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20031Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse20031Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
