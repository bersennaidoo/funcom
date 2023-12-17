package handlers

type CreateResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"code"`
}
