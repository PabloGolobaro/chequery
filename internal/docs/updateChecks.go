package docs

// swagger:route PUT /check check updateChecksStatus
//
// Update status of some checks to be printed.
//
// This will update status of checks in DB to "printed".
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       500: defaultError

// swagger:parameters updateChecksStatus
type UpdateChecksListIDs struct {
	// The IDs of a checks set status printed
	//
	// in:query
	IDs []string `json:"id"`
}
