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

// GetDatabasesReader is a Reader for the GetDatabases structure.
type GetDatabasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatabasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatabasesOK creates a GetDatabasesOK with default headers values
func NewGetDatabasesOK() *GetDatabasesOK {
	return &GetDatabasesOK{}
}

/*GetDatabasesOK handles this case with default header values.

successful
*/
type GetDatabasesOK struct {
	Payload *models.InlineResponse20018
}

func (o *GetDatabasesOK) Error() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesOK) GetPayload() *models.InlineResponse20018 {
	return o.Payload
}

func (o *GetDatabasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20018)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesDefault creates a GetDatabasesDefault with default headers values
func NewGetDatabasesDefault(code int) *GetDatabasesDefault {
	return &GetDatabasesDefault{
		_statusCode: code,
	}
}

/*GetDatabasesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDatabasesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get databases default response
func (o *GetDatabasesDefault) Code() int {
	return o._statusCode
}

func (o *GetDatabasesDefault) Error() string {
	return fmt.Sprintf("[GET /databases][%d] GetDatabases default  %+v", o._statusCode, o.Payload)
}

func (o *GetDatabasesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDatabasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
