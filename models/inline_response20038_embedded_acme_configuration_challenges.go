// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineResponse20038EmbeddedAcmeConfigurationChallenges inline response 200 38 embedded acme configuration challenges
// swagger:model inline_response_200_38__embedded_acme_configuration_challenges
type InlineResponse20038EmbeddedAcmeConfigurationChallenges struct {

	// from
	From *InlineResponse20038EmbeddedAcmeConfigurationFrom `json:"from,omitempty"`

	// method
	// Enum: [http01 dns01]
	Method string `json:"method,omitempty"`

	// to
	To []*InlineResponse20038EmbeddedAcmeConfigurationTo `json:"to"`
}

// Validate validates this inline response 200 38 embedded acme configuration challenges
func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

var inlineResponse20038EmbeddedAcmeConfigurationChallengesTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http01","dns01"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inlineResponse20038EmbeddedAcmeConfigurationChallengesTypeMethodPropEnum = append(inlineResponse20038EmbeddedAcmeConfigurationChallengesTypeMethodPropEnum, v)
	}
}

const (

	// InlineResponse20038EmbeddedAcmeConfigurationChallengesMethodHttp01 captures enum value "http01"
	InlineResponse20038EmbeddedAcmeConfigurationChallengesMethodHttp01 string = "http01"

	// InlineResponse20038EmbeddedAcmeConfigurationChallengesMethodDns01 captures enum value "dns01"
	InlineResponse20038EmbeddedAcmeConfigurationChallengesMethodDns01 string = "dns01"
)

// prop value enum
func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inlineResponse20038EmbeddedAcmeConfigurationChallengesTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	for i := 0; i < len(m.To); i++ {
		if swag.IsZero(m.To[i]) { // not required
			continue
		}

		if m.To[i] != nil {
			if err := m.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20038EmbeddedAcmeConfigurationChallenges) UnmarshalBinary(b []byte) error {
	var res InlineResponse20038EmbeddedAcmeConfigurationChallenges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}