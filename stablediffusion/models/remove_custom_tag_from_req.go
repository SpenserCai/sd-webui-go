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

// RemoveCustomTagFromReq RemoveCustomTagFromReq
//
// swagger:model RemoveCustomTagFromReq
type RemoveCustomTagFromReq struct {

	// Img Id
	// Required: true
	ImgID *int64 `json:"img_id"`

	// Tag Id
	// Required: true
	TagID *string `json:"tag_id"`
}

// Validate validates this remove custom tag from req
func (m *RemoveCustomTagFromReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveCustomTagFromReq) validateImgID(formats strfmt.Registry) error {

	if err := validate.Required("img_id", "body", m.ImgID); err != nil {
		return err
	}

	return nil
}

func (m *RemoveCustomTagFromReq) validateTagID(formats strfmt.Registry) error {

	if err := validate.Required("tag_id", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this remove custom tag from req based on context it is used
func (m *RemoveCustomTagFromReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveCustomTagFromReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveCustomTagFromReq) UnmarshalBinary(b []byte) error {
	var res RemoveCustomTagFromReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
