package httputil

// HTTPError is the struct for errors (this is for swagger doc purposes)
type HTTPError struct {
	Message string `json:"message" example:"Bad Request"`
}
