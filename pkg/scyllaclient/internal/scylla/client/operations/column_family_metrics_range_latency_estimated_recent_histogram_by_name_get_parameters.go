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

// NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams creates a new ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams() *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics range latency estimated recent histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics range latency estimated recent histogram by name get params
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsRangeLatencyEstimatedRecentHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
