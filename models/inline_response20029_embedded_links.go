// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20029EmbeddedLinks inline response 200 29 embedded links
// swagger:model inline_response_200_29__embedded__links
type InlineResponse20029EmbeddedLinks struct {

	// download csv
	DownloadCsv *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"download_csv,omitempty"`

	// download pdf
	DownloadPdf *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"download_pdf,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`

	// stack
	Stack *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"stack,omitempty"`
}

// Validate validates this inline response 200 29 embedded links
func (m *InlineResponse20029EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloadCsv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadPdf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20029EmbeddedLinks) validateDownloadCsv(formats strfmt.Registry) error {

	if swag.IsZero(m.DownloadCsv) { // not required
		return nil
	}

	if m.DownloadCsv != nil {
		if err := m.DownloadCsv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download_csv")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20029EmbeddedLinks) validateDownloadPdf(formats strfmt.Registry) error {

	if swag.IsZero(m.DownloadPdf) { // not required
		return nil
	}

	if m.DownloadPdf != nil {
		if err := m.DownloadPdf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download_pdf")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20029EmbeddedLinks) validateSelf(formats strfmt.Registry) error {

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

func (m *InlineResponse20029EmbeddedLinks) validateStack(formats strfmt.Registry) error {

	if swag.IsZero(m.Stack) { // not required
		return nil
	}

	if m.Stack != nil {
		if err := m.Stack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20029EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20029EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse20029EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
