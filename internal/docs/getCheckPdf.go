package docs

import (
	"os"
)

// swagger:route GET /api/v1/check/{check_id}/pdf check getPDF
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
//     Security:
//        - Bearer:
//
//     Deprecated: false
//
//     Responses:
//       200: FileResponse
//       500: genericError

// swagger:parameters getPDF
type _ struct {
	// The ID of a check
	// in:path
	CheckID int `json:"check_id"`
}

// Pdf file of check
// swagger:response FileResponse
type FileResponse struct {
	// In: body
	// swagger:file
	File os.File
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
