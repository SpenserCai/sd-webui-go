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

// NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams creates a new APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams() *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	return &APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithTimeout creates a new APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams object
// with the ability to set a timeout on a request.
func NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithTimeout(timeout time.Duration) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	return &APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams{
		timeout: timeout,
	}
}

// NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithContext creates a new APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams object
// with the ability to set a context for a request.
func NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithContext(ctx context.Context) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	return &APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams{
		Context: ctx,
	}
}

// NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithHTTPClient creates a new APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPISetSendImgPathInfiniteImageBrowsingSendImgPathPostParamsWithHTTPClient(client *http.Client) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	return &APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams{
		HTTPClient: client,
	}
}

/*
APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams contains all the parameters to send to the API endpoint

	for the api set send img path infinite image browsing send img path post operation.

	Typically these are written to a http.Request.
*/
type APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams struct {

	// Path.
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the api set send img path infinite image browsing send img path post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WithDefaults() *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the api set send img path infinite image browsing send img path post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WithTimeout(timeout time.Duration) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WithContext(ctx context.Context) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WithHTTPClient(client *http.Client) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WithPath(path string) *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the api set send img path infinite image browsing send img path post params
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *APISetSendImgPathInfiniteImageBrowsingSendImgPathPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
