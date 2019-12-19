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

// InlineResponse20048 inline response 200 48
// swagger:model inline_response_200_48
type InlineResponse20048 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse20047EmbeddedLinks `json:"_links"`

	// acme
	// Required: true
	Acme *bool `json:"acme"`

	// acme configuration
	// Required: true
	AcmeConfiguration *InlineResponse20047EmbeddedAcmeConfiguration `json:"acme_configuration"`

	// acme dns challenge host
	// Required: true
	AcmeDNSChallengeHost *string `json:"acme_dns_challenge_host"`

	// acme status
	// Required: true
	AcmeStatus *string `json:"acme_status"`

	// application load balancer arn
	// Required: true
	ApplicationLoadBalancerArn *string `json:"application_load_balancer_arn"`

	// container exposed ports
	// Required: true
	ContainerExposedPorts []int64 `json:"container_exposed_ports"`

	// container port
	// Required: true
	ContainerPort *int64 `json:"container_port"`

	// container ports
	// Required: true
	ContainerPorts []int64 `json:"container_ports"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// default
	// Required: true
	Default *bool `json:"default"`

	// docker name
	// Required: true
	DockerName *string `json:"docker_name"`

	// elastic load balancer name
	// Required: true
	ElasticLoadBalancerName *string `json:"elastic_load_balancer_name"`

	// external host
	// Required: true
	ExternalHost *string `json:"external_host"`

	// external http port
	// Required: true
	ExternalHTTPPort *int64 `json:"external_http_port"`

	// external https port
	// Required: true
	ExternalHTTPSPort *int64 `json:"external_https_port"`

	// host mapped ports
	// Required: true
	HostMappedPorts []int64 `json:"host_mapped_ports"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// internal health port
	// Required: true
	InternalHealthPort *int64 `json:"internal_health_port"`

	// internal host
	// Required: true
	InternalHost *string `json:"internal_host"`

	// internal http port
	// Required: true
	InternalHTTPPort *int64 `json:"internal_http_port"`

	// internal https port
	// Required: true
	InternalHTTPSPort *int64 `json:"internal_https_port"`

	// ip whitelist
	// Required: true
	IPWhitelist []string `json:"ip_whitelist"`

	// platform
	// Required: true
	Platform *string `json:"platform"`

	// security group id
	// Required: true
	SecurityGroupID *string `json:"security_group_id"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`

	// user domain
	// Required: true
	UserDomain *string `json:"user_domain"`

	// virtual domain
	// Required: true
	VirtualDomain *string `json:"virtual_domain"`
}

// Validate validates this inline response 200 48
func (m *InlineResponse20048) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcmeConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcmeDNSChallengeHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcmeStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationLoadBalancerArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerExposedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticLoadBalancerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalHTTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalHTTPSPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostMappedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalHealthPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalHTTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalHTTPSPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20048) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateLinks(formats strfmt.Registry) error {

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

func (m *InlineResponse20048) validateAcme(formats strfmt.Registry) error {

	if err := validate.Required("acme", "body", m.Acme); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateAcmeConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("acme_configuration", "body", m.AcmeConfiguration); err != nil {
		return err
	}

	if m.AcmeConfiguration != nil {
		if err := m.AcmeConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acme_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20048) validateAcmeDNSChallengeHost(formats strfmt.Registry) error {

	if err := validate.Required("acme_dns_challenge_host", "body", m.AcmeDNSChallengeHost); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateAcmeStatus(formats strfmt.Registry) error {

	if err := validate.Required("acme_status", "body", m.AcmeStatus); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateApplicationLoadBalancerArn(formats strfmt.Registry) error {

	if err := validate.Required("application_load_balancer_arn", "body", m.ApplicationLoadBalancerArn); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateContainerExposedPorts(formats strfmt.Registry) error {

	if err := validate.Required("container_exposed_ports", "body", m.ContainerExposedPorts); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateContainerPort(formats strfmt.Registry) error {

	if err := validate.Required("container_port", "body", m.ContainerPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateContainerPorts(formats strfmt.Registry) error {

	if err := validate.Required("container_ports", "body", m.ContainerPorts); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateDockerName(formats strfmt.Registry) error {

	if err := validate.Required("docker_name", "body", m.DockerName); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateElasticLoadBalancerName(formats strfmt.Registry) error {

	if err := validate.Required("elastic_load_balancer_name", "body", m.ElasticLoadBalancerName); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateExternalHost(formats strfmt.Registry) error {

	if err := validate.Required("external_host", "body", m.ExternalHost); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateExternalHTTPPort(formats strfmt.Registry) error {

	if err := validate.Required("external_http_port", "body", m.ExternalHTTPPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateExternalHTTPSPort(formats strfmt.Registry) error {

	if err := validate.Required("external_https_port", "body", m.ExternalHTTPSPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateHostMappedPorts(formats strfmt.Registry) error {

	if err := validate.Required("host_mapped_ports", "body", m.HostMappedPorts); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateInternalHealthPort(formats strfmt.Registry) error {

	if err := validate.Required("internal_health_port", "body", m.InternalHealthPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateInternalHost(formats strfmt.Registry) error {

	if err := validate.Required("internal_host", "body", m.InternalHost); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateInternalHTTPPort(formats strfmt.Registry) error {

	if err := validate.Required("internal_http_port", "body", m.InternalHTTPPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateInternalHTTPSPort(formats strfmt.Registry) error {

	if err := validate.Required("internal_https_port", "body", m.InternalHTTPSPort); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateIPWhitelist(formats strfmt.Registry) error {

	if err := validate.Required("ip_whitelist", "body", m.IPWhitelist); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateSecurityGroupID(formats strfmt.Registry) error {

	if err := validate.Required("security_group_id", "body", m.SecurityGroupID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateUserDomain(formats strfmt.Registry) error {

	if err := validate.Required("user_domain", "body", m.UserDomain); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20048) validateVirtualDomain(formats strfmt.Registry) error {

	if err := validate.Required("virtual_domain", "body", m.VirtualDomain); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20048) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20048) UnmarshalBinary(b []byte) error {
	var res InlineResponse20048
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
