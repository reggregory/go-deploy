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

// GetDisksIDReader is a Reader for the GetDisksID structure.
type GetDisksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDisksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDisksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDisksIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDisksIDOK creates a GetDisksIDOK with default headers values
func NewGetDisksIDOK() *GetDisksIDOK {
	return &GetDisksIDOK{}
}

/*GetDisksIDOK handles this case with default header values.

successful
*/
type GetDisksIDOK struct {
	Payload *models.InlineResponse20021
}

func (o *GetDisksIDOK) Error() string {
	return fmt.Sprintf("[GET /disks/{id}][%d] getDisksIdOK  %+v", 200, o.Payload)
}

func (o *GetDisksIDOK) GetPayload() *models.InlineResponse20021 {
	return o.Payload
}

func (o *GetDisksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20021)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDisksIDDefault creates a GetDisksIDDefault with default headers values
func NewGetDisksIDDefault(code int) *GetDisksIDDefault {
	return &GetDisksIDDefault{
		_statusCode: code,
	}
}

/*GetDisksIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetDisksIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get disks ID default response
func (o *GetDisksIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDisksIDDefault) Error() string {
	return fmt.Sprintf("[GET /disks/{id}][%d] GetDisksID default  %+v", o._statusCode, o.Payload)
}

func (o *GetDisksIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetDisksIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
