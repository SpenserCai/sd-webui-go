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

// ThemeCSSThemeCSSGetReader is a Reader for the ThemeCSSThemeCSSGet structure.
type ThemeCSSThemeCSSGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThemeCSSThemeCSSGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThemeCSSThemeCSSGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewThemeCSSThemeCSSGetOK creates a ThemeCSSThemeCSSGetOK with default headers values
func NewThemeCSSThemeCSSGetOK() *ThemeCSSThemeCSSGetOK {
	return &ThemeCSSThemeCSSGetOK{}
}

/*
ThemeCSSThemeCSSGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type ThemeCSSThemeCSSGetOK struct {
	Payload string
}

// IsSuccess returns true when this theme Css theme Css get o k response has a 2xx status code
func (o *ThemeCSSThemeCSSGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this theme Css theme Css get o k response has a 3xx status code
func (o *ThemeCSSThemeCSSGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this theme Css theme Css get o k response has a 4xx status code
func (o *ThemeCSSThemeCSSGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this theme Css theme Css get o k response has a 5xx status code
func (o *ThemeCSSThemeCSSGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this theme Css theme Css get o k response a status code equal to that given
func (o *ThemeCSSThemeCSSGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the theme Css theme Css get o k response
func (o *ThemeCSSThemeCSSGetOK) Code() int {
	return 200
}

func (o *ThemeCSSThemeCSSGetOK) Error() string {
	return fmt.Sprintf("[GET /theme.css][%d] themeCssThemeCssGetOK  %+v", 200, o.Payload)
}

func (o *ThemeCSSThemeCSSGetOK) String() string {
	return fmt.Sprintf("[GET /theme.css][%d] themeCssThemeCssGetOK  %+v", 200, o.Payload)
}

func (o *ThemeCSSThemeCSSGetOK) GetPayload() string {
	return o.Payload
}

func (o *ThemeCSSThemeCSSGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
