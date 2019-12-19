// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2002Links inline response 200 2 links
// swagger:model inline_response_200_2__links
type InlineResponse2002Links struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"account,omitempty"`

	// next
	Next *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"next,omitempty"`

	// prev
	Prev *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"prev,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`
}

// Validate validates this inline response 200 2 links
func (m *InlineResponse2002Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
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

func (m *InlineResponse2002Links) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2002Links) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2002Links) validatePrev(formats strfmt.Registry) error {

	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2002Links) validateSelf(formats strfmt.Registry) error {

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
func (m *InlineResponse2002Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2002Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse2002Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
