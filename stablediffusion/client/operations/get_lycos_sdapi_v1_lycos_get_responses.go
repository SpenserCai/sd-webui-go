// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLycosSdapiV1LycosGetReader is a Reader for the GetLycosSdapiV1LycosGet structure.
type GetLycosSdapiV1LycosGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLycosSdapiV1LycosGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLycosSdapiV1LycosGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLycosSdapiV1LycosGetOK creates a GetLycosSdapiV1LycosGetOK with default headers values
func NewGetLycosSdapiV1LycosGetOK() *GetLycosSdapiV1LycosGetOK {
	return &GetLycosSdapiV1LycosGetOK{}
}

/*
GetLycosSdapiV1LycosGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetLycosSdapiV1LycosGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get lycos sdapi v1 lycos get o k response has a 2xx status code
func (o *GetLycosSdapiV1LycosGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get lycos sdapi v1 lycos get o k response has a 3xx status code
func (o *GetLycosSdapiV1LycosGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lycos sdapi v1 lycos get o k response has a 4xx status code
func (o *GetLycosSdapiV1LycosGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get lycos sdapi v1 lycos get o k response has a 5xx status code
func (o *GetLycosSdapiV1LycosGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get lycos sdapi v1 lycos get o k response a status code equal to that given
func (o *GetLycosSdapiV1LycosGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get lycos sdapi v1 lycos get o k response
func (o *GetLycosSdapiV1LycosGetOK) Code() int {
	return 200
}

func (o *GetLycosSdapiV1LycosGetOK) Error() string {
	return fmt.Sprintf("[GET /sdapi/v1/lycos][%d] getLycosSdapiV1LycosGetOK  %+v", 200, o.Payload)
}

func (o *GetLycosSdapiV1LycosGetOK) String() string {
	return fmt.Sprintf("[GET /sdapi/v1/lycos][%d] getLycosSdapiV1LycosGetOK  %+v", 200, o.Payload)
}

func (o *GetLycosSdapiV1LycosGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLycosSdapiV1LycosGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
