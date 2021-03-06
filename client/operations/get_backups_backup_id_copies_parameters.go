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

// NewGetBackupsBackupIDCopiesParams creates a new GetBackupsBackupIDCopiesParams object
// with the default values initialized.
func NewGetBackupsBackupIDCopiesParams() *GetBackupsBackupIDCopiesParams {
	var ()
	return &GetBackupsBackupIDCopiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupsBackupIDCopiesParamsWithTimeout creates a new GetBackupsBackupIDCopiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBackupsBackupIDCopiesParamsWithTimeout(timeout time.Duration) *GetBackupsBackupIDCopiesParams {
	var ()
	return &GetBackupsBackupIDCopiesParams{

		timeout: timeout,
	}
}

// NewGetBackupsBackupIDCopiesParamsWithContext creates a new GetBackupsBackupIDCopiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBackupsBackupIDCopiesParamsWithContext(ctx context.Context) *GetBackupsBackupIDCopiesParams {
	var ()
	return &GetBackupsBackupIDCopiesParams{

		Context: ctx,
	}
}

// NewGetBackupsBackupIDCopiesParamsWithHTTPClient creates a new GetBackupsBackupIDCopiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBackupsBackupIDCopiesParamsWithHTTPClient(client *http.Client) *GetBackupsBackupIDCopiesParams {
	var ()
	return &GetBackupsBackupIDCopiesParams{
		HTTPClient: client,
	}
}

/*GetBackupsBackupIDCopiesParams contains all the parameters to send to the API endpoint
for the get backups backup ID copies operation typically these are written to a http.Request
*/
type GetBackupsBackupIDCopiesParams struct {

	/*BackupID
	  backup_id

	*/
	BackupID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) WithTimeout(timeout time.Duration) *GetBackupsBackupIDCopiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) WithContext(ctx context.Context) *GetBackupsBackupIDCopiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) WithHTTPClient(client *http.Client) *GetBackupsBackupIDCopiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) WithBackupID(backupID int64) *GetBackupsBackupIDCopiesParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) SetBackupID(backupID int64) {
	o.BackupID = backupID
}

// WithPage adds the page to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) WithPage(page *int64) *GetBackupsBackupIDCopiesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get backups backup ID copies params
func (o *GetBackupsBackupIDCopiesParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupsBackupIDCopiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backup_id
	if err := r.SetPathParam("backup_id", swag.FormatInt64(o.BackupID)); err != nil {
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
