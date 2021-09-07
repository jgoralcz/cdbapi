package handlers

import (
	_ "github.com/jgoralcz/cdbapi/src/lib/structs" // For Swagger purposes
	"github.com/labstack/echo/v4"
)

// Healthz is a simple response to see if the api is working.
// @Summary Gets 200 OK
// @Description checks if API is working
// @Success 200 {string} OK
// @Router /healthz [get]
// @Tags Health Check
func Healthz(c echo.Context) (err error) {
	return c.String(200, "OK")
}
