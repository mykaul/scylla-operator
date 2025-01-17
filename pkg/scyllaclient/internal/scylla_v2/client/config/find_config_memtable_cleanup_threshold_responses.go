// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigMemtableCleanupThresholdReader is a Reader for the FindConfigMemtableCleanupThreshold structure.
type FindConfigMemtableCleanupThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableCleanupThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableCleanupThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableCleanupThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableCleanupThresholdOK creates a FindConfigMemtableCleanupThresholdOK with default headers values
func NewFindConfigMemtableCleanupThresholdOK() *FindConfigMemtableCleanupThresholdOK {
	return &FindConfigMemtableCleanupThresholdOK{}
}

/*
FindConfigMemtableCleanupThresholdOK handles this case with default header values.

Config value
*/
type FindConfigMemtableCleanupThresholdOK struct {
	Payload float64
}

func (o *FindConfigMemtableCleanupThresholdOK) GetPayload() float64 {
	return o.Payload
}

func (o *FindConfigMemtableCleanupThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableCleanupThresholdDefault creates a FindConfigMemtableCleanupThresholdDefault with default headers values
func NewFindConfigMemtableCleanupThresholdDefault(code int) *FindConfigMemtableCleanupThresholdDefault {
	return &FindConfigMemtableCleanupThresholdDefault{
		_statusCode: code,
	}
}

/*
FindConfigMemtableCleanupThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableCleanupThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable cleanup threshold default response
func (o *FindConfigMemtableCleanupThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableCleanupThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableCleanupThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemtableCleanupThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
