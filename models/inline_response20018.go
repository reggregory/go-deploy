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

// InlineResponse20018 inline response 200 18
// swagger:model inline_response_200_18
type InlineResponse20018 struct {

	// embedded
	// Required: true
	Embedded *InlineResponse20018Embedded `json:"_embedded"`

	// links
	// Required: true
	Links *InlineResponse2002Links `json:"_links"`

	// current page
	// Required: true
	CurrentPage *int64 `json:"current_page"`

	// per page
	// Required: true
	PerPage *int64 `json:"per_page"`

	// total count
	// Required: true
	TotalCount *int64 `json:"total_count"`
}

// Validate validates this inline response 200 18
func (m *InlineResponse20018) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20018) validateEmbedded(formats strfmt.Registry) error {

	if err := validate.Required("_embedded", "body", m.Embedded); err != nil {
		return err
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20018) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
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

func (m *InlineResponse20018) validateCurrentPage(formats strfmt.Registry) error {

	if err := validate.Required("current_page", "body", m.CurrentPage); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validatePerPage(formats strfmt.Registry) error {

	if err := validate.Required("per_page", "body", m.PerPage); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20018) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20018) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20018) UnmarshalBinary(b []byte) error {
	var res InlineResponse20018
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}