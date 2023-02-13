// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/nzin/golang-skeleton/swagger_gen/models"
)

// ListTodosHandlerFunc turns a function with the right signature into a list todos handler
type ListTodosHandlerFunc func(ListTodosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTodosHandlerFunc) Handle(params ListTodosParams) middleware.Responder {
	return fn(params)
}

// ListTodosHandler interface for that can handle valid list todos params
type ListTodosHandler interface {
	Handle(ListTodosParams) middleware.Responder
}

// NewListTodos creates a new http.Handler for the list todos operation
func NewListTodos(ctx *middleware.Context, handler ListTodosHandler) *ListTodos {
	return &ListTodos{Context: ctx, Handler: handler}
}

/* ListTodos swagger:route GET /todos app listTodos

App: List todos

*/
type ListTodos struct {
	Context *middleware.Context
	Handler ListTodosHandler
}

func (o *ListTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTodosParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListTodosOKBody list todos o k body
//
// swagger:model ListTodosOKBody
type ListTodosOKBody struct {

	// todos
	Todos []*models.Todo `json:"todos"`
}

// Validate validates this list todos o k body
func (o *ListTodosOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTodos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListTodosOKBody) validateTodos(formats strfmt.Registry) error {
	if swag.IsZero(o.Todos) { // not required
		return nil
	}

	for i := 0; i < len(o.Todos); i++ {
		if swag.IsZero(o.Todos[i]) { // not required
			continue
		}

		if o.Todos[i] != nil {
			if err := o.Todos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listTodosOK" + "." + "todos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listTodosOK" + "." + "todos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list todos o k body based on the context it is used
func (o *ListTodosOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTodos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListTodosOKBody) contextValidateTodos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Todos); i++ {

		if o.Todos[i] != nil {
			if err := o.Todos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listTodosOK" + "." + "todos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listTodosOK" + "." + "todos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListTodosOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListTodosOKBody) UnmarshalBinary(b []byte) error {
	var res ListTodosOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}