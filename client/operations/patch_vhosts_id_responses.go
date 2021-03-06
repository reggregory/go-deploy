// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/reggregory/go-deploy/models"
)

// PatchVhostsIDReader is a Reader for the PatchVhostsID structure.
type PatchVhostsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVhostsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchVhostsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchVhostsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchVhostsIDOK creates a PatchVhostsIDOK with default headers values
func NewPatchVhostsIDOK() *PatchVhostsIDOK {
	return &PatchVhostsIDOK{}
}

/*PatchVhostsIDOK handles this case with default header values.

successful
*/
type PatchVhostsIDOK struct {
}

func (o *PatchVhostsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /vhosts/{id}][%d] patchVhostsIdOK ", 200)
}

func (o *PatchVhostsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchVhostsIDDefault creates a PatchVhostsIDDefault with default headers values
func NewPatchVhostsIDDefault(code int) *PatchVhostsIDDefault {
	return &PatchVhostsIDDefault{
		_statusCode: code,
	}
}

/*PatchVhostsIDDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PatchVhostsIDDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the patch vhosts ID default response
func (o *PatchVhostsIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchVhostsIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /vhosts/{id}][%d] PatchVhostsID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchVhostsIDDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PatchVhostsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
