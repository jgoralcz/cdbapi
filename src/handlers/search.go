package handlers

import (
	"github.com/jgoralcz/cdbapi/src/db/search"
	"github.com/jgoralcz/cdbapi/src/lib/helpers"
	_ "github.com/jgoralcz/cdbapi/src/lib/structs" // For Swagger purposes
	"github.com/labstack/echo/v4"
)

// Search is a handler for echo that gets the series metadata based off the name and user's filters.
// @Summary Gets a series based off the user's query parameters.
// @Description Get series metadata by nsfw (boolean), game (boolean), western (boolean), limit (1-20), name (string). You must use name to get a result back.
// @Produce json
// @Param name query string true "name to search"
// @Param limit query int false "limit 1-20; Default 1"
// @Param nsfw query boolean false "whether the series is nsfw or not"
// @Param western query boolean false "whether the series is western (Cartoon) or not (Anime)"
// @Param game query boolean false "whether the series is from a game or not"
// @Success 200 {array} series.Series
// @Failure 400 {object} httputil.HTTPError "Must have a valid name query parameter"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the series"
// @Router /series [get]
// @Tags Series
func Search(c echo.Context) (err error) {
	initLimit := c.QueryParam("limit")
	nsfw := c.QueryParam("nsfw")
	western := c.QueryParam("western")
	game := c.QueryParam("game")
	name := c.QueryParam("name")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	if name == "" {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid name query parameter"}
	}

	json := search.GetResults(name, limit, isNSFW, isWestern, isGame)

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the series"}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}
