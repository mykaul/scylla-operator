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

// NewFindConfigShutdownAnnounceInMsParams creates a new FindConfigShutdownAnnounceInMsParams object
// with the default values initialized.
func NewFindConfigShutdownAnnounceInMsParams() *FindConfigShutdownAnnounceInMsParams {

	return &FindConfigShutdownAnnounceInMsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigShutdownAnnounceInMsParamsWithTimeout creates a new FindConfigShutdownAnnounceInMsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigShutdownAnnounceInMsParamsWithTimeout(timeout time.Duration) *FindConfigShutdownAnnounceInMsParams {

	return &FindConfigShutdownAnnounceInMsParams{

		timeout: timeout,
	}
}

// NewFindConfigShutdownAnnounceInMsParamsWithContext creates a new FindConfigShutdownAnnounceInMsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigShutdownAnnounceInMsParamsWithContext(ctx context.Context) *FindConfigShutdownAnnounceInMsParams {

	return &FindConfigShutdownAnnounceInMsParams{

		Context: ctx,
	}
}

// NewFindConfigShutdownAnnounceInMsParamsWithHTTPClient creates a new FindConfigShutdownAnnounceInMsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigShutdownAnnounceInMsParamsWithHTTPClient(client *http.Client) *FindConfigShutdownAnnounceInMsParams {

	return &FindConfigShutdownAnnounceInMsParams{
		HTTPClient: client,
	}
}

/*
FindConfigShutdownAnnounceInMsParams contains all the parameters to send to the API endpoint
for the find config shutdown announce in ms operation typically these are written to a http.Request
*/
type FindConfigShutdownAnnounceInMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) WithTimeout(timeout time.Duration) *FindConfigShutdownAnnounceInMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) WithContext(ctx context.Context) *FindConfigShutdownAnnounceInMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) WithHTTPClient(client *http.Client) *FindConfigShutdownAnnounceInMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config shutdown announce in ms params
func (o *FindConfigShutdownAnnounceInMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigShutdownAnnounceInMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}