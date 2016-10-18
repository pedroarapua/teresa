package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

Get users
*/
type GetUsersOK struct {
	Payload GetUsersOKBodyBody
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*GetUsersBadRequest handles this case with default header values.

User sent a bad request
*/
type GetUsersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*GetUsersUnauthorized handles this case with default header values.

User not authorized
*/
type GetUsersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*GetUsersForbidden handles this case with default header values.

User does not have the credentials to access this resource

*/
type GetUsersForbidden struct {
	Payload *models.Unauthorized
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersNotFound creates a GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {
	return &GetUsersNotFound{}
}

/*GetUsersNotFound handles this case with default header values.

Resource not found
*/
type GetUsersNotFound struct {
	Payload *models.NotFound
}

func (o *GetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDefault creates a GetUsersDefault with default headers values
func NewGetUsersDefault(code int) *GetUsersDefault {
	return &GetUsersDefault{
		_statusCode: code,
	}
}

/*GetUsersDefault handles this case with default header values.

Error
*/
type GetUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get users default response
func (o *GetUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsers default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUsersOKBodyBody get users o k body body

swagger:model GetUsersOKBodyBody
*/
type GetUsersOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.User `json:"items"`
}

// Validate validates this get users o k body body
func (o *GetUsersOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getUsersOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {

		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {

			if err := o.Items[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}