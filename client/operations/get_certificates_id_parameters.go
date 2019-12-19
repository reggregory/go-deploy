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

// NewGetCertificatesIDParams creates a new GetCertificatesIDParams object
// with the default values initialized.
func NewGetCertificatesIDParams() *GetCertificatesIDParams {
	var ()
	return &GetCertificatesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificatesIDParamsWithTimeout creates a new GetCertificatesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificatesIDParamsWithTimeout(timeout time.Duration) *GetCertificatesIDParams {
	var ()
	return &GetCertificatesIDParams{

		timeout: timeout,
	}
}

// NewGetCertificatesIDParamsWithContext creates a new GetCertificatesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificatesIDParamsWithContext(ctx context.Context) *GetCertificatesIDParams {
	var ()
	return &GetCertificatesIDParams{

		Context: ctx,
	}
}

// NewGetCertificatesIDParamsWithHTTPClient creates a new GetCertificatesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificatesIDParamsWithHTTPClient(client *http.Client) *GetCertificatesIDParams {
	var ()
	return &GetCertificatesIDParams{
		HTTPClient: client,
	}
}

/*GetCertificatesIDParams contains all the parameters to send to the API endpoint
for the get certificates ID operation typically these are written to a http.Request
*/
type GetCertificatesIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get certificates ID params
func (o *GetCertificatesIDParams) WithTimeout(timeout time.Duration) *GetCertificatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificates ID params
func (o *GetCertificatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificates ID params
func (o *GetCertificatesIDParams) WithContext(ctx context.Context) *GetCertificatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificates ID params
func (o *GetCertificatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificates ID params
func (o *GetCertificatesIDParams) WithHTTPClient(client *http.Client) *GetCertificatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificates ID params
func (o *GetCertificatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificates ID params
func (o *GetCertificatesIDParams) WithID(id int64) *GetCertificatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificates ID params
func (o *GetCertificatesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
