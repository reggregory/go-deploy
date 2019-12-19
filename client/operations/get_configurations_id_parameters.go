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

// NewGetConfigurationsIDParams creates a new GetConfigurationsIDParams object
// with the default values initialized.
func NewGetConfigurationsIDParams() *GetConfigurationsIDParams {
	var ()
	return &GetConfigurationsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationsIDParamsWithTimeout creates a new GetConfigurationsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigurationsIDParamsWithTimeout(timeout time.Duration) *GetConfigurationsIDParams {
	var ()
	return &GetConfigurationsIDParams{

		timeout: timeout,
	}
}

// NewGetConfigurationsIDParamsWithContext creates a new GetConfigurationsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigurationsIDParamsWithContext(ctx context.Context) *GetConfigurationsIDParams {
	var ()
	return &GetConfigurationsIDParams{

		Context: ctx,
	}
}

// NewGetConfigurationsIDParamsWithHTTPClient creates a new GetConfigurationsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigurationsIDParamsWithHTTPClient(client *http.Client) *GetConfigurationsIDParams {
	var ()
	return &GetConfigurationsIDParams{
		HTTPClient: client,
	}
}

/*GetConfigurationsIDParams contains all the parameters to send to the API endpoint
for the get configurations ID operation typically these are written to a http.Request
*/
type GetConfigurationsIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get configurations ID params
func (o *GetConfigurationsIDParams) WithTimeout(timeout time.Duration) *GetConfigurationsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configurations ID params
func (o *GetConfigurationsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configurations ID params
func (o *GetConfigurationsIDParams) WithContext(ctx context.Context) *GetConfigurationsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configurations ID params
func (o *GetConfigurationsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configurations ID params
func (o *GetConfigurationsIDParams) WithHTTPClient(client *http.Client) *GetConfigurationsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configurations ID params
func (o *GetConfigurationsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get configurations ID params
func (o *GetConfigurationsIDParams) WithID(id int64) *GetConfigurationsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get configurations ID params
func (o *GetConfigurationsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
