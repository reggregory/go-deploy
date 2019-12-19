// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// GetEphemeralSessionsIDReader is a Reader for the GetEphemeralSessionsID structure.
type GetEphemeralSessionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEphemeralSessionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEphemeralSessionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEphemeralSessionsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEphemeralSessionsIDOK creates a GetEphemeralSessionsIDOK with default headers values
func NewGetEphemeralSessionsIDOK() *GetEphemeralSessionsIDOK {
	return &GetEphemeralSessionsIDOK{}
}

/*GetEphemeralSessionsIDOK handles this case with default header values.

successful
*/
type GetEphemeralSessionsIDOK struct {
	Payload *models.InlineResponse20025
}

func (o *GetEphemeralSessionsIDOK) Error() string {
	return fmt.Sprintf("[GET /ephemeral_sessions/{id}][%d] getEphemeralSessionsIdOK  %+v", 200, o.Payload)
}

func (o *GetEphemeralSessionsIDOK) GetPayload() *models.InlineResponse20025 {
	return o.Payload
}

func (o *GetEphemeralSessionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20025)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEphemeralSessionsIDDefault creates a GetEphemeralSessionsIDDefault with default headers values
func NewGetEphemeralSessionsIDDefault(code int) *GetEphemeralSessionsIDDefault {
	return &GetEphemeralSessionsIDDefault{
		_statusCode: code,
	}
}

/*GetEphemeralSessionsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetEphemeralSessionsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get ephemeral sessions ID default response
func (o *GetEphemeralSessionsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetEphemeralSessionsIDDefault) Error() string {
	return fmt.Sprintf("[GET /ephemeral_sessions/{id}][%d] GetEphemeralSessionsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetEphemeralSessionsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetEphemeralSessionsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
