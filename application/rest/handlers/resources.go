package handlers

type CreateResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"code"`
}

type UpdateResponse struct {
	Error     string "json:error"
	ErrorCode int    "json:code"
}
