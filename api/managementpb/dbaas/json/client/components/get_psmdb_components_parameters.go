// Code generated by go-swagger; DO NOT EDIT.

package components

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

// NewGetPSMDBComponentsParams creates a new GetPSMDBComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPSMDBComponentsParams() *GetPSMDBComponentsParams {
	return &GetPSMDBComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPSMDBComponentsParamsWithTimeout creates a new GetPSMDBComponentsParams object
// with the ability to set a timeout on a request.
func NewGetPSMDBComponentsParamsWithTimeout(timeout time.Duration) *GetPSMDBComponentsParams {
	return &GetPSMDBComponentsParams{
		timeout: timeout,
	}
}

// NewGetPSMDBComponentsParamsWithContext creates a new GetPSMDBComponentsParams object
// with the ability to set a context for a request.
func NewGetPSMDBComponentsParamsWithContext(ctx context.Context) *GetPSMDBComponentsParams {
	return &GetPSMDBComponentsParams{
		Context: ctx,
	}
}

// NewGetPSMDBComponentsParamsWithHTTPClient creates a new GetPSMDBComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPSMDBComponentsParamsWithHTTPClient(client *http.Client) *GetPSMDBComponentsParams {
	return &GetPSMDBComponentsParams{
		HTTPClient: client,
	}
}

/* GetPSMDBComponentsParams contains all the parameters to send to the API endpoint
   for the get PSMDB components operation.

   Typically these are written to a http.Request.
*/
type GetPSMDBComponentsParams struct {
	// Body.
	Body GetPSMDBComponentsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get PSMDB components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPSMDBComponentsParams) WithDefaults() *GetPSMDBComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get PSMDB components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPSMDBComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get PSMDB components params
func (o *GetPSMDBComponentsParams) WithTimeout(timeout time.Duration) *GetPSMDBComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get PSMDB components params
func (o *GetPSMDBComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get PSMDB components params
func (o *GetPSMDBComponentsParams) WithContext(ctx context.Context) *GetPSMDBComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get PSMDB components params
func (o *GetPSMDBComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get PSMDB components params
func (o *GetPSMDBComponentsParams) WithHTTPClient(client *http.Client) *GetPSMDBComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get PSMDB components params
func (o *GetPSMDBComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get PSMDB components params
func (o *GetPSMDBComponentsParams) WithBody(body GetPSMDBComponentsBody) *GetPSMDBComponentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get PSMDB components params
func (o *GetPSMDBComponentsParams) SetBody(body GetPSMDBComponentsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetPSMDBComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
