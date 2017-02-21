package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFileParams creates a new DeleteFileParams object
// with the default values initialized.
func NewDeleteFileParams() *DeleteFileParams {
	var ()
	return &DeleteFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFileParamsWithTimeout creates a new DeleteFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFileParamsWithTimeout(timeout time.Duration) *DeleteFileParams {
	var ()
	return &DeleteFileParams{

		timeout: timeout,
	}
}

// NewDeleteFileParamsWithContext creates a new DeleteFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFileParamsWithContext(ctx context.Context) *DeleteFileParams {
	var ()
	return &DeleteFileParams{

		Context: ctx,
	}
}

// NewDeleteFileParamsWithHTTPClient creates a new DeleteFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFileParamsWithHTTPClient(client *http.Client) *DeleteFileParams {
	var ()
	return &DeleteFileParams{
		HTTPClient: client,
	}
}

/*DeleteFileParams contains all the parameters to send to the API endpoint
for the delete file operation typically these are written to a http.Request
*/
type DeleteFileParams struct {

	/*Path*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete file params
func (o *DeleteFileParams) WithTimeout(timeout time.Duration) *DeleteFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete file params
func (o *DeleteFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete file params
func (o *DeleteFileParams) WithContext(ctx context.Context) *DeleteFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete file params
func (o *DeleteFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete file params
func (o *DeleteFileParams) WithHTTPClient(client *http.Client) *DeleteFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete file params
func (o *DeleteFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the delete file params
func (o *DeleteFileParams) WithPath(path string) *DeleteFileParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete file params
func (o *DeleteFileParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}