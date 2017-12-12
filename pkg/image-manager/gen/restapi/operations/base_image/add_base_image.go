///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddBaseImageHandlerFunc turns a function with the right signature into a add base image handler
type AddBaseImageHandlerFunc func(AddBaseImageParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddBaseImageHandlerFunc) Handle(params AddBaseImageParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddBaseImageHandler interface for that can handle valid add base image params
type AddBaseImageHandler interface {
	Handle(AddBaseImageParams, interface{}) middleware.Responder
}

// NewAddBaseImage creates a new http.Handler for the add base image operation
func NewAddBaseImage(ctx *middleware.Context, handler AddBaseImageHandler) *AddBaseImage {
	return &AddBaseImage{Context: ctx, Handler: handler}
}

/*AddBaseImage swagger:route POST /base baseImage addBaseImage

Add a new base image

*/
type AddBaseImage struct {
	Context *middleware.Context
	Handler AddBaseImageHandler
}

func (o *AddBaseImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddBaseImageParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
