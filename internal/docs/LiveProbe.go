package docs

import "github.com/pablogolobaro/chequery/internal/handlers/rest/v1/health"

// swagger:route GET /health-check health liveProbe
//
// Get live probe.
//
// This will ensure that server router is working.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/pdf
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       200: liveProbeSuccess
//       500:

// Generic server live probe message
// swagger:response liveProbeSuccess
type LiveResponce struct {
	// The router live message.
	// in: body
	Body health.LiveProbeResponse
}
