///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteAPIHandlerFunc turns a function with the right signature into a delete API handler
type DeleteAPIHandlerFunc func(DeleteAPIParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAPIHandlerFunc) Handle(params DeleteAPIParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAPIHandler interface for that can handle valid delete API params
type DeleteAPIHandler interface {
	Handle(DeleteAPIParams, interface{}) middleware.Responder
}

// NewDeleteAPI creates a new http.Handler for the delete API operation
func NewDeleteAPI(ctx *middleware.Context, handler DeleteAPIHandler) *DeleteAPI {
	return &DeleteAPI{Context: ctx, Handler: handler}
}

/*DeleteAPI swagger:route DELETE /{api} endpoint deleteApi

Deletes an API

*/
type DeleteAPI struct {
	Context *middleware.Context
	Handler DeleteAPIHandler
}

func (o *DeleteAPI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAPIParams()

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
