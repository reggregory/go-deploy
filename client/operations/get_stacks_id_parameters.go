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

// NewGetStacksIDParams creates a new GetStacksIDParams object
// with the default values initialized.
func NewGetStacksIDParams() *GetStacksIDParams {
	var ()
	return &GetStacksIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStacksIDParamsWithTimeout creates a new GetStacksIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStacksIDParamsWithTimeout(timeout time.Duration) *GetStacksIDParams {
	var ()
	return &GetStacksIDParams{

		timeout: timeout,
	}
}

// NewGetStacksIDParamsWithContext creates a new GetStacksIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStacksIDParamsWithContext(ctx context.Context) *GetStacksIDParams {
	var ()
	return &GetStacksIDParams{

		Context: ctx,
	}
}

// NewGetStacksIDParamsWithHTTPClient creates a new GetStacksIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStacksIDParamsWithHTTPClient(client *http.Client) *GetStacksIDParams {
	var ()
	return &GetStacksIDParams{
		HTTPClient: client,
	}
}

/*GetStacksIDParams contains all the parameters to send to the API endpoint
for the get stacks ID operation typically these are written to a http.Request
*/
type GetStacksIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stacks ID params
func (o *GetStacksIDParams) WithTimeout(timeout time.Duration) *GetStacksIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stacks ID params
func (o *GetStacksIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stacks ID params
func (o *GetStacksIDParams) WithContext(ctx context.Context) *GetStacksIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stacks ID params
func (o *GetStacksIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stacks ID params
func (o *GetStacksIDParams) WithHTTPClient(client *http.Client) *GetStacksIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stacks ID params
func (o *GetStacksIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get stacks ID params
func (o *GetStacksIDParams) WithID(id int64) *GetStacksIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stacks ID params
func (o *GetStacksIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStacksIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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