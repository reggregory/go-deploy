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

// GetBackupsBackupIDOperationsReader is a Reader for the GetBackupsBackupIDOperations structure.
type GetBackupsBackupIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupsBackupIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupsBackupIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupsBackupIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupsBackupIDOperationsOK creates a GetBackupsBackupIDOperationsOK with default headers values
func NewGetBackupsBackupIDOperationsOK() *GetBackupsBackupIDOperationsOK {
	return &GetBackupsBackupIDOperationsOK{}
}

/*GetBackupsBackupIDOperationsOK handles this case with default header values.

successful
*/
type GetBackupsBackupIDOperationsOK struct {
	Payload *models.InlineResponse20036
}

func (o *GetBackupsBackupIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /backups/{backup_id}/operations][%d] getBackupsBackupIdOperationsOK  %+v", 200, o.Payload)
}

func (o *GetBackupsBackupIDOperationsOK) GetPayload() *models.InlineResponse20036 {
	return o.Payload
}

func (o *GetBackupsBackupIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20036)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupsBackupIDOperationsDefault creates a GetBackupsBackupIDOperationsDefault with default headers values
func NewGetBackupsBackupIDOperationsDefault(code int) *GetBackupsBackupIDOperationsDefault {
	return &GetBackupsBackupIDOperationsDefault{
		_statusCode: code,
	}
}

/*GetBackupsBackupIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetBackupsBackupIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get backups backup ID operations default response
func (o *GetBackupsBackupIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupsBackupIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /backups/{backup_id}/operations][%d] GetBackupsBackupIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackupsBackupIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetBackupsBackupIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
