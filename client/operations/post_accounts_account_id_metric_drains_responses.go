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

// PostAccountsAccountIDMetricDrainsReader is a Reader for the PostAccountsAccountIDMetricDrains structure.
type PostAccountsAccountIDMetricDrainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsAccountIDMetricDrainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsAccountIDMetricDrainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountsAccountIDMetricDrainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountsAccountIDMetricDrainsCreated creates a PostAccountsAccountIDMetricDrainsCreated with default headers values
func NewPostAccountsAccountIDMetricDrainsCreated() *PostAccountsAccountIDMetricDrainsCreated {
	return &PostAccountsAccountIDMetricDrainsCreated{}
}

/*PostAccountsAccountIDMetricDrainsCreated handles this case with default header values.

successful
*/
type PostAccountsAccountIDMetricDrainsCreated struct {
}

func (o *PostAccountsAccountIDMetricDrainsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/metric_drains][%d] postAccountsAccountIdMetricDrainsCreated ", 201)
}

func (o *PostAccountsAccountIDMetricDrainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAccountsAccountIDMetricDrainsDefault creates a PostAccountsAccountIDMetricDrainsDefault with default headers values
func NewPostAccountsAccountIDMetricDrainsDefault(code int) *PostAccountsAccountIDMetricDrainsDefault {
	return &PostAccountsAccountIDMetricDrainsDefault{
		_statusCode: code,
	}
}

/*PostAccountsAccountIDMetricDrainsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAccountsAccountIDMetricDrainsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post accounts account ID metric drains default response
func (o *PostAccountsAccountIDMetricDrainsDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountsAccountIDMetricDrainsDefault) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/metric_drains][%d] PostAccountsAccountIDMetricDrains default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountsAccountIDMetricDrainsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAccountsAccountIDMetricDrainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
