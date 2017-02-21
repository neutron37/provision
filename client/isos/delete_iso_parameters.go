package isos

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

// NewDeleteIsoParams creates a new DeleteIsoParams object
// with the default values initialized.
func NewDeleteIsoParams() *DeleteIsoParams {
	var ()
	return &DeleteIsoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIsoParamsWithTimeout creates a new DeleteIsoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIsoParamsWithTimeout(timeout time.Duration) *DeleteIsoParams {
	var ()
	return &DeleteIsoParams{

		timeout: timeout,
	}
}

// NewDeleteIsoParamsWithContext creates a new DeleteIsoParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIsoParamsWithContext(ctx context.Context) *DeleteIsoParams {
	var ()
	return &DeleteIsoParams{

		Context: ctx,
	}
}

// NewDeleteIsoParamsWithHTTPClient creates a new DeleteIsoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIsoParamsWithHTTPClient(client *http.Client) *DeleteIsoParams {
	var ()
	return &DeleteIsoParams{
		HTTPClient: client,
	}
}

/*DeleteIsoParams contains all the parameters to send to the API endpoint
for the delete iso operation typically these are written to a http.Request
*/
type DeleteIsoParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete iso params
func (o *DeleteIsoParams) WithTimeout(timeout time.Duration) *DeleteIsoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iso params
func (o *DeleteIsoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iso params
func (o *DeleteIsoParams) WithContext(ctx context.Context) *DeleteIsoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iso params
func (o *DeleteIsoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iso params
func (o *DeleteIsoParams) WithHTTPClient(client *http.Client) *DeleteIsoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iso params
func (o *DeleteIsoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete iso params
func (o *DeleteIsoParams) WithName(name string) *DeleteIsoParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete iso params
func (o *DeleteIsoParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIsoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}