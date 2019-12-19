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

// NewGetAccountsAccountIDLogDrainsParams creates a new GetAccountsAccountIDLogDrainsParams object
// with the default values initialized.
func NewGetAccountsAccountIDLogDrainsParams() *GetAccountsAccountIDLogDrainsParams {
	var ()
	return &GetAccountsAccountIDLogDrainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDLogDrainsParamsWithTimeout creates a new GetAccountsAccountIDLogDrainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsAccountIDLogDrainsParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDLogDrainsParams {
	var ()
	return &GetAccountsAccountIDLogDrainsParams{

		timeout: timeout,
	}
}

// NewGetAccountsAccountIDLogDrainsParamsWithContext creates a new GetAccountsAccountIDLogDrainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsAccountIDLogDrainsParamsWithContext(ctx context.Context) *GetAccountsAccountIDLogDrainsParams {
	var ()
	return &GetAccountsAccountIDLogDrainsParams{

		Context: ctx,
	}
}

// NewGetAccountsAccountIDLogDrainsParamsWithHTTPClient creates a new GetAccountsAccountIDLogDrainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsAccountIDLogDrainsParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDLogDrainsParams {
	var ()
	return &GetAccountsAccountIDLogDrainsParams{
		HTTPClient: client,
	}
}

/*GetAccountsAccountIDLogDrainsParams contains all the parameters to send to the API endpoint
for the get accounts account ID log drains operation typically these are written to a http.Request
*/
type GetAccountsAccountIDLogDrainsParams struct {

	/*AccountID
	  account_id

	*/
	AccountID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDLogDrainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) WithContext(ctx context.Context) *GetAccountsAccountIDLogDrainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDLogDrainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) WithAccountID(accountID int64) *GetAccountsAccountIDLogDrainsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithPage adds the page to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) WithPage(page *int64) *GetAccountsAccountIDLogDrainsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get accounts account ID log drains params
func (o *GetAccountsAccountIDLogDrainsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDLogDrainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", swag.FormatInt64(o.AccountID)); err != nil {
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
