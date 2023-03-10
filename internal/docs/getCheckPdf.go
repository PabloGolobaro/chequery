package docs

import (
	"github.com/go-openapi/runtime"
)

// swagger:route GET /check/{check_id}/pdf check getPDF
//
// Get pdf file for check.
//
// This will download pdf file of explicit check.
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
//       200: PdfFile
//       500: genericError

// swagger:parameters getPDF
type _ struct {
	// The ID of a check
	// in:path
	CheckId string `json:"check_id"`
}

// Pdf file of check
// swagger:response PdfFile
type Pdf struct {
	// Pdf file of check
	//
	// In: body
	// Example: -
	File runtime.File
}

// Generic server error
// swagger:response genericError
type DefaultError struct {
	// The error message.
	// in: body
	Body struct {
		// The error message
		//
		// Required: true
		// Example: Error loading file
		Message string
	}
}
