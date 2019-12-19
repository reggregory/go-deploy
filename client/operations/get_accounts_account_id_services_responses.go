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

// GetAccountsAccountIDServicesReader is a Reader for the GetAccountsAccountIDServices structure.
type GetAccountsAccountIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountsAccountIDServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountsAccountIDServicesOK creates a GetAccountsAccountIDServicesOK with default headers values
func NewGetAccountsAccountIDServicesOK() *GetAccountsAccountIDServicesOK {
	return &GetAccountsAccountIDServicesOK{}
}

/*GetAccountsAccountIDServicesOK handles this case with default header values.

successful
*/
type GetAccountsAccountIDServicesOK struct {
	Payload *models.InlineResponse20041
}

func (o *GetAccountsAccountIDServicesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/services][%d] getAccountsAccountIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDServicesOK) GetPayload() *models.InlineResponse20041 {
	return o.Payload
}

func (o *GetAccountsAccountIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20041)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDServicesDefault creates a GetAccountsAccountIDServicesDefault with default headers values
func NewGetAccountsAccountIDServicesDefault(code int) *GetAccountsAccountIDServicesDefault {
	return &GetAccountsAccountIDServicesDefault{
		_statusCode: code,
	}
}

/*GetAccountsAccountIDServicesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAccountsAccountIDServicesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get accounts account ID services default response
func (o *GetAccountsAccountIDServicesDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountsAccountIDServicesDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/services][%d] GetAccountsAccountIDServices default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountsAccountIDServicesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAccountsAccountIDServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
