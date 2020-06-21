package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jgoralcz/go_cdbapi/src/db/characters"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
)

// CharacterByIDHandler handles routes when a user searches by a character ID.
// Returns 400 if the id is not an int, and returns 404 if the id is not found.
func CharacterByIDHandler(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Must have a valid id parameter"})
		return
	}

	json := characters.GetCharacterByID(id)

	if json == "[]" {
		c.JSON(404, gin.H{"error": "Could not find a character with id " + strID})
		return
	}

	if json == "" {
		c.JSON(500, gin.H{"error": "An unexpected error has occurred when retrieving the character with id " + strID})
		return
	}

	// TODO: make this into a function?
	c.Header("Content-Type", "application/json")
	c.String(200, json)
}

// CharacterNameHandler handles the logic for searching for a character with the
// user's specified parameters that act as filters: limit (int), nsfw (true or false),
// western (true or false), game (true or false)
func CharacterNameHandler(c *gin.Context) {
	initLimit := c.Query("limit")
	nsfw := c.Query("nsfw")
	western := c.Query("western")
	game := c.Query("game")
	name := c.Query("name")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	if name == "" {
		c.JSON(400, gin.H{"error": "Must have a valid name query parameter"})
	}

	json := characters.SearchCharacter(name, limit, isNSFW, isWestern, isGame)

	if json == "" {
		c.JSON(500, gin.H{"error": "An unexpected error has occurred when retrieving the character with name " + name})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(200, json)
}

// CharacterRandomHandler handles the logic for a random character request with
// the user's specified parameters that act as filters: limit (int), nsfw (true or false),
// western (true or false), game (true or false)
func CharacterRandomHandler(c *gin.Context) {
	initLimit := c.Query("limit")
	nsfw := c.Query("nsfw")
	western := c.Query("western")
	game := c.Query("game")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	json := characters.GetRandomCharacter(limit, isNSFW, isWestern, isGame)

	if json == "" {
		c.JSON(500, gin.H{"error": "An unexpected error has occurred when retrieving the random character"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(200, json)
}
