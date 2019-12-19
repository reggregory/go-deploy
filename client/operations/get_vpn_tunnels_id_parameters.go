// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVpnTunnelsIDParams creates a new GetVpnTunnelsIDParams object
// with the default values initialized.
func NewGetVpnTunnelsIDParams() *GetVpnTunnelsIDParams {
	var ()
	return &GetVpnTunnelsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpnTunnelsIDParamsWithTimeout creates a new GetVpnTunnelsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpnTunnelsIDParamsWithTimeout(timeout time.Duration) *GetVpnTunnelsIDParams {
	var ()
	return &GetVpnTunnelsIDParams{

		timeout: timeout,
	}
}

// NewGetVpnTunnelsIDParamsWithContext creates a new GetVpnTunnelsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpnTunnelsIDParamsWithContext(ctx context.Context) *GetVpnTunnelsIDParams {
	var ()
	return &GetVpnTunnelsIDParams{

		Context: ctx,
	}
}

// NewGetVpnTunnelsIDParamsWithHTTPClient creates a new GetVpnTunnelsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpnTunnelsIDParamsWithHTTPClient(client *http.Client) *GetVpnTunnelsIDParams {
	var ()
	return &GetVpnTunnelsIDParams{
		HTTPClient: client,
	}
}

/*GetVpnTunnelsIDParams contains all the parameters to send to the API endpoint
for the get vpn tunnels ID operation typically these are written to a http.Request
*/
type GetVpnTunnelsIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) WithTimeout(timeout time.Duration) *GetVpnTunnelsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) WithContext(ctx context.Context) *GetVpnTunnelsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) WithHTTPClient(client *http.Client) *GetVpnTunnelsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) WithID(id int64) *GetVpnTunnelsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpn tunnels ID params
func (o *GetVpnTunnelsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpnTunnelsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}