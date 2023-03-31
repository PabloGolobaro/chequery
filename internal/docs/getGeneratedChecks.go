package docs

import "github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"

// swagger:route GET /api/v1/check/generated check getGenerated
//
// Get a list of IDs of generated checks.
//
// This will return list of check IDs.
//
//     Schemes: http
//
//     Deprecated: false
//
//     Security:
//        - Bearer:
//
//     Responses:
//       200: generatedIds
//       500: genericError

// List of generated checks ids
// swagger:response generatedIds
type WrapperGeneratedIds struct {
	// in: body
	Body check.GeneratedChecksResponse
}
