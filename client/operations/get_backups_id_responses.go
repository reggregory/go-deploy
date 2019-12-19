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

// GetBackupsIDReader is a Reader for the GetBackupsID structure.
type GetBackupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupsIDOK creates a GetBackupsIDOK with default headers values
func NewGetBackupsIDOK() *GetBackupsIDOK {
	return &GetBackupsIDOK{}
}

/*GetBackupsIDOK handles this case with default header values.

successful
*/
type GetBackupsIDOK struct {
	Payload *models.InlineResponse2007
}

func (o *GetBackupsIDOK) Error() string {
	return fmt.Sprintf("[GET /backups/{id}][%d] getBackupsIdOK  %+v", 200, o.Payload)
}

func (o *GetBackupsIDOK) GetPayload() *models.InlineResponse2007 {
	return o.Payload
}

func (o *GetBackupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2007)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupsIDDefault creates a GetBackupsIDDefault with default headers values
func NewGetBackupsIDDefault(code int) *GetBackupsIDDefault {
	return &GetBackupsIDDefault{
		_statusCode: code,
	}
}

/*GetBackupsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetBackupsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get backups ID default response
func (o *GetBackupsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupsIDDefault) Error() string {
	return fmt.Sprintf("[GET /backups/{id}][%d] GetBackupsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackupsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetBackupsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
