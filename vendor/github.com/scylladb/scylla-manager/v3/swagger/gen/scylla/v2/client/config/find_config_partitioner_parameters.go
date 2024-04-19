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

// NewFindConfigPartitionerParams creates a new FindConfigPartitionerParams object
// with the default values initialized.
func NewFindConfigPartitionerParams() *FindConfigPartitionerParams {

	return &FindConfigPartitionerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigPartitionerParamsWithTimeout creates a new FindConfigPartitionerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigPartitionerParamsWithTimeout(timeout time.Duration) *FindConfigPartitionerParams {

	return &FindConfigPartitionerParams{

		timeout: timeout,
	}
}

// NewFindConfigPartitionerParamsWithContext creates a new FindConfigPartitionerParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigPartitionerParamsWithContext(ctx context.Context) *FindConfigPartitionerParams {

	return &FindConfigPartitionerParams{

		Context: ctx,
	}
}

// NewFindConfigPartitionerParamsWithHTTPClient creates a new FindConfigPartitionerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigPartitionerParamsWithHTTPClient(client *http.Client) *FindConfigPartitionerParams {

	return &FindConfigPartitionerParams{
		HTTPClient: client,
	}
}

/*
FindConfigPartitionerParams contains all the parameters to send to the API endpoint
for the find config partitioner operation typically these are written to a http.Request
*/
type FindConfigPartitionerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config partitioner params
func (o *FindConfigPartitionerParams) WithTimeout(timeout time.Duration) *FindConfigPartitionerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config partitioner params
func (o *FindConfigPartitionerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config partitioner params
func (o *FindConfigPartitionerParams) WithContext(ctx context.Context) *FindConfigPartitionerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config partitioner params
func (o *FindConfigPartitionerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config partitioner params
func (o *FindConfigPartitionerParams) WithHTTPClient(client *http.Client) *FindConfigPartitionerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config partitioner params
func (o *FindConfigPartitionerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigPartitionerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}