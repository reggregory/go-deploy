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

// NewGetContainersIDParams creates a new GetContainersIDParams object
// with the default values initialized.
func NewGetContainersIDParams() *GetContainersIDParams {
	var ()
	return &GetContainersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContainersIDParamsWithTimeout creates a new GetContainersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContainersIDParamsWithTimeout(timeout time.Duration) *GetContainersIDParams {
	var ()
	return &GetContainersIDParams{

		timeout: timeout,
	}
}

// NewGetContainersIDParamsWithContext creates a new GetContainersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContainersIDParamsWithContext(ctx context.Context) *GetContainersIDParams {
	var ()
	return &GetContainersIDParams{

		Context: ctx,
	}
}

// NewGetContainersIDParamsWithHTTPClient creates a new GetContainersIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContainersIDParamsWithHTTPClient(client *http.Client) *GetContainersIDParams {
	var ()
	return &GetContainersIDParams{
		HTTPClient: client,
	}
}

/*GetContainersIDParams contains all the parameters to send to the API endpoint
for the get containers ID operation typically these are written to a http.Request
*/
type GetContainersIDParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get containers ID params
func (o *GetContainersIDParams) WithTimeout(timeout time.Duration) *GetContainersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get containers ID params
func (o *GetContainersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get containers ID params
func (o *GetContainersIDParams) WithContext(ctx context.Context) *GetContainersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get containers ID params
func (o *GetContainersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get containers ID params
func (o *GetContainersIDParams) WithHTTPClient(client *http.Client) *GetContainersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get containers ID params
func (o *GetContainersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get containers ID params
func (o *GetContainersIDParams) WithID(id int64) *GetContainersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get containers ID params
func (o *GetContainersIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetContainersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
