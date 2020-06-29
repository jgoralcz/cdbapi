package characters

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetSearchCharacter searches for a character based off the user's input for:
// name, limit (1-20), nsfw (true or false), western (true or false), game (true or false)
func GetSearchCharacter(name string, limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(CharacterSearch, name, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetRandomCharacter retrieves a random character from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomCharacter(limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(CharacterRandom, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetCharacterByID gets the character information based off the user's input for an ID.
func GetCharacterByID(id int) string {
	row := db.PoolQueryRow(CharacterByID, id)
	return handleRow(row)
}

// GetCharacterImages gets the basic information on the character images.
func GetCharacterImages(id int, limit int, offset int, nsfw string) string {
	rows := db.PoolQueryRows(CharacterImagesByIDOffsetLimit, id, limit, offset, nsfw)
	return handleBasicImage(rows)
}
