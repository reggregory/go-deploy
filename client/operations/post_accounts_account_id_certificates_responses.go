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

// PostAccountsAccountIDCertificatesReader is a Reader for the PostAccountsAccountIDCertificates structure.
type PostAccountsAccountIDCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountsAccountIDCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountsAccountIDCertificatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountsAccountIDCertificatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountsAccountIDCertificatesCreated creates a PostAccountsAccountIDCertificatesCreated with default headers values
func NewPostAccountsAccountIDCertificatesCreated() *PostAccountsAccountIDCertificatesCreated {
	return &PostAccountsAccountIDCertificatesCreated{}
}

/*PostAccountsAccountIDCertificatesCreated handles this case with default header values.

successful
*/
type PostAccountsAccountIDCertificatesCreated struct {
}

func (o *PostAccountsAccountIDCertificatesCreated) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/certificates][%d] postAccountsAccountIdCertificatesCreated ", 201)
}

func (o *PostAccountsAccountIDCertificatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAccountsAccountIDCertificatesDefault creates a PostAccountsAccountIDCertificatesDefault with default headers values
func NewPostAccountsAccountIDCertificatesDefault(code int) *PostAccountsAccountIDCertificatesDefault {
	return &PostAccountsAccountIDCertificatesDefault{
		_statusCode: code,
	}
}

/*PostAccountsAccountIDCertificatesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type PostAccountsAccountIDCertificatesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the post accounts account ID certificates default response
func (o *PostAccountsAccountIDCertificatesDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountsAccountIDCertificatesDefault) Error() string {
	return fmt.Sprintf("[POST /accounts/{account_id}/certificates][%d] PostAccountsAccountIDCertificates default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountsAccountIDCertificatesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *PostAccountsAccountIDCertificatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}