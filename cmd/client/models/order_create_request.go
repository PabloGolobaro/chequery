// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderCreateRequest order create request
//
// swagger:model OrderCreateRequest
type OrderCreateRequest struct {

	// order
	Order string `json:"order,omitempty"`
}

// Validate validates this order create request
func (m *OrderCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this order create request based on context it is used
func (m *OrderCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderCreateRequest) UnmarshalBinary(b []byte) error {
	var res OrderCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}