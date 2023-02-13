// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nzin/golang-skeleton/swagger_gen/models"
)

// GetTodoOKCode is the HTTP code returned for type GetTodoOK
const GetTodoOKCode int = 200

/*GetTodoOK returns the todo

swagger:response getTodoOK
*/
type GetTodoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Todo `json:"body,omitempty"`
}

// NewGetTodoOK creates GetTodoOK with default headers values
func NewGetTodoOK() *GetTodoOK {

	return &GetTodoOK{}
}

// WithPayload adds the payload to the get todo o k response
func (o *GetTodoOK) WithPayload(payload *models.Todo) *GetTodoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo o k response
func (o *GetTodoOK) SetPayload(payload *models.Todo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTodoDefault generic error response

swagger:response getTodoDefault
*/
type GetTodoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTodoDefault creates GetTodoDefault with default headers values
func NewGetTodoDefault(code int) *GetTodoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTodoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get todo default response
func (o *GetTodoDefault) WithStatusCode(code int) *GetTodoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get todo default response
func (o *GetTodoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get todo default response
func (o *GetTodoDefault) WithPayload(payload *models.Error) *GetTodoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo default response
func (o *GetTodoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}