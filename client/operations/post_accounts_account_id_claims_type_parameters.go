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

// NewPostAccountsAccountIDClaimsTypeParams creates a new PostAccountsAccountIDClaimsTypeParams object
// with the default values initialized.
func NewPostAccountsAccountIDClaimsTypeParams() *PostAccountsAccountIDClaimsTypeParams {
	var ()
	return &PostAccountsAccountIDClaimsTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountsAccountIDClaimsTypeParamsWithTimeout creates a new PostAccountsAccountIDClaimsTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountsAccountIDClaimsTypeParamsWithTimeout(timeout time.Duration) *PostAccountsAccountIDClaimsTypeParams {
	var ()
	return &PostAccountsAccountIDClaimsTypeParams{

		timeout: timeout,
	}
}

// NewPostAccountsAccountIDClaimsTypeParamsWithContext creates a new PostAccountsAccountIDClaimsTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAccountsAccountIDClaimsTypeParamsWithContext(ctx context.Context) *PostAccountsAccountIDClaimsTypeParams {
	var ()
	return &PostAccountsAccountIDClaimsTypeParams{

		Context: ctx,
	}
}

// NewPostAccountsAccountIDClaimsTypeParamsWithHTTPClient creates a new PostAccountsAccountIDClaimsTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAccountsAccountIDClaimsTypeParamsWithHTTPClient(client *http.Client) *PostAccountsAccountIDClaimsTypeParams {
	var ()
	return &PostAccountsAccountIDClaimsTypeParams{
		HTTPClient: client,
	}
}

/*PostAccountsAccountIDClaimsTypeParams contains all the parameters to send to the API endpoint
for the post accounts account ID claims type operation typically these are written to a http.Request
*/
type PostAccountsAccountIDClaimsTypeParams struct {

	/*AccountID
	  account_id

	*/
	AccountID int64
	/*AppRequest*/
	AppRequest *models.AppRequest7
	/*Type
	  type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithTimeout(timeout time.Duration) *PostAccountsAccountIDClaimsTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithContext(ctx context.Context) *PostAccountsAccountIDClaimsTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithHTTPClient(client *http.Client) *PostAccountsAccountIDClaimsTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithAccountID(accountID int64) *PostAccountsAccountIDClaimsTypeParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetAccountID(accountID int64) {
	o.AccountID = accountID
}

// WithAppRequest adds the appRequest to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithAppRequest(appRequest *models.AppRequest7) *PostAccountsAccountIDClaimsTypeParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetAppRequest(appRequest *models.AppRequest7) {
	o.AppRequest = appRequest
}

// WithType adds the typeVar to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) WithType(typeVar string) *PostAccountsAccountIDClaimsTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post accounts account ID claims type params
func (o *PostAccountsAccountIDClaimsTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountsAccountIDClaimsTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", swag.FormatInt64(o.AccountID)); err != nil {
		return err
	}

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