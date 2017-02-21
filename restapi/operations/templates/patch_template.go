package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/rackn/rocket-skates/models"
)

// PatchTemplateHandlerFunc turns a function with the right signature into a patch template handler
type PatchTemplateHandlerFunc func(PatchTemplateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchTemplateHandlerFunc) Handle(params PatchTemplateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PatchTemplateHandler interface for that can handle valid patch template params
type PatchTemplateHandler interface {
	Handle(PatchTemplateParams, *models.Principal) middleware.Responder
}

// NewPatchTemplate creates a new http.Handler for the patch template operation
func NewPatchTemplate(ctx *middleware.Context, handler PatchTemplateHandler) *PatchTemplate {
	return &PatchTemplate{Context: ctx, Handler: handler}
}

/*PatchTemplate swagger:route PATCH /template/{uuid} Templates patchTemplate

Patch Template

*/
type PatchTemplate struct {
	Context *middleware.Context
	Handler PatchTemplateHandler
}

func (o *PatchTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchTemplateParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}