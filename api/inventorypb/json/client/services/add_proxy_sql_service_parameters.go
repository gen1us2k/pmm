// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAddProxySQLServiceParams creates a new AddProxySQLServiceParams object
// with the default values initialized.
func NewAddProxySQLServiceParams() *AddProxySQLServiceParams {
	var ()
	return &AddProxySQLServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddProxySQLServiceParamsWithTimeout creates a new AddProxySQLServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddProxySQLServiceParamsWithTimeout(timeout time.Duration) *AddProxySQLServiceParams {
	var ()
	return &AddProxySQLServiceParams{

		timeout: timeout,
	}
}

// NewAddProxySQLServiceParamsWithContext creates a new AddProxySQLServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddProxySQLServiceParamsWithContext(ctx context.Context) *AddProxySQLServiceParams {
	var ()
	return &AddProxySQLServiceParams{

		Context: ctx,
	}
}

// NewAddProxySQLServiceParamsWithHTTPClient creates a new AddProxySQLServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddProxySQLServiceParamsWithHTTPClient(client *http.Client) *AddProxySQLServiceParams {
	var ()
	return &AddProxySQLServiceParams{
		HTTPClient: client,
	}
}

/*AddProxySQLServiceParams contains all the parameters to send to the API endpoint
for the add proxy SQL service operation typically these are written to a http.Request
*/
type AddProxySQLServiceParams struct {

	/*Body*/
	Body AddProxySQLServiceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add proxy SQL service params
func (o *AddProxySQLServiceParams) WithTimeout(timeout time.Duration) *AddProxySQLServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add proxy SQL service params
func (o *AddProxySQLServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add proxy SQL service params
func (o *AddProxySQLServiceParams) WithContext(ctx context.Context) *AddProxySQLServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add proxy SQL service params
func (o *AddProxySQLServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add proxy SQL service params
func (o *AddProxySQLServiceParams) WithHTTPClient(client *http.Client) *AddProxySQLServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add proxy SQL service params
func (o *AddProxySQLServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add proxy SQL service params
func (o *AddProxySQLServiceParams) WithBody(body AddProxySQLServiceBody) *AddProxySQLServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add proxy SQL service params
func (o *AddProxySQLServiceParams) SetBody(body AddProxySQLServiceBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddProxySQLServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
