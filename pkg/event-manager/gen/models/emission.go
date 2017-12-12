///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Emission emission
// swagger:model Emission

type Emission struct {

	// emitted time
	// Read Only: true
	EmittedTime int64 `json:"emittedTime,omitempty"`

	// id
	// Read Only: true
	ID strfmt.UUID `json:"id,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// topic
	// Required: true
	// Pattern: ^[\w\d\-\.]+$
	Topic *string `json:"topic"`
}

/* polymorph Emission emittedTime false */

/* polymorph Emission id false */

/* polymorph Emission payload false */

/* polymorph Emission topic false */

// Validate validates this emission
func (m *Emission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Emission) validateTopic(formats strfmt.Registry) error {

	if err := validate.Required("topic", "body", m.Topic); err != nil {
		return err
	}

	if err := validate.Pattern("topic", "body", string(*m.Topic), `^[\w\d\-\.]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Emission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Emission) UnmarshalBinary(b []byte) error {
	var res Emission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
