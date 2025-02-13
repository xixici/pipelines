// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	experiment_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/experiment_model"
)

// ExperimentServiceListExperimentsV1Reader is a Reader for the ExperimentServiceListExperimentsV1 structure.
type ExperimentServiceListExperimentsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExperimentServiceListExperimentsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExperimentServiceListExperimentsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewExperimentServiceListExperimentsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExperimentServiceListExperimentsV1OK creates a ExperimentServiceListExperimentsV1OK with default headers values
func NewExperimentServiceListExperimentsV1OK() *ExperimentServiceListExperimentsV1OK {
	return &ExperimentServiceListExperimentsV1OK{}
}

/*ExperimentServiceListExperimentsV1OK handles this case with default header values.

A successful response.
*/
type ExperimentServiceListExperimentsV1OK struct {
	Payload *experiment_model.APIListExperimentsResponse
}

func (o *ExperimentServiceListExperimentsV1OK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/experiments][%d] experimentServiceListExperimentsV1OK  %+v", 200, o.Payload)
}

func (o *ExperimentServiceListExperimentsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(experiment_model.APIListExperimentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExperimentServiceListExperimentsV1Default creates a ExperimentServiceListExperimentsV1Default with default headers values
func NewExperimentServiceListExperimentsV1Default(code int) *ExperimentServiceListExperimentsV1Default {
	return &ExperimentServiceListExperimentsV1Default{
		_statusCode: code,
	}
}

/*ExperimentServiceListExperimentsV1Default handles this case with default header values.

An unexpected error response.
*/
type ExperimentServiceListExperimentsV1Default struct {
	_statusCode int

	Payload *experiment_model.GatewayruntimeError
}

// Code gets the status code for the experiment service list experiments v1 default response
func (o *ExperimentServiceListExperimentsV1Default) Code() int {
	return o._statusCode
}

func (o *ExperimentServiceListExperimentsV1Default) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/experiments][%d] ExperimentService_ListExperimentsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ExperimentServiceListExperimentsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(experiment_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
