// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/run_model"
)

// RunServiceGetRunReader is a Reader for the RunServiceGetRun structure.
type RunServiceGetRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunServiceGetRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunServiceGetRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRunServiceGetRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunServiceGetRunOK creates a RunServiceGetRunOK with default headers values
func NewRunServiceGetRunOK() *RunServiceGetRunOK {
	return &RunServiceGetRunOK{}
}

/*RunServiceGetRunOK handles this case with default header values.

A successful response.
*/
type RunServiceGetRunOK struct {
	Payload *run_model.V2beta1Run
}

func (o *RunServiceGetRunOK) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/runs/{run_id}][%d] runServiceGetRunOK  %+v", 200, o.Payload)
}

func (o *RunServiceGetRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V2beta1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunServiceGetRunDefault creates a RunServiceGetRunDefault with default headers values
func NewRunServiceGetRunDefault(code int) *RunServiceGetRunDefault {
	return &RunServiceGetRunDefault{
		_statusCode: code,
	}
}

/*RunServiceGetRunDefault handles this case with default header values.

An unexpected error response.
*/
type RunServiceGetRunDefault struct {
	_statusCode int

	Payload *run_model.RuntimeError
}

// Code gets the status code for the run service get run default response
func (o *RunServiceGetRunDefault) Code() int {
	return o._statusCode
}

func (o *RunServiceGetRunDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/runs/{run_id}][%d] RunService_GetRun default  %+v", o._statusCode, o.Payload)
}

func (o *RunServiceGetRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
