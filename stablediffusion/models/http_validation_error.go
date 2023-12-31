// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HTTPValidationError HTTPValidationError
//
// swagger:model HTTPValidationError
type HTTPValidationError struct {

	// Detail
	Detail []*ValidationError `json:"detail"`
}

// Validate validates this HTTP validation error
func (m *HTTPValidationError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPValidationError) validateDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.Detail) { // not required
		return nil
	}

	for i := 0; i < len(m.Detail); i++ {
		if swag.IsZero(m.Detail[i]) { // not required
			continue
		}

		if m.Detail[i] != nil {
			if err := m.Detail[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detail" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detail" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this HTTP validation error based on the context it is used
func (m *HTTPValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPValidationError) contextValidateDetail(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Detail); i++ {

		if m.Detail[i] != nil {
			if err := m.Detail[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("detail" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("detail" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPValidationError) UnmarshalBinary(b []byte) error {
	var res HTTPValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
