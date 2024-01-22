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

	"github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

// NewSamAreaToolsSamAreaPostParams creates a new SamAreaToolsSamAreaPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSamAreaToolsSamAreaPostParams() *SamAreaToolsSamAreaPostParams {
	return &SamAreaToolsSamAreaPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSamAreaToolsSamAreaPostParamsWithTimeout creates a new SamAreaToolsSamAreaPostParams object
// with the ability to set a timeout on a request.
func NewSamAreaToolsSamAreaPostParamsWithTimeout(timeout time.Duration) *SamAreaToolsSamAreaPostParams {
	return &SamAreaToolsSamAreaPostParams{
		timeout: timeout,
	}
}

// NewSamAreaToolsSamAreaPostParamsWithContext creates a new SamAreaToolsSamAreaPostParams object
// with the ability to set a context for a request.
func NewSamAreaToolsSamAreaPostParamsWithContext(ctx context.Context) *SamAreaToolsSamAreaPostParams {
	return &SamAreaToolsSamAreaPostParams{
		Context: ctx,
	}
}

// NewSamAreaToolsSamAreaPostParamsWithHTTPClient creates a new SamAreaToolsSamAreaPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSamAreaToolsSamAreaPostParamsWithHTTPClient(client *http.Client) *SamAreaToolsSamAreaPostParams {
	return &SamAreaToolsSamAreaPostParams{
		HTTPClient: client,
	}
}

/*
SamAreaToolsSamAreaPostParams contains all the parameters to send to the API endpoint

	for the sam area tools sam area post operation.

	Typically these are written to a http.Request.
*/
type SamAreaToolsSamAreaPostParams struct {

	// Body.
	Body *models.SamAreaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sam area tools sam area post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SamAreaToolsSamAreaPostParams) WithDefaults() *SamAreaToolsSamAreaPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sam area tools sam area post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SamAreaToolsSamAreaPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) WithTimeout(timeout time.Duration) *SamAreaToolsSamAreaPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) WithContext(ctx context.Context) *SamAreaToolsSamAreaPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) WithHTTPClient(client *http.Client) *SamAreaToolsSamAreaPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) WithBody(body *models.SamAreaRequest) *SamAreaToolsSamAreaPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the sam area tools sam area post params
func (o *SamAreaToolsSamAreaPostParams) SetBody(body *models.SamAreaRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SamAreaToolsSamAreaPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
