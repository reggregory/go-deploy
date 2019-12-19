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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPostClaimsTypeParams creates a new PostClaimsTypeParams object
// with the default values initialized.
func NewPostClaimsTypeParams() *PostClaimsTypeParams {
	var ()
	return &PostClaimsTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClaimsTypeParamsWithTimeout creates a new PostClaimsTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClaimsTypeParamsWithTimeout(timeout time.Duration) *PostClaimsTypeParams {
	var ()
	return &PostClaimsTypeParams{

		timeout: timeout,
	}
}

// NewPostClaimsTypeParamsWithContext creates a new PostClaimsTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClaimsTypeParamsWithContext(ctx context.Context) *PostClaimsTypeParams {
	var ()
	return &PostClaimsTypeParams{

		Context: ctx,
	}
}

// NewPostClaimsTypeParamsWithHTTPClient creates a new PostClaimsTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClaimsTypeParamsWithHTTPClient(client *http.Client) *PostClaimsTypeParams {
	var ()
	return &PostClaimsTypeParams{
		HTTPClient: client,
	}
}

/*PostClaimsTypeParams contains all the parameters to send to the API endpoint
for the post claims type operation typically these are written to a http.Request
*/
type PostClaimsTypeParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest9
	/*Type
	  type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post claims type params
func (o *PostClaimsTypeParams) WithTimeout(timeout time.Duration) *PostClaimsTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post claims type params
func (o *PostClaimsTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post claims type params
func (o *PostClaimsTypeParams) WithContext(ctx context.Context) *PostClaimsTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post claims type params
func (o *PostClaimsTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post claims type params
func (o *PostClaimsTypeParams) WithHTTPClient(client *http.Client) *PostClaimsTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post claims type params
func (o *PostClaimsTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post claims type params
func (o *PostClaimsTypeParams) WithAppRequest(appRequest *models.AppRequest9) *PostClaimsTypeParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post claims type params
func (o *PostClaimsTypeParams) SetAppRequest(appRequest *models.AppRequest9) {
	o.AppRequest = appRequest
}

// WithType adds the typeVar to the post claims type params
func (o *PostClaimsTypeParams) WithType(typeVar string) *PostClaimsTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post claims type params
func (o *PostClaimsTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostClaimsTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}