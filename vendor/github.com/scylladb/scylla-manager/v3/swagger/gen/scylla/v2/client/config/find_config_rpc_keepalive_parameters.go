// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindConfigRPCKeepaliveParams creates a new FindConfigRPCKeepaliveParams object
// with the default values initialized.
func NewFindConfigRPCKeepaliveParams() *FindConfigRPCKeepaliveParams {

	return &FindConfigRPCKeepaliveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRPCKeepaliveParamsWithTimeout creates a new FindConfigRPCKeepaliveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRPCKeepaliveParamsWithTimeout(timeout time.Duration) *FindConfigRPCKeepaliveParams {

	return &FindConfigRPCKeepaliveParams{

		timeout: timeout,
	}
}

// NewFindConfigRPCKeepaliveParamsWithContext creates a new FindConfigRPCKeepaliveParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRPCKeepaliveParamsWithContext(ctx context.Context) *FindConfigRPCKeepaliveParams {

	return &FindConfigRPCKeepaliveParams{

		Context: ctx,
	}
}

// NewFindConfigRPCKeepaliveParamsWithHTTPClient creates a new FindConfigRPCKeepaliveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRPCKeepaliveParamsWithHTTPClient(client *http.Client) *FindConfigRPCKeepaliveParams {

	return &FindConfigRPCKeepaliveParams{
		HTTPClient: client,
	}
}

/*
FindConfigRPCKeepaliveParams contains all the parameters to send to the API endpoint
for the find config rpc keepalive operation typically these are written to a http.Request
*/
type FindConfigRPCKeepaliveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) WithTimeout(timeout time.Duration) *FindConfigRPCKeepaliveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) WithContext(ctx context.Context) *FindConfigRPCKeepaliveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) WithHTTPClient(client *http.Client) *FindConfigRPCKeepaliveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config rpc keepalive params
func (o *FindConfigRPCKeepaliveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRPCKeepaliveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}