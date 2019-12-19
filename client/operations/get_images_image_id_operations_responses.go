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

// GetImagesImageIDOperationsReader is a Reader for the GetImagesImageIDOperations structure.
type GetImagesImageIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesImageIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImagesImageIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetImagesImageIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImagesImageIDOperationsOK creates a GetImagesImageIDOperationsOK with default headers values
func NewGetImagesImageIDOperationsOK() *GetImagesImageIDOperationsOK {
	return &GetImagesImageIDOperationsOK{}
}

/*GetImagesImageIDOperationsOK handles this case with default header values.

successful
*/
type GetImagesImageIDOperationsOK struct {
	Payload *models.InlineResponse20036
}

func (o *GetImagesImageIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /images/{image_id}/operations][%d] getImagesImageIdOperationsOK  %+v", 200, o.Payload)
}

func (o *GetImagesImageIDOperationsOK) GetPayload() *models.InlineResponse20036 {
	return o.Payload
}

func (o *GetImagesImageIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20036)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesImageIDOperationsDefault creates a GetImagesImageIDOperationsDefault with default headers values
func NewGetImagesImageIDOperationsDefault(code int) *GetImagesImageIDOperationsDefault {
	return &GetImagesImageIDOperationsDefault{
		_statusCode: code,
	}
}

/*GetImagesImageIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetImagesImageIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get images image ID operations default response
func (o *GetImagesImageIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetImagesImageIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /images/{image_id}/operations][%d] GetImagesImageIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesImageIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetImagesImageIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
