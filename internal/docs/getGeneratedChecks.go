package docs

import "github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"

// swagger:route GET /check/generated check getGenerated
//
// Get a list of IDs of generated checks.
//
// This will return list of check IDs.
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       200: generatedIds
//       500: defaultError

// List of generated checks ids
// swagger:response generatedIds
type WrapperGeneratedIds struct {
	// in: body
	Body check.GeneratedChecksResponse
}
