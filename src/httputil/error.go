package httputil

import "github.com/gin-gonic/gin"

// NewError generates and sends a new error with a status code and message
func NewError(ctx *gin.Context, status int, err string) {
	e := HTTPError{Code: status, Message: err}
	ctx.JSON(status, e)
}

// HTTPError is the struct for errors
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}
