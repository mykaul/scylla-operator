// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceJoinRingGetParams creates a new StorageServiceJoinRingGetParams object
// with the default values initialized.
func NewStorageServiceJoinRingGetParams() *StorageServiceJoinRingGetParams {

	return &StorageServiceJoinRingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceJoinRingGetParamsWithTimeout creates a new StorageServiceJoinRingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceJoinRingGetParamsWithTimeout(timeout time.Duration) *StorageServiceJoinRingGetParams {

	return &StorageServiceJoinRingGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceJoinRingGetParamsWithContext creates a new StorageServiceJoinRingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceJoinRingGetParamsWithContext(ctx context.Context) *StorageServiceJoinRingGetParams {

	return &StorageServiceJoinRingGetParams{

		Context: ctx,
	}
}

// NewStorageServiceJoinRingGetParamsWithHTTPClient creates a new StorageServiceJoinRingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceJoinRingGetParamsWithHTTPClient(client *http.Client) *StorageServiceJoinRingGetParams {

	return &StorageServiceJoinRingGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceJoinRingGetParams contains all the parameters to send to the API endpoint
for the storage service join ring get operation typically these are written to a http.Request
*/
type StorageServiceJoinRingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) WithTimeout(timeout time.Duration) *StorageServiceJoinRingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) WithContext(ctx context.Context) *StorageServiceJoinRingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) WithHTTPClient(client *http.Client) *StorageServiceJoinRingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service join ring get params
func (o *StorageServiceJoinRingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceJoinRingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
