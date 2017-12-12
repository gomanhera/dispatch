///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewLogoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/*LogoutOK handles this case with default header values.

successfully logged out
*/
type LogoutOK struct {
	SetCookie string

	Payload *models.Message
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[GET /v1/iam/logout][%d] logoutOK  %+v", 200, o.Payload)
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Set-Cookie
	o.SetCookie = response.GetHeader("Set-Cookie")

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutDefault creates a LogoutDefault with default headers values
func NewLogoutDefault(code int) *LogoutDefault {
	return &LogoutDefault{
		_statusCode: code,
	}
}

/*LogoutDefault handles this case with default header values.

error
*/
type LogoutDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the logout default response
func (o *LogoutDefault) Code() int {
	return o._statusCode
}

func (o *LogoutDefault) Error() string {
	return fmt.Sprintf("[GET /v1/iam/logout][%d] logout default  %+v", o._statusCode, o.Payload)
}

func (o *LogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
