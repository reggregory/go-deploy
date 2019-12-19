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

// PostAccountsAccountIDAppsReader is a Reader for the PostAccountsAccountIDApps structure.
type PostAccountsAccountIDAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsAccountIDAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsAccountIDAppsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountsAccountIDAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountsAccountIDAppsCreated creates a PostAccountsAccountIDAppsCreated with default headers values
func NewPostAccountsAccountIDAppsCreated() *PostAccountsAccountIDAppsCreated {
	return &PostAccountsAccountIDAppsCreated{}
}

/*PostAccountsAccountIDAppsCreated handles this case with default header values.

successful
*/
type PostAccountsAccountIDAppsCreated struct {
}

func (o *PostAccountsAccountIDAppsCreated) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/apps][%d] postAccountsAccountIdAppsCreated ", 201)
}

func (o *PostAccountsAccountIDAppsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAccountsAccountIDAppsDefault creates a PostAccountsAccountIDAppsDefault with default headers values
func NewPostAccountsAccountIDAppsDefault(code int) *PostAccountsAccountIDAppsDefault {
	return &PostAccountsAccountIDAppsDefault{
		_statusCode: code,
	}
}

/*PostAccountsAccountIDAppsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAccountsAccountIDAppsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post accounts account ID apps default response
func (o *PostAccountsAccountIDAppsDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountsAccountIDAppsDefault) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/apps][%d] PostAccountsAccountIDApps default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountsAccountIDAppsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAccountsAccountIDAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
