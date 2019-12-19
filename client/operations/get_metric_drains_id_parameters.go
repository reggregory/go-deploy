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

// NewGetMetricDrainsIDParams creates a new GetMetricDrainsIDParams object
// with the default values initialized.
func NewGetMetricDrainsIDParams() *GetMetricDrainsIDParams {
	var ()
	return &GetMetricDrainsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricDrainsIDParamsWithTimeout creates a new GetMetricDrainsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetricDrainsIDParamsWithTimeout(timeout time.Duration) *GetMetricDrainsIDParams {
	var ()
	return &GetMetricDrainsIDParams{

		timeout: timeout,
	}
}

// NewGetMetricDrainsIDParamsWithContext creates a new GetMetricDrainsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetricDrainsIDParamsWithContext(ctx context.Context) *GetMetricDrainsIDParams {
	var ()
	return &GetMetricDrainsIDParams{

		Context: ctx,
	}
}

// NewGetMetricDrainsIDParamsWithHTTPClient creates a new GetMetricDrainsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetricDrainsIDParamsWithHTTPClient(client *http.Client) *GetMetricDrainsIDParams {
	var ()
	return &GetMetricDrainsIDParams{
		HTTPClient: client,
	}
}

/*GetMetricDrainsIDParams contains all the parameters to send to the API endpoint
for the get metric drains ID operation typically these are written to a http.Request
*/
type GetMetricDrainsIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metric drains ID params
func (o *GetMetricDrainsIDParams) WithTimeout(timeout time.Duration) *GetMetricDrainsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metric drains ID params
func (o *GetMetricDrainsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metric drains ID params
func (o *GetMetricDrainsIDParams) WithContext(ctx context.Context) *GetMetricDrainsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metric drains ID params
func (o *GetMetricDrainsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metric drains ID params
func (o *GetMetricDrainsIDParams) WithHTTPClient(client *http.Client) *GetMetricDrainsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metric drains ID params
func (o *GetMetricDrainsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get metric drains ID params
func (o *GetMetricDrainsIDParams) WithID(id int64) *GetMetricDrainsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get metric drains ID params
func (o *GetMetricDrainsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricDrainsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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