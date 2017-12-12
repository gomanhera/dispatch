///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// RunFunctionReader is a Reader for the RunFunction structure.
type RunFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewRunFunctionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRunFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRunFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewRunFunctionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRunFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewRunFunctionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunFunctionOK creates a RunFunctionOK with default headers values
func NewRunFunctionOK() *RunFunctionOK {
	return &RunFunctionOK{}
}

/*RunFunctionOK handles this case with default header values.

Successful execution (blocking call)
*/
type RunFunctionOK struct {
	Payload *models.Run
}

func (o *RunFunctionOK) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionOK  %+v", 200, o.Payload)
}

func (o *RunFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionAccepted creates a RunFunctionAccepted with default headers values
func NewRunFunctionAccepted() *RunFunctionAccepted {
	return &RunFunctionAccepted{}
}

/*RunFunctionAccepted handles this case with default header values.

Execution started (non-blocking call)
*/
type RunFunctionAccepted struct {
	Payload *models.Run
}

func (o *RunFunctionAccepted) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionAccepted  %+v", 202, o.Payload)
}

func (o *RunFunctionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionBadRequest creates a RunFunctionBadRequest with default headers values
func NewRunFunctionBadRequest() *RunFunctionBadRequest {
	return &RunFunctionBadRequest{}
}

/*RunFunctionBadRequest handles this case with default header values.

User error
*/
type RunFunctionBadRequest struct {
	Payload *models.Error
}

func (o *RunFunctionBadRequest) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *RunFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionNotFound creates a RunFunctionNotFound with default headers values
func NewRunFunctionNotFound() *RunFunctionNotFound {
	return &RunFunctionNotFound{}
}

/*RunFunctionNotFound handles this case with default header values.

Function not found
*/
type RunFunctionNotFound struct {
	Payload *models.Error
}

func (o *RunFunctionNotFound) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionNotFound  %+v", 404, o.Payload)
}

func (o *RunFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionUnprocessableEntity creates a RunFunctionUnprocessableEntity with default headers values
func NewRunFunctionUnprocessableEntity() *RunFunctionUnprocessableEntity {
	return &RunFunctionUnprocessableEntity{}
}

/*RunFunctionUnprocessableEntity handles this case with default header values.

Input object validation failed
*/
type RunFunctionUnprocessableEntity struct {
	Payload *models.Error
}

func (o *RunFunctionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RunFunctionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionInternalServerError creates a RunFunctionInternalServerError with default headers values
func NewRunFunctionInternalServerError() *RunFunctionInternalServerError {
	return &RunFunctionInternalServerError{}
}

/*RunFunctionInternalServerError handles this case with default header values.

Internal error
*/
type RunFunctionInternalServerError struct {
	Payload *models.Error
}

func (o *RunFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *RunFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunFunctionBadGateway creates a RunFunctionBadGateway with default headers values
func NewRunFunctionBadGateway() *RunFunctionBadGateway {
	return &RunFunctionBadGateway{}
}

/*RunFunctionBadGateway handles this case with default header values.

Function error occurred (blocking call)
*/
type RunFunctionBadGateway struct {
	Payload *models.Error
}

func (o *RunFunctionBadGateway) Error() string {
	return fmt.Sprintf("[POST /{functionName}/runs][%d] runFunctionBadGateway  %+v", 502, o.Payload)
}

func (o *RunFunctionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
