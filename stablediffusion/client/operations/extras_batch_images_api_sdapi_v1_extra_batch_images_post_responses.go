// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

// ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostReader is a Reader for the ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPost structure.
type ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK creates a ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK with default headers values
func NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK() *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK {
	return &ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK{}
}

/*
ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK struct {
	Payload *models.ExtrasBatchImagesResponse
}

// IsSuccess returns true when this extras batch images Api sdapi v1 extra batch images post o k response has a 2xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras batch images Api sdapi v1 extra batch images post o k response has a 3xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras batch images Api sdapi v1 extra batch images post o k response has a 4xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras batch images Api sdapi v1 extra batch images post o k response has a 5xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras batch images Api sdapi v1 extra batch images post o k response a status code equal to that given
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras batch images Api sdapi v1 extra batch images post o k response
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) Code() int {
	return 200
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostOK  %+v", 200, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostOK  %+v", 200, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) GetPayload() *models.ExtrasBatchImagesResponse {
	return o.Payload
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtrasBatchImagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity creates a ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity with default headers values
func NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity() *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity {
	return &ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity{}
}

/*
ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this extras batch images Api sdapi v1 extra batch images post unprocessable entity response has a 2xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras batch images Api sdapi v1 extra batch images post unprocessable entity response has a 3xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras batch images Api sdapi v1 extra batch images post unprocessable entity response has a 4xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras batch images Api sdapi v1 extra batch images post unprocessable entity response has a 5xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this extras batch images Api sdapi v1 extra batch images post unprocessable entity response a status code equal to that given
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the extras batch images Api sdapi v1 extra batch images post unprocessable entity response
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) Code() int {
	return 422
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError creates a ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError with default headers values
func NewExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError() *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError {
	return &ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError{}
}

/*
ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this extras batch images Api sdapi v1 extra batch images post internal server error response has a 2xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras batch images Api sdapi v1 extra batch images post internal server error response has a 3xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras batch images Api sdapi v1 extra batch images post internal server error response has a 4xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras batch images Api sdapi v1 extra batch images post internal server error response has a 5xx status code
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this extras batch images Api sdapi v1 extra batch images post internal server error response a status code equal to that given
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the extras batch images Api sdapi v1 extra batch images post internal server error response
func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) Code() int {
	return 500
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/extra-batch-images][%d] extrasBatchImagesApiSdapiV1ExtraBatchImagesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *ExtrasBatchImagesAPISdapiV1ExtraBatchImagesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
