// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetTodoParams creates a new GetTodoParams object
//
// There are no default values defined in the spec.
func NewGetTodoParams() GetTodoParams {

	return GetTodoParams{}
}

// GetTodoParams contains all the bound params for the get todo operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTodo
type GetTodoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*numeric ID of the todo to get
	  Required: true
	  Minimum: 1
	  In: path
	*/
	TodoID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTodoParams() beforehand.
func (o *GetTodoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTodoID, rhkTodoID, _ := route.Params.GetOK("todoID")
	if err := o.bindTodoID(rTodoID, rhkTodoID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTodoID binds and validates parameter TodoID from path.
func (o *GetTodoParams) bindTodoID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("todoID", "path", "int64", raw)
	}
	o.TodoID = value

	if err := o.validateTodoID(formats); err != nil {
		return err
	}

	return nil
}

// validateTodoID carries on validations for parameter TodoID
func (o *GetTodoParams) validateTodoID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("todoID", "path", o.TodoID, 1, false); err != nil {
		return err
	}

	return nil
}