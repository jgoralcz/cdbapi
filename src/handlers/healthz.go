package handlers

import (
	"github.com/jgoralcz/cdbapi/src/db/health"
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

// Healthz is a simple response to see if the database is working.
// @Summary Gets 200 OK
// @Description checks if database is working
// @Success 200 {string} OK
// @Router /healthz/db [get]
// @Tags Health Check
func HealthzDb(c echo.Context) (err error) {
	data, err := health.GetNow()

	if err != nil {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the date from database"}
	}

	return c.JSON(200, data)
}
