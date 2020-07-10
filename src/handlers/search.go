package handlers

import (
	"github.com/jgoralcz/cdbapi/src/db/search"
	"github.com/jgoralcz/cdbapi/src/lib/helpers"
	_ "github.com/jgoralcz/cdbapi/src/lib/structs" // For Swagger purposes
	"github.com/labstack/echo/v4"
)

// Search is a handler for echo that gets all of the metadata based off the name and user's filters.
// @Summary Gets a "search" based off the user's query parameters.
// @Description Search all resources by name (string), nsfw (boolean), game (boolean), western (boolean), limit (1-5), name (string). You must use name to get a result back.
// @Produce json
// @Param name query string true "name to search"
// @Param limit query int false "limit 1-5; Default 5"
// @Param nsfw query boolean false "whether the search is nsfw or not"
// @Param western query boolean false "whether the search is western (Cartoon) or not (Anime)"
// @Param game query boolean false "whether the search is from a game or not"
// @Success 200 {array} search.Search
// @Failure 400 {object} httputil.HTTPError "Must have a valid name query parameter"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when searching"
// @Router /search [get]
// @Tags Search
func Search(c echo.Context) (err error) {
	initLimit := c.QueryParam("limit")
	nsfw := c.QueryParam("nsfw")
	western := c.QueryParam("western")
	game := c.QueryParam("game")
	name := c.QueryParam("name")

	limit := helpers.MaxLimit(initLimit, 5, 5)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	if name == "" {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid name query parameter"}
	}

	json := search.GetResults(name, limit, isNSFW, isWestern, isGame)

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when searching"}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}
