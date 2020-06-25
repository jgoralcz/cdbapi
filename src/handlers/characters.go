package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jgoralcz/go_cdbapi/src/db/characters"
	"github.com/jgoralcz/go_cdbapi/src/lib/helpers"
	"github.com/jgoralcz/go_cdbapi/src/lib/httputil"
)

// @Summary Gets a character by the ID
// @Description get character metadata by id
// @Produce json
// @Param id path int true "Some ID"
// @Success 200 {object} characters.Character
// @Failure 400 {object} httputil.HTTPError "Must have a valid id parameter"
// @Failure 404 {object} httputil.HTTPError "Could not find a character with id Some ID"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the character with id Some ID"
// @Router /characters/{id} [get]
func CharacterByID(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Must have a valid id parameter"})
		return
	}

	json := characters.GetCharacterByID(id)

	if json == "[]" {
		httputil.NewError(c, 404, "Could not find a character with id "+strID)
		return
	}

	if json == "" {
		httputil.NewError(c, 500, "An unexpected error has occurred when retrieving the character with id "+strID)
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(200, json)
}

// @Summary Gets a character based off the user's query parameters.
// @Description Get character metadata by nsfw (boolean), game (boolean), western (boolean), limit (1-20), name (string), random (boolean). You must use either name or random in order to get a result back.
// @Produce json
// @Param name query string true "name to search (use without random param)"
// @Param random query boolean true "true or false (use without name param)"
// @Param limit query int false "limit 1-20; Default 1"
// @Param nsfw query boolean false "whether the character is nsfw or not"
// @Param western query boolean false "whether the character is western (Cartoon) or not (Anime)"
// @Param game query boolean false "whether the character is from a game or not"
// @Success 200 {array} characters.Character
// @Failure 400 {object} httputil.HTTPError "Must have a valid name query parameter or specify that the query parameter random is true"
// @Failure 500 {object} httputil.HTTPError "An unexpected error has occurred when retrieving the character"
// @Router /characters [get]
func Character(c *gin.Context) {
	initLimit := c.Query("limit")
	nsfw := c.Query("nsfw")
	western := c.Query("western")
	game := c.Query("game")
	name := c.Query("name")
	random := c.Query("random")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)
	isRandom := helpers.DefaultBoolean(random)

	if name == "" && isRandom != "true" {
		httputil.NewError(c, 400, "Must have a valid name query parameter or specify that the query parameter random is true")
		return
	}

	var json string
	if name != "" {
		json = characters.SearchCharacter(name, limit, isNSFW, isWestern, isGame)
	} else {
		json = characters.GetRandomCharacter(limit, isNSFW, isWestern, isGame)
	}

	if json == "" {
		httputil.NewError(c, 500, "An unexpected error has occurred when retrieving the character")
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(200, json)
}
