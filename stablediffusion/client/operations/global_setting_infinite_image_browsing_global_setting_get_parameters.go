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

// NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParams creates a new GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParams() *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	return &GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithTimeout creates a new GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams object
// with the ability to set a timeout on a request.
func NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithTimeout(timeout time.Duration) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	return &GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams{
		timeout: timeout,
	}
}

// NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithContext creates a new GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams object
// with the ability to set a context for a request.
func NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithContext(ctx context.Context) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	return &GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams{
		Context: ctx,
	}
}

// NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithHTTPClient creates a new GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalSettingInfiniteImageBrowsingGlobalSettingGetParamsWithHTTPClient(client *http.Client) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	return &GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams{
		HTTPClient: client,
	}
}

/*
GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams contains all the parameters to send to the API endpoint

	for the global setting infinite image browsing global setting get operation.

	Typically these are written to a http.Request.
*/
type GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global setting infinite image browsing global setting get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) WithDefaults() *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global setting infinite image browsing global setting get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) WithTimeout(timeout time.Duration) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) WithContext(ctx context.Context) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) WithHTTPClient(client *http.Client) *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global setting infinite image browsing global setting get params
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalSettingInfiniteImageBrowsingGlobalSettingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
