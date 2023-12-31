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
	"github.com/go-openapi/strfmt"
)

// NewGetCurrentUserUserGetParams creates a new GetCurrentUserUserGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCurrentUserUserGetParams() *GetCurrentUserUserGetParams {
	return &GetCurrentUserUserGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentUserUserGetParamsWithTimeout creates a new GetCurrentUserUserGetParams object
// with the ability to set a timeout on a request.
func NewGetCurrentUserUserGetParamsWithTimeout(timeout time.Duration) *GetCurrentUserUserGetParams {
	return &GetCurrentUserUserGetParams{
		timeout: timeout,
	}
}

// NewGetCurrentUserUserGetParamsWithContext creates a new GetCurrentUserUserGetParams object
// with the ability to set a context for a request.
func NewGetCurrentUserUserGetParamsWithContext(ctx context.Context) *GetCurrentUserUserGetParams {
	return &GetCurrentUserUserGetParams{
		Context: ctx,
	}
}

// NewGetCurrentUserUserGetParamsWithHTTPClient creates a new GetCurrentUserUserGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCurrentUserUserGetParamsWithHTTPClient(client *http.Client) *GetCurrentUserUserGetParams {
	return &GetCurrentUserUserGetParams{
		HTTPClient: client,
	}
}

/*
GetCurrentUserUserGetParams contains all the parameters to send to the API endpoint

	for the get current user user get operation.

	Typically these are written to a http.Request.
*/
type GetCurrentUserUserGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get current user user get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserUserGetParams) WithDefaults() *GetCurrentUserUserGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get current user user get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserUserGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get current user user get params
func (o *GetCurrentUserUserGetParams) WithTimeout(timeout time.Duration) *GetCurrentUserUserGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current user user get params
func (o *GetCurrentUserUserGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current user user get params
func (o *GetCurrentUserUserGetParams) WithContext(ctx context.Context) *GetCurrentUserUserGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current user user get params
func (o *GetCurrentUserUserGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current user user get params
func (o *GetCurrentUserUserGetParams) WithHTTPClient(client *http.Client) *GetCurrentUserUserGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current user user get params
func (o *GetCurrentUserUserGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentUserUserGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
