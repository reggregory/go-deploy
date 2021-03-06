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

// InlineResponse2003EmbeddedEmbedded inline response 200 3 embedded embedded
// swagger:model inline_response_200_3__embedded__embedded
type InlineResponse2003EmbeddedEmbedded struct {

	// current image
	CurrentImage *InlineResponse2003EmbeddedEmbeddedCurrentImage `json:"current_image,omitempty"`

	// last deploy operation
	LastDeployOperation *InlineResponse2003EmbeddedEmbeddedLastOperation `json:"last_deploy_operation,omitempty"`

	// last operation
	LastOperation *InlineResponse2003EmbeddedEmbeddedLastOperation `json:"last_operation,omitempty"`

	// services
	Services []*InlineResponse2003EmbeddedEmbeddedServices `json:"services"`
}

// Validate validates this inline response 200 3 embedded embedded
func (m *InlineResponse2003EmbeddedEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastDeployOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2003EmbeddedEmbedded) validateCurrentImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentImage) { // not required
		return nil
	}

	if m.CurrentImage != nil {
		if err := m.CurrentImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_image")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedEmbedded) validateLastDeployOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.LastDeployOperation) { // not required
		return nil
	}

	if m.LastDeployOperation != nil {
		if err := m.LastDeployOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_deploy_operation")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedEmbedded) validateLastOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.LastOperation) { // not required
		return nil
	}

	if m.LastOperation != nil {
		if err := m.LastOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_operation")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedEmbedded) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedEmbedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003EmbeddedEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
