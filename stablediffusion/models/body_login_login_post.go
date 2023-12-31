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

// BodyLoginLoginPost Body_login_login_post
//
// swagger:model Body_login_login_post
type BodyLoginLoginPost struct {

	// Client Id
	ClientID string `json:"client_id,omitempty"`

	// Client Secret
	ClientSecret string `json:"client_secret,omitempty"`

	// Grant Type
	// Pattern: password
	GrantType string `json:"grant_type,omitempty"`

	// Password
	// Required: true
	Password *string `json:"password"`

	// Scope
	Scope string `json:"scope,omitempty"`

	// Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this body login login post
func (m *BodyLoginLoginPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BodyLoginLoginPost) validateGrantType(formats strfmt.Registry) error {
	if swag.IsZero(m.GrantType) { // not required
		return nil
	}

	if err := validate.Pattern("grant_type", "body", m.GrantType, `password`); err != nil {
		return err
	}

	return nil
}

func (m *BodyLoginLoginPost) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *BodyLoginLoginPost) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this body login login post based on context it is used
func (m *BodyLoginLoginPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BodyLoginLoginPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BodyLoginLoginPost) UnmarshalBinary(b []byte) error {
	var res BodyLoginLoginPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
