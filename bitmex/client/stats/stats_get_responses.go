// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// StatsGetReader is a Reader for the StatsGet structure.
type StatsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewStatsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewStatsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStatsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatsGetOK creates a StatsGetOK with default headers values
func NewStatsGetOK() *StatsGetOK {
	return &StatsGetOK{}
}

/*StatsGetOK handles this case with default header values.

Request was successful
*/
type StatsGetOK struct {
	Payload []*models.Stats
}

func (o *StatsGetOK) Error() string {
	return fmt.Sprintf("[GET /stats][%d] statsGetOK  %+v", 200, o.Payload)
}

func (o *StatsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsGetBadRequest creates a StatsGetBadRequest with default headers values
func NewStatsGetBadRequest() *StatsGetBadRequest {
	return &StatsGetBadRequest{}
}

/*StatsGetBadRequest handles this case with default header values.

Parameter Error
*/
type StatsGetBadRequest struct {
	Payload *models.Error
}

func (o *StatsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /stats][%d] statsGetBadRequest  %+v", 400, *o.Payload)
}

func (o *StatsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsGetUnauthorized creates a StatsGetUnauthorized with default headers values
func NewStatsGetUnauthorized() *StatsGetUnauthorized {
	return &StatsGetUnauthorized{}
}

/*StatsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type StatsGetUnauthorized struct {
	Payload *models.Error
}

func (o *StatsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /stats][%d] statsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *StatsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsGetNotFound creates a StatsGetNotFound with default headers values
func NewStatsGetNotFound() *StatsGetNotFound {
	return &StatsGetNotFound{}
}

/*StatsGetNotFound handles this case with default header values.

Not Found
*/
type StatsGetNotFound struct {
	Payload *models.Error
}

func (o *StatsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /stats][%d] statsGetNotFound  %+v", 404, o.Payload)
}

func (o *StatsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
