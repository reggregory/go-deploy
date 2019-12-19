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

	models "github.com/aptible/go-deploy/models"
)

// NewPostOperationsOperationIDSSHPortalConnectionsParams creates a new PostOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized.
func NewPostOperationsOperationIDSSHPortalConnectionsParams() *PostOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &PostOperationsOperationIDSSHPortalConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOperationsOperationIDSSHPortalConnectionsParamsWithTimeout creates a new PostOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOperationsOperationIDSSHPortalConnectionsParamsWithTimeout(timeout time.Duration) *PostOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &PostOperationsOperationIDSSHPortalConnectionsParams{

		timeout: timeout,
	}
}

// NewPostOperationsOperationIDSSHPortalConnectionsParamsWithContext creates a new PostOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOperationsOperationIDSSHPortalConnectionsParamsWithContext(ctx context.Context) *PostOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &PostOperationsOperationIDSSHPortalConnectionsParams{

		Context: ctx,
	}
}

// NewPostOperationsOperationIDSSHPortalConnectionsParamsWithHTTPClient creates a new PostOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOperationsOperationIDSSHPortalConnectionsParamsWithHTTPClient(client *http.Client) *PostOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &PostOperationsOperationIDSSHPortalConnectionsParams{
		HTTPClient: client,
	}
}

/*PostOperationsOperationIDSSHPortalConnectionsParams contains all the parameters to send to the API endpoint
for the post operations operation ID SSH portal connections operation typically these are written to a http.Request
*/
type PostOperationsOperationIDSSHPortalConnectionsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest32
	/*OperationID
	  operation_id

	*/
	OperationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WithTimeout(timeout time.Duration) *PostOperationsOperationIDSSHPortalConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WithContext(ctx context.Context) *PostOperationsOperationIDSSHPortalConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WithHTTPClient(client *http.Client) *PostOperationsOperationIDSSHPortalConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WithAppRequest(appRequest *models.AppRequest32) *PostOperationsOperationIDSSHPortalConnectionsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) SetAppRequest(appRequest *models.AppRequest32) {
	o.AppRequest = appRequest
}

// WithOperationID adds the operationID to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WithOperationID(operationID int64) *PostOperationsOperationIDSSHPortalConnectionsParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the post operations operation ID SSH portal connections params
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) SetOperationID(operationID int64) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOperationsOperationIDSSHPortalConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param operation_id
	if err := r.SetPathParam("operation_id", swag.FormatInt64(o.OperationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}