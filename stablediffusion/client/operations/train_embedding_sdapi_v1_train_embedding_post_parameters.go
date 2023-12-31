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

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostParams creates a new TrainEmbeddingSdapiV1TrainEmbeddingPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostParams() *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithTimeout creates a new TrainEmbeddingSdapiV1TrainEmbeddingPostParams object
// with the ability to set a timeout on a request.
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithTimeout(timeout time.Duration) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostParams{
		timeout: timeout,
	}
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithContext creates a new TrainEmbeddingSdapiV1TrainEmbeddingPostParams object
// with the ability to set a context for a request.
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithContext(ctx context.Context) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostParams{
		Context: ctx,
	}
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithHTTPClient creates a new TrainEmbeddingSdapiV1TrainEmbeddingPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostParamsWithHTTPClient(client *http.Client) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostParams{
		HTTPClient: client,
	}
}

/*
TrainEmbeddingSdapiV1TrainEmbeddingPostParams contains all the parameters to send to the API endpoint

	for the train embedding sdapi v1 train embedding post operation.

	Typically these are written to a http.Request.
*/
type TrainEmbeddingSdapiV1TrainEmbeddingPostParams struct {

	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the train embedding sdapi v1 train embedding post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WithDefaults() *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the train embedding sdapi v1 train embedding post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WithTimeout(timeout time.Duration) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WithContext(ctx context.Context) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WithHTTPClient(client *http.Client) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WithBody(body interface{}) *TrainEmbeddingSdapiV1TrainEmbeddingPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the train embedding sdapi v1 train embedding post params
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
