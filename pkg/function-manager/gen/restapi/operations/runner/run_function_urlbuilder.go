///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RunFunctionURL generates an URL for the run function operation
type RunFunctionURL struct {
	FunctionName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RunFunctionURL) WithBasePath(bp string) *RunFunctionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RunFunctionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RunFunctionURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/{functionName}/runs"

	functionName := o.FunctionName
	if functionName != "" {
		_path = strings.Replace(_path, "{functionName}", functionName, -1)
	} else {
		return nil, errors.New("FunctionName is required on RunFunctionURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1/function"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RunFunctionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RunFunctionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RunFunctionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RunFunctionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RunFunctionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *RunFunctionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
