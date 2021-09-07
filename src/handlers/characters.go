package handlers

import (
	"strconv"

	"github.com/jgoralcz/cdbapi/src/db/characters"
	"github.com/jgoralcz/cdbapi/src/lib/helpers"
	_ "github.com/jgoralcz/cdbapi/src/lib/structs" // For Swagger purposes
	"github.com/labstack/echo/v4"
)

// CharacterByID is a handler for echo that gets the character metadata by id.
// @Summary Gets a character by the ID
// @Description get character metadata by id
// @Produce json
// @Param id path int true "Some ID"
// @Success 200 {object} characters.Character
// @Failure 400 {object} httputil.HTTPError "Must have a valid id parameter"
// @Failure 404 {object} httputil.HTTPError "Could not find a character with id Some ID"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the character with id Some ID"
// @Router /characters/{id} [get]\
// @Tags Character
func CharacterByID(c echo.Context) (err error) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)

	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid id parameter"}
	}

	json := characters.GetCharacterByID(id)

	if json == "{}" {
		return &echo.HTTPError{Code: 404, Message: "Could not find a character with id " + strID}
	}

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the character with id " + strID}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}

// CharacterImages is a handler for echo that gets a character's images based off the character ID, offset, and limit.
// @Summary Gets a character's image based off the character's ID, offset, and limit
// @Description Get the character's images between the requested offset and limit with the requested ID.
// @Produce json
// @Param id path int true "Some ID"
// @Param limit query int false "limit 1-10; Default 10"
// @Param offset query int false "the offset for the images for pagination"
// @Param nsfw query boolean false "whether the image is nsfw or not"
// @Param crop query boolean false "filter down by cropped images only"
// @Success 200 {array} characters.CharacterImage
// @Failure 400 {object} httputil.HTTPError "Must have a valid id parameter"
// @Failure 400 {object} httputil.HTTPError "Invalid limit provided. Limit must be a valid number less than 10 and greater than 0"
// @Failure 400 {object} httputil.HTTPError "Invalid offset provided. Offset must be a valid number and greater than 0"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the images"
// @Router /characters/{id}/images [get]
// @Tags Character
func CharacterImages(c echo.Context) (err error) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)

	nsfw := c.QueryParam("nsfw")
	isNSFW := helpers.DefaultBoolean(nsfw)

	crop := c.QueryParam("crop")
	isCrop := helpers.DefaultBoolean(crop)

	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid id parameter"}
	}

	strLimit := c.QueryParam("limit")
	limit := helpers.DefaultNumber(strLimit, 10)

	if limit == -1 || limit > 10 || limit < 1 {
		return &echo.HTTPError{Code: 400, Message: "Invalid limit provided. Limit must be a valid number less than 10 and greater than 0"}
	}

	strOffset := c.QueryParam("offset")
	offset := helpers.DefaultNumber(strOffset, 0)

	if offset == -1 || offset < 0 {
		return &echo.HTTPError{Code: 400, Message: "Invalid offset provided. Offset must be a valid number and greater than 0"}
	}

	json := characters.GetCharacterImages(id, limit, offset, isNSFW, isCrop)

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the images for the character"}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}

// CharacterRandom is a handler for echo that gets a random character based off a user's filters.
// @Summary Gets a random character based off the user's query parameters.
// @Description Get a random character metadata by nsfw (boolean), game (boolean), western (boolean), limit (1-20).
// @Produce json
// @Param limit query int false "limit 1-20; Default 1"
// @Param nsfw query boolean false "whether the character is nsfw or not"
// @Param western query boolean false "whether the character is western (Cartoon) or not (Anime)"
// @Param game query boolean false "whether the character is from a game or not"
// @Success 200 {array} characters.Character
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the character"
// @Router /characters/random [get]
// @Tags Character
func CharacterRandom(c echo.Context) (err error) {
	initLimit := c.QueryParam("limit")
	nsfw := c.QueryParam("nsfw")
	western := c.QueryParam("western")
	game := c.QueryParam("game")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	json := characters.GetRandomCharacter(limit, isNSFW, isWestern, isGame)

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the character"}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}

// TODO: how to do name or series required
// Character is a handler for echo that gets the character metadata based off the name or series and user's filters.
// @Summary Gets a character based off the user's query parameters.
// @Description Get character metadata by nsfw (boolean), game (boolean), western (boolean), limit (1-20), name (string). You must use name to get a result back.
// @Produce json
// @Param name query string true "name to search"
// @Param series query string true "series to search by"
// @Param limit query int false "limit 1-20; Default 1"
// @Param nsfw query boolean false "whether the character is nsfw or not"
// @Param western query boolean false "whether the character is western (Cartoon) or not (Anime)"
// @Param game query boolean false "whether the character is from a game or not"
// @Success 200 {array} characters.Character
// @Failure 400 {object} httputil.HTTPError "Must have a valid name or series query parameter"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the character"
// @Router /characters [get]
// @Tags Character
func Character(c echo.Context) (err error) {
	initLimit := c.QueryParam("limit")
	initOffset := c.QueryParam("offset")
	nsfw := c.QueryParam("nsfw")
	western := c.QueryParam("western")
	game := c.QueryParam("game")
	name := c.QueryParam("name")
	series := c.QueryParam("series")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)
	offset := helpers.DefaultNumber(initOffset, 0)

	if name == "" && series == "" {
		return &echo.HTTPError{Code: 400, Message: "Must have a valid name or series query parameter"}
	}

	var json string
	if name != "" {
		json = characters.GetSearchCharacter(name, limit, isNSFW, isWestern, isGame)
	} else {
		json = characters.GetCharactersBySearchSeries(series, limit, offset)
	}

	if json == "" {
		return &echo.HTTPError{Code: 500, Message: "An unexpected error has occurred when retrieving the character"}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(200, json)
}
