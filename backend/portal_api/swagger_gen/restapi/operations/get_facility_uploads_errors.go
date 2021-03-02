// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetFacilityUploadsErrorsHandlerFunc turns a function with the right signature into a get facility uploads errors handler
type GetFacilityUploadsErrorsHandlerFunc func(GetFacilityUploadsErrorsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFacilityUploadsErrorsHandlerFunc) Handle(params GetFacilityUploadsErrorsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetFacilityUploadsErrorsHandler interface for that can handle valid get facility uploads errors params
type GetFacilityUploadsErrorsHandler interface {
	Handle(GetFacilityUploadsErrorsParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetFacilityUploadsErrors creates a new http.Handler for the get facility uploads errors operation
func NewGetFacilityUploadsErrors(ctx *middleware.Context, handler GetFacilityUploadsErrorsHandler) *GetFacilityUploadsErrors {
	return &GetFacilityUploadsErrors{Context: ctx, Handler: handler}
}

/*GetFacilityUploadsErrors swagger:route GET /facility/uploads/{uploadId}/errors getFacilityUploadsErrors

Get all the error rows associated with given uploadId

*/
type GetFacilityUploadsErrors struct {
	Context *middleware.Context
	Handler GetFacilityUploadsErrorsHandler
}

func (o *GetFacilityUploadsErrors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFacilityUploadsErrorsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
