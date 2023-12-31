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

// NewFetchFileSdExtraNetworksThumbGetParams creates a new FetchFileSdExtraNetworksThumbGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFetchFileSdExtraNetworksThumbGetParams() *FetchFileSdExtraNetworksThumbGetParams {
	return &FetchFileSdExtraNetworksThumbGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFetchFileSdExtraNetworksThumbGetParamsWithTimeout creates a new FetchFileSdExtraNetworksThumbGetParams object
// with the ability to set a timeout on a request.
func NewFetchFileSdExtraNetworksThumbGetParamsWithTimeout(timeout time.Duration) *FetchFileSdExtraNetworksThumbGetParams {
	return &FetchFileSdExtraNetworksThumbGetParams{
		timeout: timeout,
	}
}

// NewFetchFileSdExtraNetworksThumbGetParamsWithContext creates a new FetchFileSdExtraNetworksThumbGetParams object
// with the ability to set a context for a request.
func NewFetchFileSdExtraNetworksThumbGetParamsWithContext(ctx context.Context) *FetchFileSdExtraNetworksThumbGetParams {
	return &FetchFileSdExtraNetworksThumbGetParams{
		Context: ctx,
	}
}

// NewFetchFileSdExtraNetworksThumbGetParamsWithHTTPClient creates a new FetchFileSdExtraNetworksThumbGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFetchFileSdExtraNetworksThumbGetParamsWithHTTPClient(client *http.Client) *FetchFileSdExtraNetworksThumbGetParams {
	return &FetchFileSdExtraNetworksThumbGetParams{
		HTTPClient: client,
	}
}

/*
FetchFileSdExtraNetworksThumbGetParams contains all the parameters to send to the API endpoint

	for the fetch file sd extra networks thumb get operation.

	Typically these are written to a http.Request.
*/
type FetchFileSdExtraNetworksThumbGetParams struct {

	// Filename.
	Filename *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fetch file sd extra networks thumb get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchFileSdExtraNetworksThumbGetParams) WithDefaults() *FetchFileSdExtraNetworksThumbGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fetch file sd extra networks thumb get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchFileSdExtraNetworksThumbGetParams) SetDefaults() {
	var (
		filenameDefault = string("")
	)

	val := FetchFileSdExtraNetworksThumbGetParams{
		Filename: &filenameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) WithTimeout(timeout time.Duration) *FetchFileSdExtraNetworksThumbGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) WithContext(ctx context.Context) *FetchFileSdExtraNetworksThumbGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) WithHTTPClient(client *http.Client) *FetchFileSdExtraNetworksThumbGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) WithFilename(filename *string) *FetchFileSdExtraNetworksThumbGetParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the fetch file sd extra networks thumb get params
func (o *FetchFileSdExtraNetworksThumbGetParams) SetFilename(filename *string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *FetchFileSdExtraNetworksThumbGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filename != nil {

		// query param filename
		var qrFilename string

		if o.Filename != nil {
			qrFilename = *o.Filename
		}
		qFilename := qrFilename
		if qFilename != "" {

			if err := r.SetQueryParam("filename", qFilename); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
