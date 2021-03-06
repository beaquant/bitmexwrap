// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserCommission user commission
// swagger:model UserCommission
type UserCommission struct {

	// maker fee
	MakerFee float64 `json:"makerFee,omitempty"`

	// max fee
	MaxFee float64 `json:"maxFee,omitempty"`

	// settlement fee
	SettlementFee float64 `json:"settlementFee,omitempty"`

	// taker fee
	TakerFee float64 `json:"takerFee,omitempty"`
}

// Validate validates this user commission
func (m *UserCommission) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserCommission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCommission) UnmarshalBinary(b []byte) error {
	var res UserCommission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
