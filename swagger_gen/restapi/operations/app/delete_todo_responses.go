// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nzin/golang-skeleton/swagger_gen/models"
)

// DeleteTodoOKCode is the HTTP code returned for type DeleteTodoOK
const DeleteTodoOKCode int = 200

/*DeleteTodoOK OK deleted

swagger:response deleteTodoOK
*/
type DeleteTodoOK struct {
}

// NewDeleteTodoOK creates DeleteTodoOK with default headers values
func NewDeleteTodoOK() *DeleteTodoOK {

	return &DeleteTodoOK{}
}

// WriteResponse to the client
func (o *DeleteTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteTodoDefault generic error response

swagger:response deleteTodoDefault
*/
type DeleteTodoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTodoDefault creates DeleteTodoDefault with default headers values
func NewDeleteTodoDefault(code int) *DeleteTodoDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTodoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete todo default response
func (o *DeleteTodoDefault) WithStatusCode(code int) *DeleteTodoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete todo default response
func (o *DeleteTodoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete todo default response
func (o *DeleteTodoDefault) WithPayload(payload *models.Error) *DeleteTodoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete todo default response
func (o *DeleteTodoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTodoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
