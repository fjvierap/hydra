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

// DeleteOAuth2ClientPublicReader is a Reader for the DeleteOAuth2ClientPublic structure.
type DeleteOAuth2ClientPublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOAuth2ClientPublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOAuth2ClientPublicNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteOAuth2ClientPublicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOAuth2ClientPublicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOAuth2ClientPublicNoContent creates a DeleteOAuth2ClientPublicNoContent with default headers values
func NewDeleteOAuth2ClientPublicNoContent() *DeleteOAuth2ClientPublicNoContent {
	return &DeleteOAuth2ClientPublicNoContent{}
}

/* DeleteOAuth2ClientPublicNoContent describes a response with status code 204, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteOAuth2ClientPublicNoContent struct {
}

func (o *DeleteOAuth2ClientPublicNoContent) Error() string {
	return fmt.Sprintf("[DELETE /connect/register][%d] deleteOAuth2ClientPublicNoContent ", 204)
}

func (o *DeleteOAuth2ClientPublicNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOAuth2ClientPublicNotFound creates a DeleteOAuth2ClientPublicNotFound with default headers values
func NewDeleteOAuth2ClientPublicNotFound() *DeleteOAuth2ClientPublicNotFound {
	return &DeleteOAuth2ClientPublicNotFound{}
}

/* DeleteOAuth2ClientPublicNotFound describes a response with status code 404, with default header values.

genericError
*/
type DeleteOAuth2ClientPublicNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2ClientPublicNotFound) Error() string {
	return fmt.Sprintf("[DELETE /connect/register][%d] deleteOAuth2ClientPublicNotFound  %+v", 404, o.Payload)
}
func (o *DeleteOAuth2ClientPublicNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteOAuth2ClientPublicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOAuth2ClientPublicInternalServerError creates a DeleteOAuth2ClientPublicInternalServerError with default headers values
func NewDeleteOAuth2ClientPublicInternalServerError() *DeleteOAuth2ClientPublicInternalServerError {
	return &DeleteOAuth2ClientPublicInternalServerError{}
}

/* DeleteOAuth2ClientPublicInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type DeleteOAuth2ClientPublicInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2ClientPublicInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /connect/register][%d] deleteOAuth2ClientPublicInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteOAuth2ClientPublicInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteOAuth2ClientPublicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
