// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// UpdateOAuth2ClientPublicReader is a Reader for the UpdateOAuth2ClientPublic structure.
type UpdateOAuth2ClientPublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOAuth2ClientPublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOAuth2ClientPublicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateOAuth2ClientPublicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOAuth2ClientPublicOK creates a UpdateOAuth2ClientPublicOK with default headers values
func NewUpdateOAuth2ClientPublicOK() *UpdateOAuth2ClientPublicOK {
	return &UpdateOAuth2ClientPublicOK{}
}

/* UpdateOAuth2ClientPublicOK describes a response with status code 200, with default header values.

oAuth2Client
*/
type UpdateOAuth2ClientPublicOK struct {
	Payload *models.OAuth2Client
}

func (o *UpdateOAuth2ClientPublicOK) Error() string {
	return fmt.Sprintf("[PUT /connect/register][%d] updateOAuth2ClientPublicOK  %+v", 200, o.Payload)
}
func (o *UpdateOAuth2ClientPublicOK) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *UpdateOAuth2ClientPublicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOAuth2ClientPublicInternalServerError creates a UpdateOAuth2ClientPublicInternalServerError with default headers values
func NewUpdateOAuth2ClientPublicInternalServerError() *UpdateOAuth2ClientPublicInternalServerError {
	return &UpdateOAuth2ClientPublicInternalServerError{}
}

/* UpdateOAuth2ClientPublicInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type UpdateOAuth2ClientPublicInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateOAuth2ClientPublicInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /connect/register][%d] updateOAuth2ClientPublicInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateOAuth2ClientPublicInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateOAuth2ClientPublicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
