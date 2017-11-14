package api_handlers

// A ValidationError is an error that is used when the required input fails validation.
// swagger:parameters listParams
type listParams struct {
	// in: path
	BotID    string

}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response errorResponse
type errorResponse struct {
	ErrorCode    int
	ErrorMessage string
}
