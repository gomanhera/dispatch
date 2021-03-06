///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/image-manager/gen/models"
)

// GetBaseImagesOKCode is the HTTP code returned for type GetBaseImagesOK
const GetBaseImagesOKCode int = 200

/*GetBaseImagesOK successful operation

swagger:response getBaseImagesOK
*/
type GetBaseImagesOK struct {

	/*
	  In: Body
	*/
	Payload models.GetBaseImagesOKBody `json:"body,omitempty"`
}

// NewGetBaseImagesOK creates GetBaseImagesOK with default headers values
func NewGetBaseImagesOK() *GetBaseImagesOK {
	return &GetBaseImagesOK{}
}

// WithPayload adds the payload to the get base images o k response
func (o *GetBaseImagesOK) WithPayload(payload models.GetBaseImagesOKBody) *GetBaseImagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base images o k response
func (o *GetBaseImagesOK) SetPayload(payload models.GetBaseImagesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetBaseImagesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetBaseImagesDefault Generic error response

swagger:response getBaseImagesDefault
*/
type GetBaseImagesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBaseImagesDefault creates GetBaseImagesDefault with default headers values
func NewGetBaseImagesDefault(code int) *GetBaseImagesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBaseImagesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get base images default response
func (o *GetBaseImagesDefault) WithStatusCode(code int) *GetBaseImagesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get base images default response
func (o *GetBaseImagesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get base images default response
func (o *GetBaseImagesDefault) WithPayload(payload *models.Error) *GetBaseImagesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get base images default response
func (o *GetBaseImagesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBaseImagesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
