// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Funding Swap Funding History
// swagger:model Funding
type Funding struct {

	// funding interval
	// Format: date-time
	FundingInterval strfmt.DateTime `json:"fundingInterval,omitempty"`

	// funding rate
	FundingRate float64 `json:"fundingRate,omitempty"`

	// funding rate daily
	FundingRateDaily float64 `json:"fundingRateDaily,omitempty"`

	// symbol
	// Required: true
	Symbol *string `json:"symbol"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this funding
func (m *Funding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFundingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Funding) validateFundingInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.FundingInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("fundingInterval", "body", "date-time", m.FundingInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Funding) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	return nil
}

func (m *Funding) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Funding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Funding) UnmarshalBinary(b []byte) error {
	var res Funding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
