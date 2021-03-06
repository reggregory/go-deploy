// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20021Links inline response 200 21 links
// swagger:model inline_response_200_21__links
type InlineResponse20021Links struct {

	// accounts
	Accounts *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"accounts,omitempty"`

	// apps
	Apps *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"apps,omitempty"`

	// database images
	DatabaseImages *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"database_images,omitempty"`

	// databases
	Databases *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"databases,omitempty"`

	// stacks
	Stacks *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"stacks,omitempty"`
}

// Validate validates this inline response 200 21 links
func (m *InlineResponse20021Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20021Links) validateAccounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	if m.Accounts != nil {
		if err := m.Accounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validateApps(formats strfmt.Registry) error {

	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	if m.Apps != nil {
		if err := m.Apps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validateDatabaseImages(formats strfmt.Registry) error {

	if swag.IsZero(m.DatabaseImages) { // not required
		return nil
	}

	if m.DatabaseImages != nil {
		if err := m.DatabaseImages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database_images")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if m.Databases != nil {
		if err := m.Databases.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databases")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20021Links) validateStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Stacks) { // not required
		return nil
	}

	if m.Stacks != nil {
		if err := m.Stacks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stacks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20021Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20021Links) UnmarshalBinary(b []byte) error {
	var res InlineResponse20021Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
