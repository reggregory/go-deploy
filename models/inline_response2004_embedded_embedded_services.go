// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2004EmbeddedEmbeddedServices inline response 200 4 embedded embedded services
// swagger:model inline_response_200_4__embedded__embedded_services
type InlineResponse2004EmbeddedEmbeddedServices struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse2004EmbeddedEmbeddedLinks `json:"_links,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// container count
	ContainerCount int64 `json:"container_count,omitempty"`

	// container memory limit mb
	ContainerMemoryLimitMb *int64 `json:"container_memory_limit_mb,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// docker ref
	DockerRef *string `json:"docker_ref,omitempty"`

	// docker repo
	DockerRepo *string `json:"docker_repo,omitempty"`

	// handle
	Handle string `json:"handle,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// process type
	ProcessType string `json:"process_type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this inline response 200 4 embedded embedded services
func (m *InlineResponse2004EmbeddedEmbeddedServices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2004EmbeddedEmbeddedServices) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse2004EmbeddedEmbeddedServices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2004EmbeddedEmbeddedServices) UnmarshalBinary(b []byte) error {
	var res InlineResponse2004EmbeddedEmbeddedServices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
