// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nzin/golang-skeleton/swagger_gen/models"
)

// CreateTodoOKCode is the HTTP code returned for type CreateTodoOK
const CreateTodoOKCode int = 200

/*CreateTodoOK returns the created todo

swagger:response createTodoOK
*/
type CreateTodoOK struct {

	/*
	  In: Body
	*/
	Payload *CreateTodoOKBody `json:"body,omitempty"`
}

// NewCreateTodoOK creates CreateTodoOK with default headers values
func NewCreateTodoOK() *CreateTodoOK {

	return &CreateTodoOK{}
}

// WithPayload adds the payload to the create todo o k response
func (o *CreateTodoOK) WithPayload(payload *CreateTodoOKBody) *CreateTodoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo o k response
func (o *CreateTodoOK) SetPayload(payload *CreateTodoOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTodoDefault generic error response

swagger:response createTodoDefault
*/
type CreateTodoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTodoDefault creates CreateTodoDefault with default headers values
func NewCreateTodoDefault(code int) *CreateTodoDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTodoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create todo default response
func (o *CreateTodoDefault) WithStatusCode(code int) *CreateTodoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create todo default response
func (o *CreateTodoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create todo default response
func (o *CreateTodoDefault) WithPayload(payload *models.Error) *CreateTodoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo default response
func (o *CreateTodoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
