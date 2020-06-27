package httputil

// HTTPError is the struct for errors (this is for swagger doc purposes)
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}
