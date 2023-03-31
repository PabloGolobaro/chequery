package docs

import "github.com/pablogolobaro/chequery/internal/handlers/rest/v1/order"

// swagger:route POST /api/v1/order order createOrder
//
// Create checks from new order.
//
// This will create new checks in db and starts generating pdf files from it.
//
//     Schemes: http
//
//     Deprecated: false
//
//     Security:
//        - Bearer:
//
//     Responses:
//       201: orderCreateResponse
//       default: genericError

// JSON of unknown type with order information
// swagger:parameters createOrder
type WrapperOrderCreateRequest struct {
	// in: body
	Body order.OrderCreateRequest
}

// List of created checks IDs
// swagger:response orderCreateResponse
type WrappedOrderCreateResponse struct {
	// in: body
	Body order.OrderCreateResponse
}
