// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// EnterContainerOKCode is the HTTP code returned for type EnterContainerOK
const EnterContainerOKCode int = 200

/*EnterContainerOK enter to the container

swagger:response enterContainerOK
*/
type EnterContainerOK struct {
}

// NewEnterContainerOK creates EnterContainerOK with default headers values
func NewEnterContainerOK() *EnterContainerOK {

	return &EnterContainerOK{}
}

// WriteResponse to the client
func (o *EnterContainerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
