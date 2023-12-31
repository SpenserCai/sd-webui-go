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

// DINOPredictRequest DINOPredictRequest
//
// swagger:model DINOPredictRequest
type DINOPredictRequest struct {

	// Box Threshold
	BoxThreshold *float64 `json:"box_threshold,omitempty"`

	// Dino Model Name
	DinoModelName *string `json:"dino_model_name,omitempty"`

	// Input Image
	// Required: true
	InputImage *string `json:"input_image"`

	// Text Prompt
	// Required: true
	TextPrompt *string `json:"text_prompt"`
}

// Validate validates this d i n o predict request
func (m *DINOPredictRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextPrompt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DINOPredictRequest) validateInputImage(formats strfmt.Registry) error {

	if err := validate.Required("input_image", "body", m.InputImage); err != nil {
		return err
	}

	return nil
}

func (m *DINOPredictRequest) validateTextPrompt(formats strfmt.Registry) error {

	if err := validate.Required("text_prompt", "body", m.TextPrompt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this d i n o predict request based on context it is used
func (m *DINOPredictRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DINOPredictRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DINOPredictRequest) UnmarshalBinary(b []byte) error {
	var res DINOPredictRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
