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

// InlineResponse2002Embedded inline response 200 2 embedded
// swagger:model inline_response_200_2__embedded
type InlineResponse2002Embedded struct {

	// activity reports
	ActivityReports []*InlineResponse2002EmbeddedActivityReports `json:"activity_reports"`
}

// Validate validates this inline response 200 2 embedded
func (m *InlineResponse2002Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityReports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2002Embedded) validateActivityReports(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityReports) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivityReports); i++ {
		if swag.IsZero(m.ActivityReports[i]) { // not required
			continue
		}

		if m.ActivityReports[i] != nil {
			if err := m.ActivityReports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activity_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2002Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2002Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse2002Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
