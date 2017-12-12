///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewGetImageByNameParams creates a new GetImageByNameParams object
// with the default values initialized.
func NewGetImageByNameParams() *GetImageByNameParams {
	var ()
	return &GetImageByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageByNameParamsWithTimeout creates a new GetImageByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageByNameParamsWithTimeout(timeout time.Duration) *GetImageByNameParams {
	var ()
	return &GetImageByNameParams{

		timeout: timeout,
	}
}

// NewGetImageByNameParamsWithContext creates a new GetImageByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageByNameParamsWithContext(ctx context.Context) *GetImageByNameParams {
	var ()
	return &GetImageByNameParams{

		Context: ctx,
	}
}

// NewGetImageByNameParamsWithHTTPClient creates a new GetImageByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageByNameParamsWithHTTPClient(client *http.Client) *GetImageByNameParams {
	var ()
	return &GetImageByNameParams{
		HTTPClient: client,
	}
}

/*GetImageByNameParams contains all the parameters to send to the API endpoint
for the get image by name operation typically these are written to a http.Request
*/
type GetImageByNameParams struct {

	/*ImageName
	  Name of image to return

	*/
	ImageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image by name params
func (o *GetImageByNameParams) WithTimeout(timeout time.Duration) *GetImageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image by name params
func (o *GetImageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image by name params
func (o *GetImageByNameParams) WithContext(ctx context.Context) *GetImageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image by name params
func (o *GetImageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image by name params
func (o *GetImageByNameParams) WithHTTPClient(client *http.Client) *GetImageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image by name params
func (o *GetImageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageName adds the imageName to the get image by name params
func (o *GetImageByNameParams) WithImageName(imageName string) *GetImageByNameParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the get image by name params
func (o *GetImageByNameParams) SetImageName(imageName string) {
	o.ImageName = imageName
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageName
	if err := r.SetPathParam("imageName", o.ImageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
