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

// BodyAPIControlnetSegSamControlnetSegPost Body_api_controlnet_seg_sam_controlnet_seg_post
//
// swagger:model Body_api_controlnet_seg_sam_controlnet_seg_post
type BodyAPIControlnetSegSamControlnetSegPost struct {

	// autosam conf
	// Required: true
	AutosamConf *AutoSAMConfig `json:"autosam_conf"`

	// payload
	// Required: true
	Payload *ControlNetSegRequest `json:"payload"`
}

// Validate validates this body api controlnet seg sam controlnet seg post
func (m *BodyAPIControlnetSegSamControlnetSegPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutosamConf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BodyAPIControlnetSegSamControlnetSegPost) validateAutosamConf(formats strfmt.Registry) error {

	if err := validate.Required("autosam_conf", "body", m.AutosamConf); err != nil {
		return err
	}

	if m.AutosamConf != nil {
		if err := m.AutosamConf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosam_conf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autosam_conf")
			}
			return err
		}
	}

	return nil
}

func (m *BodyAPIControlnetSegSamControlnetSegPost) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this body api controlnet seg sam controlnet seg post based on the context it is used
func (m *BodyAPIControlnetSegSamControlnetSegPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutosamConf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BodyAPIControlnetSegSamControlnetSegPost) contextValidateAutosamConf(ctx context.Context, formats strfmt.Registry) error {

	if m.AutosamConf != nil {
		if err := m.AutosamConf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autosam_conf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autosam_conf")
			}
			return err
		}
	}

	return nil
}

func (m *BodyAPIControlnetSegSamControlnetSegPost) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.Payload != nil {
		if err := m.Payload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BodyAPIControlnetSegSamControlnetSegPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BodyAPIControlnetSegSamControlnetSegPost) UnmarshalBinary(b []byte) error {
	var res BodyAPIControlnetSegSamControlnetSegPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
