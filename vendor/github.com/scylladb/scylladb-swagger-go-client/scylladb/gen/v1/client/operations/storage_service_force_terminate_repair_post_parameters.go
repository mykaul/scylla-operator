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
	"github.com/go-openapi/strfmt"
)

// NewStorageServiceForceTerminateRepairPostParams creates a new StorageServiceForceTerminateRepairPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceForceTerminateRepairPostParams() *StorageServiceForceTerminateRepairPostParams {
	return &StorageServiceForceTerminateRepairPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceForceTerminateRepairPostParamsWithTimeout creates a new StorageServiceForceTerminateRepairPostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceForceTerminateRepairPostParamsWithTimeout(timeout time.Duration) *StorageServiceForceTerminateRepairPostParams {
	return &StorageServiceForceTerminateRepairPostParams{
		timeout: timeout,
	}
}

// NewStorageServiceForceTerminateRepairPostParamsWithContext creates a new StorageServiceForceTerminateRepairPostParams object
// with the ability to set a context for a request.
func NewStorageServiceForceTerminateRepairPostParamsWithContext(ctx context.Context) *StorageServiceForceTerminateRepairPostParams {
	return &StorageServiceForceTerminateRepairPostParams{
		Context: ctx,
	}
}

// NewStorageServiceForceTerminateRepairPostParamsWithHTTPClient creates a new StorageServiceForceTerminateRepairPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceForceTerminateRepairPostParamsWithHTTPClient(client *http.Client) *StorageServiceForceTerminateRepairPostParams {
	return &StorageServiceForceTerminateRepairPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceForceTerminateRepairPostParams contains all the parameters to send to the API endpoint

	for the storage service force terminate repair post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceForceTerminateRepairPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service force terminate repair post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceForceTerminateRepairPostParams) WithDefaults() *StorageServiceForceTerminateRepairPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service force terminate repair post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceForceTerminateRepairPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) WithTimeout(timeout time.Duration) *StorageServiceForceTerminateRepairPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) WithContext(ctx context.Context) *StorageServiceForceTerminateRepairPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) WithHTTPClient(client *http.Client) *StorageServiceForceTerminateRepairPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service force terminate repair post params
func (o *StorageServiceForceTerminateRepairPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceForceTerminateRepairPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}