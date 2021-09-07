package handlers

import (
	"strconv"

	"github.com/jgoralcz/cdbapi/src/db/images"
	"github.com/labstack/echo/v4"
)

// ImageByID is a handler for echo that gets an image's metadata based off the ID provided.
// @Summary Gets an image's metadata based off the provided ID.
// @Description Get detailed information on an image metadata based off the provided ID.
// @Produce json
// @Param id path int true "Some ID"
// @Success 200 {array} images.Image
// @Failure 400 {object} httputil.HTTPError "Must have a valid id parameter"
// @Failure 404 {object} httputil.HTTPError "Image not found with id some ID"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the images for the character"
// @Router /v1/images/{id} [get]
// @Tags Images
func ImageByID(c echo.Context) (err error) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)

	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid id parameter"}
	}

	json := images.GetImageByID(id)

	if json == "{}" {
		return &echo.HTTPError{Code: 404, Message: "Image not found with id " + strID}
	}

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the images for the character"}
	}

	return c.JSON(200, json)
}
