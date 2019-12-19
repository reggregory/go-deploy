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

// NewGetLogDrainsLogDrainIDContainersParams creates a new GetLogDrainsLogDrainIDContainersParams object
// with the default values initialized.
func NewGetLogDrainsLogDrainIDContainersParams() *GetLogDrainsLogDrainIDContainersParams {
	var ()
	return &GetLogDrainsLogDrainIDContainersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogDrainsLogDrainIDContainersParamsWithTimeout creates a new GetLogDrainsLogDrainIDContainersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogDrainsLogDrainIDContainersParamsWithTimeout(timeout time.Duration) *GetLogDrainsLogDrainIDContainersParams {
	var ()
	return &GetLogDrainsLogDrainIDContainersParams{

		timeout: timeout,
	}
}

// NewGetLogDrainsLogDrainIDContainersParamsWithContext creates a new GetLogDrainsLogDrainIDContainersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogDrainsLogDrainIDContainersParamsWithContext(ctx context.Context) *GetLogDrainsLogDrainIDContainersParams {
	var ()
	return &GetLogDrainsLogDrainIDContainersParams{

		Context: ctx,
	}
}

// NewGetLogDrainsLogDrainIDContainersParamsWithHTTPClient creates a new GetLogDrainsLogDrainIDContainersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogDrainsLogDrainIDContainersParamsWithHTTPClient(client *http.Client) *GetLogDrainsLogDrainIDContainersParams {
	var ()
	return &GetLogDrainsLogDrainIDContainersParams{
		HTTPClient: client,
	}
}

/*GetLogDrainsLogDrainIDContainersParams contains all the parameters to send to the API endpoint
for the get log drains log drain ID containers operation typically these are written to a http.Request
*/
type GetLogDrainsLogDrainIDContainersParams struct {

	/*LogDrainID
	  log_drain_id

	*/
	LogDrainID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) WithTimeout(timeout time.Duration) *GetLogDrainsLogDrainIDContainersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) WithContext(ctx context.Context) *GetLogDrainsLogDrainIDContainersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) WithHTTPClient(client *http.Client) *GetLogDrainsLogDrainIDContainersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogDrainID adds the logDrainID to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) WithLogDrainID(logDrainID int64) *GetLogDrainsLogDrainIDContainersParams {
	o.SetLogDrainID(logDrainID)
	return o
}

// SetLogDrainID adds the logDrainId to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) SetLogDrainID(logDrainID int64) {
	o.LogDrainID = logDrainID
}

// WithPage adds the page to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) WithPage(page *int64) *GetLogDrainsLogDrainIDContainersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get log drains log drain ID containers params
func (o *GetLogDrainsLogDrainIDContainersParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogDrainsLogDrainIDContainersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param log_drain_id
	if err := r.SetPathParam("log_drain_id", swag.FormatInt64(o.LogDrainID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
