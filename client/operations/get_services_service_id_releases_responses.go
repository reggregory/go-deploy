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

// GetServicesServiceIDReleasesReader is a Reader for the GetServicesServiceIDReleases structure.
type GetServicesServiceIDReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesServiceIDReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesServiceIDReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesServiceIDReleasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesServiceIDReleasesOK creates a GetServicesServiceIDReleasesOK with default headers values
func NewGetServicesServiceIDReleasesOK() *GetServicesServiceIDReleasesOK {
	return &GetServicesServiceIDReleasesOK{}
}

/*GetServicesServiceIDReleasesOK handles this case with default header values.

successful
*/
type GetServicesServiceIDReleasesOK struct {
	Payload *models.InlineResponse20039
}

func (o *GetServicesServiceIDReleasesOK) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/releases][%d] getServicesServiceIdReleasesOK  %+v", 200, o.Payload)
}

func (o *GetServicesServiceIDReleasesOK) GetPayload() *models.InlineResponse20039 {
	return o.Payload
}

func (o *GetServicesServiceIDReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20039)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesServiceIDReleasesDefault creates a GetServicesServiceIDReleasesDefault with default headers values
func NewGetServicesServiceIDReleasesDefault(code int) *GetServicesServiceIDReleasesDefault {
	return &GetServicesServiceIDReleasesDefault{
		_statusCode: code,
	}
}

/*GetServicesServiceIDReleasesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetServicesServiceIDReleasesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get services service ID releases default response
func (o *GetServicesServiceIDReleasesDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesServiceIDReleasesDefault) Error() string {
	return fmt.Sprintf("[GET /services/{service_id}/releases][%d] GetServicesServiceIDReleases default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesServiceIDReleasesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetServicesServiceIDReleasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
