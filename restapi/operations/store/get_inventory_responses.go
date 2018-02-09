// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetInventoryOKCode is the HTTP code returned for type GetInventoryOK
const GetInventoryOKCode int = 200

/*GetInventoryOK successful operation

swagger:response getInventoryOK
*/
type GetInventoryOK struct {

	/*
	  In: Body
	*/
	Payload GetInventoryOKBody `json:"body,omitempty"`
}

// NewGetInventoryOK creates GetInventoryOK with default headers values
func NewGetInventoryOK() *GetInventoryOK {
	return &GetInventoryOK{}
}

// WithPayload adds the payload to the get inventory o k response
func (o *GetInventoryOK) WithPayload(payload GetInventoryOKBody) *GetInventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get inventory o k response
func (o *GetInventoryOK) SetPayload(payload GetInventoryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}