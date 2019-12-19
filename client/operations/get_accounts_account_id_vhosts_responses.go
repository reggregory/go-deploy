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

// GetAccountsAccountIDVhostsReader is a Reader for the GetAccountsAccountIDVhosts structure.
type GetAccountsAccountIDVhostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDVhostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDVhostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountsAccountIDVhostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountsAccountIDVhostsOK creates a GetAccountsAccountIDVhostsOK with default headers values
func NewGetAccountsAccountIDVhostsOK() *GetAccountsAccountIDVhostsOK {
	return &GetAccountsAccountIDVhostsOK{}
}

/*GetAccountsAccountIDVhostsOK handles this case with default header values.

successful
*/
type GetAccountsAccountIDVhostsOK struct {
	Payload *models.InlineResponse20047
}

func (o *GetAccountsAccountIDVhostsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/vhosts][%d] getAccountsAccountIdVhostsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDVhostsOK) GetPayload() *models.InlineResponse20047 {
	return o.Payload
}

func (o *GetAccountsAccountIDVhostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20047)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDVhostsDefault creates a GetAccountsAccountIDVhostsDefault with default headers values
func NewGetAccountsAccountIDVhostsDefault(code int) *GetAccountsAccountIDVhostsDefault {
	return &GetAccountsAccountIDVhostsDefault{
		_statusCode: code,
	}
}

/*GetAccountsAccountIDVhostsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAccountsAccountIDVhostsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get accounts account ID vhosts default response
func (o *GetAccountsAccountIDVhostsDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountsAccountIDVhostsDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/vhosts][%d] GetAccountsAccountIDVhosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountsAccountIDVhostsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAccountsAccountIDVhostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
