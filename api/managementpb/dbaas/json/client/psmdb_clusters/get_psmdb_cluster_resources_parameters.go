// Code generated by go-swagger; DO NOT EDIT.

package psmdb_clusters

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

// NewGetPSMDBClusterResourcesParams creates a new GetPSMDBClusterResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPSMDBClusterResourcesParams() *GetPSMDBClusterResourcesParams {
	return &GetPSMDBClusterResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPSMDBClusterResourcesParamsWithTimeout creates a new GetPSMDBClusterResourcesParams object
// with the ability to set a timeout on a request.
func NewGetPSMDBClusterResourcesParamsWithTimeout(timeout time.Duration) *GetPSMDBClusterResourcesParams {
	return &GetPSMDBClusterResourcesParams{
		timeout: timeout,
	}
}

// NewGetPSMDBClusterResourcesParamsWithContext creates a new GetPSMDBClusterResourcesParams object
// with the ability to set a context for a request.
func NewGetPSMDBClusterResourcesParamsWithContext(ctx context.Context) *GetPSMDBClusterResourcesParams {
	return &GetPSMDBClusterResourcesParams{
		Context: ctx,
	}
}

// NewGetPSMDBClusterResourcesParamsWithHTTPClient creates a new GetPSMDBClusterResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPSMDBClusterResourcesParamsWithHTTPClient(client *http.Client) *GetPSMDBClusterResourcesParams {
	return &GetPSMDBClusterResourcesParams{
		HTTPClient: client,
	}
}

/* GetPSMDBClusterResourcesParams contains all the parameters to send to the API endpoint
   for the get PSMDB cluster resources operation.

   Typically these are written to a http.Request.
*/
type GetPSMDBClusterResourcesParams struct {
	// Body.
	Body GetPSMDBClusterResourcesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get PSMDB cluster resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPSMDBClusterResourcesParams) WithDefaults() *GetPSMDBClusterResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get PSMDB cluster resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPSMDBClusterResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) WithTimeout(timeout time.Duration) *GetPSMDBClusterResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) WithContext(ctx context.Context) *GetPSMDBClusterResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) WithHTTPClient(client *http.Client) *GetPSMDBClusterResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) WithBody(body GetPSMDBClusterResourcesBody) *GetPSMDBClusterResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get PSMDB cluster resources params
func (o *GetPSMDBClusterResourcesParams) SetBody(body GetPSMDBClusterResourcesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetPSMDBClusterResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
