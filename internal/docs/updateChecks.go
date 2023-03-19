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
//       200: empty
//       500: genericError

// List of printed checks ids
// swagger:parameters updateChecksStatus
type UpdateChecksListIDs struct {
	// The IDs of a checks set status printed
	//
	// in:query
	// required: true
	IDs []int `json:"id"`
}

// Just empty responce body
// swagger:response empty
type Empty struct {
	// in: body
	Body struct{}
}
