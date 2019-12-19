// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2004EmbeddedEmbeddedCurrentImage inline response 200 4 embedded embedded current image
// swagger:model inline_response_200_4__embedded__embedded_current_image
type InlineResponse2004EmbeddedEmbeddedCurrentImage struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse2004EmbeddedEmbeddedCurrentImageLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// docker ref
	DockerRef *string `json:"docker_ref,omitempty"`

	// docker repo
	DockerRepo *string `json:"docker_repo,omitempty"`

	// dualstack hint
	DualstackHint *int64 `json:"dualstack_hint,omitempty"`

	// exposed ports
	ExposedPorts []int64 `json:"exposed_ports"`

	// git ref
	GitRef *string `json:"git_ref,omitempty"`

	// git repo
	GitRepo *string `json:"git_repo,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// platform
	Platform *string `json:"platform,omitempty"`

	// release
	Release *string `json:"release,omitempty"`

	// scan
	Scan *string `json:"scan,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 4 embedded embedded current image
func (m *InlineResponse2004EmbeddedEmbeddedCurrentImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2004EmbeddedEmbeddedCurrentImage) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse2004EmbeddedEmbeddedCurrentImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2004EmbeddedEmbeddedCurrentImage) UnmarshalBinary(b []byte) error {
	var res InlineResponse2004EmbeddedEmbeddedCurrentImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
