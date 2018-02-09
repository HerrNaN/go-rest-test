// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"test/models"
)

// NewCreateUsersWithArrayInputParams creates a new CreateUsersWithArrayInputParams object
// with the default values initialized.
func NewCreateUsersWithArrayInputParams() CreateUsersWithArrayInputParams {
	var ()
	return CreateUsersWithArrayInputParams{}
}

// CreateUsersWithArrayInputParams contains all the bound params for the create users with array input operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUsersWithArrayInput
type CreateUsersWithArrayInputParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*List of user object
	  Required: true
	  In: body
	*/
	Body models.CreateUsersWithArrayInputParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateUsersWithArrayInputParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateUsersWithArrayInputParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {

			if len(res) == 0 {
				o.Body = body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}