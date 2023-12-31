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

// MatchImagesByTagsReq MatchImagesByTagsReq
//
// swagger:model MatchImagesByTagsReq
type MatchImagesByTagsReq struct {

	// And Tags
	// Required: true
	AndTags []int64 `json:"and_tags"`

	// Not Tags
	// Required: true
	NotTags []int64 `json:"not_tags"`

	// Or Tags
	// Required: true
	OrTags []int64 `json:"or_tags"`
}

// Validate validates this match images by tags req
func (m *MatchImagesByTagsReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAndTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchImagesByTagsReq) validateAndTags(formats strfmt.Registry) error {

	if err := validate.Required("and_tags", "body", m.AndTags); err != nil {
		return err
	}

	return nil
}

func (m *MatchImagesByTagsReq) validateNotTags(formats strfmt.Registry) error {

	if err := validate.Required("not_tags", "body", m.NotTags); err != nil {
		return err
	}

	return nil
}

func (m *MatchImagesByTagsReq) validateOrTags(formats strfmt.Registry) error {

	if err := validate.Required("or_tags", "body", m.OrTags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this match images by tags req based on context it is used
func (m *MatchImagesByTagsReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchImagesByTagsReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchImagesByTagsReq) UnmarshalBinary(b []byte) error {
	var res MatchImagesByTagsReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
