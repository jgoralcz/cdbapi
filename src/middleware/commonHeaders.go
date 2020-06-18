package middleware

import "net/http"

// CommonHeaders adds basic headers to every request.
// An example is the "Content-Type" header with the value "application/json"
func CommonHeaders(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	rw.Header().Add("Content-Type", "application/json")
	next(rw, req)
}
