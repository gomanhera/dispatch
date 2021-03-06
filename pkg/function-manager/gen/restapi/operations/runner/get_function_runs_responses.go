///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// GetFunctionRunsOKCode is the HTTP code returned for type GetFunctionRunsOK
const GetFunctionRunsOKCode int = 200

/*GetFunctionRunsOK List of function runs

swagger:response getFunctionRunsOK
*/
type GetFunctionRunsOK struct {

	/*
	  In: Body
	*/
	Payload models.GetFunctionRunsOKBody `json:"body,omitempty"`
}

// NewGetFunctionRunsOK creates GetFunctionRunsOK with default headers values
func NewGetFunctionRunsOK() *GetFunctionRunsOK {
	return &GetFunctionRunsOK{}
}

// WithPayload adds the payload to the get function runs o k response
func (o *GetFunctionRunsOK) WithPayload(payload models.GetFunctionRunsOKBody) *GetFunctionRunsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function runs o k response
func (o *GetFunctionRunsOK) SetPayload(payload models.GetFunctionRunsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionRunsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetFunctionRunsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetFunctionRunsBadRequestCode is the HTTP code returned for type GetFunctionRunsBadRequest
const GetFunctionRunsBadRequestCode int = 400

/*GetFunctionRunsBadRequest Bad Request

swagger:response getFunctionRunsBadRequest
*/
type GetFunctionRunsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionRunsBadRequest creates GetFunctionRunsBadRequest with default headers values
func NewGetFunctionRunsBadRequest() *GetFunctionRunsBadRequest {
	return &GetFunctionRunsBadRequest{}
}

// WithPayload adds the payload to the get function runs bad request response
func (o *GetFunctionRunsBadRequest) WithPayload(payload *models.Error) *GetFunctionRunsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function runs bad request response
func (o *GetFunctionRunsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionRunsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFunctionRunsNotFoundCode is the HTTP code returned for type GetFunctionRunsNotFound
const GetFunctionRunsNotFoundCode int = 404

/*GetFunctionRunsNotFound Function not found

swagger:response getFunctionRunsNotFound
*/
type GetFunctionRunsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionRunsNotFound creates GetFunctionRunsNotFound with default headers values
func NewGetFunctionRunsNotFound() *GetFunctionRunsNotFound {
	return &GetFunctionRunsNotFound{}
}

// WithPayload adds the payload to the get function runs not found response
func (o *GetFunctionRunsNotFound) WithPayload(payload *models.Error) *GetFunctionRunsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function runs not found response
func (o *GetFunctionRunsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionRunsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFunctionRunsInternalServerErrorCode is the HTTP code returned for type GetFunctionRunsInternalServerError
const GetFunctionRunsInternalServerErrorCode int = 500

/*GetFunctionRunsInternalServerError Internal error

swagger:response getFunctionRunsInternalServerError
*/
type GetFunctionRunsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFunctionRunsInternalServerError creates GetFunctionRunsInternalServerError with default headers values
func NewGetFunctionRunsInternalServerError() *GetFunctionRunsInternalServerError {
	return &GetFunctionRunsInternalServerError{}
}

// WithPayload adds the payload to the get function runs internal server error response
func (o *GetFunctionRunsInternalServerError) WithPayload(payload *models.Error) *GetFunctionRunsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get function runs internal server error response
func (o *GetFunctionRunsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFunctionRunsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
