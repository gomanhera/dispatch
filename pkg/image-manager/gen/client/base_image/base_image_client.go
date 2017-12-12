///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new base image API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for base image API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddBaseImage adds a new base image
*/
func (a *Client) AddBaseImage(params *AddBaseImageParams, authInfo runtime.ClientAuthInfoWriter) (*AddBaseImageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBaseImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addBaseImage",
		Method:             "POST",
		PathPattern:        "/base",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddBaseImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddBaseImageCreated), nil

}

/*
DeleteBaseImageByName deletes a base image
*/
func (a *Client) DeleteBaseImageByName(params *DeleteBaseImageByNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBaseImageByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBaseImageByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBaseImageByName",
		Method:             "DELETE",
		PathPattern:        "/base/{baseImageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBaseImageByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteBaseImageByNameOK), nil

}

/*
GetBaseImageByName finds base image by name

Returns a single base image
*/
func (a *Client) GetBaseImageByName(params *GetBaseImageByNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetBaseImageByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBaseImageByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBaseImageByName",
		Method:             "GET",
		PathPattern:        "/base/{baseImageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBaseImageByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBaseImageByNameOK), nil

}

/*
GetBaseImages lists all existing base images
*/
func (a *Client) GetBaseImages(params *GetBaseImagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetBaseImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBaseImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBaseImages",
		Method:             "GET",
		PathPattern:        "/base",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBaseImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBaseImagesOK), nil

}

/*
UpdateBaseImageByName updates a base image
*/
func (a *Client) UpdateBaseImageByName(params *UpdateBaseImageByNameParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBaseImageByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBaseImageByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateBaseImageByName",
		Method:             "PUT",
		PathPattern:        "/base/{baseImageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBaseImageByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateBaseImageByNameOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
