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

// PostImagesImageIDOperationsReader is a Reader for the PostImagesImageIDOperations structure.
type PostImagesImageIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostImagesImageIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostImagesImageIDOperationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostImagesImageIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostImagesImageIDOperationsCreated creates a PostImagesImageIDOperationsCreated with default headers values
func NewPostImagesImageIDOperationsCreated() *PostImagesImageIDOperationsCreated {
	return &PostImagesImageIDOperationsCreated{}
}

/*PostImagesImageIDOperationsCreated handles this case with default header values.

successful
*/
type PostImagesImageIDOperationsCreated struct {
}

func (o *PostImagesImageIDOperationsCreated) Error() string {
	return fmt.Sprintf("[POST /images/{image_id}/operations][%d] postImagesImageIdOperationsCreated ", 201)
}

func (o *PostImagesImageIDOperationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostImagesImageIDOperationsDefault creates a PostImagesImageIDOperationsDefault with default headers values
func NewPostImagesImageIDOperationsDefault(code int) *PostImagesImageIDOperationsDefault {
	return &PostImagesImageIDOperationsDefault{
		_statusCode: code,
	}
}

/*PostImagesImageIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostImagesImageIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post images image ID operations default response
func (o *PostImagesImageIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *PostImagesImageIDOperationsDefault) Error() string {
	return fmt.Sprintf("[POST /images/{image_id}/operations][%d] PostImagesImageIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *PostImagesImageIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostImagesImageIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
