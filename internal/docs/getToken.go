package docs

import "github.com/pablogolobaro/chequery/internal/handlers/auth"

// swagger:route POST /auth authHandlers TokenHandler
//
// Get JWT authorization token.
//
// This will generate JWT Token to use API if username and password are correct.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       200: JWTToken
//       400: genericError
//       401: genericError
//       500: genericError

// swagger:parameters TokenHandler
type TokenHandlerParameters struct {
	// username and password
	// In: body
	Body auth.AuthDataRequest
}

// swagger:response JWTToken
type TokenResponse struct {
	// In: body
	Body auth.AuthDataResponce
}
