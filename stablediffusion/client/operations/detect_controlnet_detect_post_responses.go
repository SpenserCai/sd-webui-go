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

// DetectControlnetDetectPostReader is a Reader for the DetectControlnetDetectPost structure.
type DetectControlnetDetectPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetectControlnetDetectPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetectControlnetDetectPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewDetectControlnetDetectPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDetectControlnetDetectPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetectControlnetDetectPostOK creates a DetectControlnetDetectPostOK with default headers values
func NewDetectControlnetDetectPostOK() *DetectControlnetDetectPostOK {
	return &DetectControlnetDetectPostOK{}
}

/*
DetectControlnetDetectPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type DetectControlnetDetectPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this detect controlnet detect post o k response has a 2xx status code
func (o *DetectControlnetDetectPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this detect controlnet detect post o k response has a 3xx status code
func (o *DetectControlnetDetectPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detect controlnet detect post o k response has a 4xx status code
func (o *DetectControlnetDetectPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this detect controlnet detect post o k response has a 5xx status code
func (o *DetectControlnetDetectPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this detect controlnet detect post o k response a status code equal to that given
func (o *DetectControlnetDetectPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the detect controlnet detect post o k response
func (o *DetectControlnetDetectPostOK) Code() int {
	return 200
}

func (o *DetectControlnetDetectPostOK) Error() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostOK  %+v", 200, o.Payload)
}

func (o *DetectControlnetDetectPostOK) String() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostOK  %+v", 200, o.Payload)
}

func (o *DetectControlnetDetectPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DetectControlnetDetectPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetectControlnetDetectPostUnprocessableEntity creates a DetectControlnetDetectPostUnprocessableEntity with default headers values
func NewDetectControlnetDetectPostUnprocessableEntity() *DetectControlnetDetectPostUnprocessableEntity {
	return &DetectControlnetDetectPostUnprocessableEntity{}
}

/*
DetectControlnetDetectPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type DetectControlnetDetectPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this detect controlnet detect post unprocessable entity response has a 2xx status code
func (o *DetectControlnetDetectPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detect controlnet detect post unprocessable entity response has a 3xx status code
func (o *DetectControlnetDetectPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detect controlnet detect post unprocessable entity response has a 4xx status code
func (o *DetectControlnetDetectPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this detect controlnet detect post unprocessable entity response has a 5xx status code
func (o *DetectControlnetDetectPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this detect controlnet detect post unprocessable entity response a status code equal to that given
func (o *DetectControlnetDetectPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the detect controlnet detect post unprocessable entity response
func (o *DetectControlnetDetectPostUnprocessableEntity) Code() int {
	return 422
}

func (o *DetectControlnetDetectPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DetectControlnetDetectPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DetectControlnetDetectPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *DetectControlnetDetectPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetectControlnetDetectPostInternalServerError creates a DetectControlnetDetectPostInternalServerError with default headers values
func NewDetectControlnetDetectPostInternalServerError() *DetectControlnetDetectPostInternalServerError {
	return &DetectControlnetDetectPostInternalServerError{}
}

/*
DetectControlnetDetectPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type DetectControlnetDetectPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this detect controlnet detect post internal server error response has a 2xx status code
func (o *DetectControlnetDetectPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detect controlnet detect post internal server error response has a 3xx status code
func (o *DetectControlnetDetectPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detect controlnet detect post internal server error response has a 4xx status code
func (o *DetectControlnetDetectPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this detect controlnet detect post internal server error response has a 5xx status code
func (o *DetectControlnetDetectPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this detect controlnet detect post internal server error response a status code equal to that given
func (o *DetectControlnetDetectPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the detect controlnet detect post internal server error response
func (o *DetectControlnetDetectPostInternalServerError) Code() int {
	return 500
}

func (o *DetectControlnetDetectPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostInternalServerError  %+v", 500, o.Payload)
}

func (o *DetectControlnetDetectPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /controlnet/detect][%d] detectControlnetDetectPostInternalServerError  %+v", 500, o.Payload)
}

func (o *DetectControlnetDetectPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *DetectControlnetDetectPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
