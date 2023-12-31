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

// ModulesAPIModelsProgressResponse ProgressResponse
//
// swagger:model modules__api__models__ProgressResponse
type ModulesAPIModelsProgressResponse struct {

	// Current image
	//
	// The current image in base64 format. opts.show_progress_every_n_steps is required for this to work.
	CurrentImage string `json:"current_image,omitempty"`

	// ETA in secs
	// Required: true
	EtaRelative *float64 `json:"eta_relative"`

	// Progress
	//
	// The progress with a range of 0 to 1
	// Required: true
	Progress *float64 `json:"progress"`

	// State
	//
	// The current state snapshot
	// Required: true
	State interface{} `json:"state"`

	// Info text
	//
	// Info text used by WebUI.
	Textinfo string `json:"textinfo,omitempty"`
}

// Validate validates this modules api models progress response
func (m *ModulesAPIModelsProgressResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEtaRelative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModulesAPIModelsProgressResponse) validateEtaRelative(formats strfmt.Registry) error {

	if err := validate.Required("eta_relative", "body", m.EtaRelative); err != nil {
		return err
	}

	return nil
}

func (m *ModulesAPIModelsProgressResponse) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *ModulesAPIModelsProgressResponse) validateState(formats strfmt.Registry) error {

	if m.State == nil {
		return errors.Required("state", "body", nil)
	}

	return nil
}

// ContextValidate validates this modules api models progress response based on context it is used
func (m *ModulesAPIModelsProgressResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModulesAPIModelsProgressResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModulesAPIModelsProgressResponse) UnmarshalBinary(b []byte) error {
	var res ModulesAPIModelsProgressResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
