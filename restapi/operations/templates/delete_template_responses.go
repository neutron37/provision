package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type DeleteTemplateNoContent
const DeleteTemplateNoContentCode int = 204

/*DeleteTemplateNoContent delete template no content

swagger:response deleteTemplateNoContent
*/
type DeleteTemplateNoContent struct {
}

// NewDeleteTemplateNoContent creates DeleteTemplateNoContent with default headers values
func NewDeleteTemplateNoContent() *DeleteTemplateNoContent {
	return &DeleteTemplateNoContent{}
}

// WriteResponse to the client
func (o *DeleteTemplateNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// HTTP code for type DeleteTemplateUnauthorized
const DeleteTemplateUnauthorizedCode int = 401

/*DeleteTemplateUnauthorized delete template unauthorized

swagger:response deleteTemplateUnauthorized
*/
type DeleteTemplateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateUnauthorized creates DeleteTemplateUnauthorized with default headers values
func NewDeleteTemplateUnauthorized() *DeleteTemplateUnauthorized {
	return &DeleteTemplateUnauthorized{}
}

// WithPayload adds the payload to the delete template unauthorized response
func (o *DeleteTemplateUnauthorized) WithPayload(payload *models.Error) *DeleteTemplateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template unauthorized response
func (o *DeleteTemplateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteTemplateNotFound
const DeleteTemplateNotFoundCode int = 404

/*DeleteTemplateNotFound delete template not found

swagger:response deleteTemplateNotFound
*/
type DeleteTemplateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateNotFound creates DeleteTemplateNotFound with default headers values
func NewDeleteTemplateNotFound() *DeleteTemplateNotFound {
	return &DeleteTemplateNotFound{}
}

// WithPayload adds the payload to the delete template not found response
func (o *DeleteTemplateNotFound) WithPayload(payload *models.Error) *DeleteTemplateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template not found response
func (o *DeleteTemplateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteTemplateConflict
const DeleteTemplateConflictCode int = 409

/*DeleteTemplateConflict delete template conflict

swagger:response deleteTemplateConflict
*/
type DeleteTemplateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateConflict creates DeleteTemplateConflict with default headers values
func NewDeleteTemplateConflict() *DeleteTemplateConflict {
	return &DeleteTemplateConflict{}
}

// WithPayload adds the payload to the delete template conflict response
func (o *DeleteTemplateConflict) WithPayload(payload *models.Error) *DeleteTemplateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template conflict response
func (o *DeleteTemplateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type DeleteTemplateInternalServerError
const DeleteTemplateInternalServerErrorCode int = 500

/*DeleteTemplateInternalServerError delete template internal server error

swagger:response deleteTemplateInternalServerError
*/
type DeleteTemplateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateInternalServerError creates DeleteTemplateInternalServerError with default headers values
func NewDeleteTemplateInternalServerError() *DeleteTemplateInternalServerError {
	return &DeleteTemplateInternalServerError{}
}

// WithPayload adds the payload to the delete template internal server error response
func (o *DeleteTemplateInternalServerError) WithPayload(payload *models.Error) *DeleteTemplateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template internal server error response
func (o *DeleteTemplateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}