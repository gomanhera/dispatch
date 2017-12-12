///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIParams creates a new GetAPIParams object
// with the default values initialized.
func NewGetAPIParams() *GetAPIParams {
	var ()
	return &GetAPIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIParamsWithTimeout creates a new GetAPIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIParamsWithTimeout(timeout time.Duration) *GetAPIParams {
	var ()
	return &GetAPIParams{

		timeout: timeout,
	}
}

// NewGetAPIParamsWithContext creates a new GetAPIParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIParamsWithContext(ctx context.Context) *GetAPIParams {
	var ()
	return &GetAPIParams{

		Context: ctx,
	}
}

// NewGetAPIParamsWithHTTPClient creates a new GetAPIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIParamsWithHTTPClient(client *http.Client) *GetAPIParams {
	var ()
	return &GetAPIParams{
		HTTPClient: client,
	}
}

/*GetAPIParams contains all the parameters to send to the API endpoint
for the get API operation typically these are written to a http.Request
*/
type GetAPIParams struct {

	/*API
	  Name of API to work on

	*/
	API string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API params
func (o *GetAPIParams) WithTimeout(timeout time.Duration) *GetAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API params
func (o *GetAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API params
func (o *GetAPIParams) WithContext(ctx context.Context) *GetAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API params
func (o *GetAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API params
func (o *GetAPIParams) WithHTTPClient(client *http.Client) *GetAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API params
func (o *GetAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPI adds the api to the get API params
func (o *GetAPIParams) WithAPI(api string) *GetAPIParams {
	o.SetAPI(api)
	return o
}

// SetAPI adds the api to the get API params
func (o *GetAPIParams) SetAPI(api string) {
	o.API = api
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api
	if err := r.SetPathParam("api", o.API); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
