///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// GetFunctionOKCode is the HTTP code returned for type GetFunctionOK
const GetFunctionOKCode int = 200

/*GetFunctionOK Successful operation

swagger:response getFunctionOK
*/
type GetFunctionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Function `json:"body,omitempty"`
}

// NewGetFunctionOK creates GetFunctionOK with default headers values
func NewGetFunctionOK() *GetFunctionOK {
	return &GetFunctionOK{}
}

// WithPayload adds the payload to the get function o k response
func (o *GetFunctionOK) WithPayload(payload *models.Function) *GetFunctionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function o k response
func (o *GetFunctionOK) SetPayload(payload *models.Function) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFunctionBadRequestCode is the HTTP code returned for type GetFunctionBadRequest
const GetFunctionBadRequestCode int = 400

/*GetFunctionBadRequest Invalid Name supplied

swagger:response getFunctionBadRequest
*/
type GetFunctionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionBadRequest creates GetFunctionBadRequest with default headers values
func NewGetFunctionBadRequest() *GetFunctionBadRequest {
	return &GetFunctionBadRequest{}
}

// WithPayload adds the payload to the get function bad request response
func (o *GetFunctionBadRequest) WithPayload(payload *models.Error) *GetFunctionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function bad request response
func (o *GetFunctionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFunctionNotFoundCode is the HTTP code returned for type GetFunctionNotFound
const GetFunctionNotFoundCode int = 404

/*GetFunctionNotFound Function not found

swagger:response getFunctionNotFound
*/
type GetFunctionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionNotFound creates GetFunctionNotFound with default headers values
func NewGetFunctionNotFound() *GetFunctionNotFound {
	return &GetFunctionNotFound{}
}

// WithPayload adds the payload to the get function not found response
func (o *GetFunctionNotFound) WithPayload(payload *models.Error) *GetFunctionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function not found response
func (o *GetFunctionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFunctionInternalServerErrorCode is the HTTP code returned for type GetFunctionInternalServerError
const GetFunctionInternalServerErrorCode int = 500

/*GetFunctionInternalServerError Internal error

swagger:response getFunctionInternalServerError
*/
type GetFunctionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionInternalServerError creates GetFunctionInternalServerError with default headers values
func NewGetFunctionInternalServerError() *GetFunctionInternalServerError {
	return &GetFunctionInternalServerError{}
}

// WithPayload adds the payload to the get function internal server error response
func (o *GetFunctionInternalServerError) WithPayload(payload *models.Error) *GetFunctionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function internal server error response
func (o *GetFunctionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
