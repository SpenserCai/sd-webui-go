// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExtrasBatchImagesResponse ExtrasBatchImagesResponse
//
// swagger:model ExtrasBatchImagesResponse
type ExtrasBatchImagesResponse struct {

	// HTML info
	//
	// A series of HTML tags containing the process info.
	// Required: true
	HTMLInfo *string `json:"html_info"`

	// Images
	//
	// The generated images in base64 format.
	// Required: true
	Images []string `json:"images"`
}

// Validate validates this extras batch images response
func (m *ExtrasBatchImagesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTMLInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtrasBatchImagesResponse) validateHTMLInfo(formats strfmt.Registry) error {

	if err := validate.Required("html_info", "body", m.HTMLInfo); err != nil {
		return err
	}

	return nil
}

func (m *ExtrasBatchImagesResponse) validateImages(formats strfmt.Registry) error {

	if err := validate.Required("images", "body", m.Images); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this extras batch images response based on context it is used
func (m *ExtrasBatchImagesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtrasBatchImagesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtrasBatchImagesResponse) UnmarshalBinary(b []byte) error {
	var res ExtrasBatchImagesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
