package characters

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetSearchCharacter searches for a character based off the user's input for:
// name, limit (1-20), nsfw (true or false), western (true or false), game (true or false)
func GetSearchCharacter(name string, limit int, nsfw string, western string, game string) ([]Character, error) {
	var query string

	if len(name) <= 3 {
		query = characterSearchShort
	} else {
		query = characterSearch
	}

	rows := db.PoolQueryRows(query, name, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetRandomCharacter retrieves a random character from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomCharacter(limit int, nsfw string, western string, game string) ([]Character, error) {
	rows := db.PoolQueryRows(characterRandom, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetCharacterByID gets the character information based off the user's input for an ID.
func GetCharacterByID(id int) (*Character, error) {
	row := db.PoolQueryRow(characterByID, id)
	return handleRow(row)
}

// GetCharactersBySearchSeries gets character information by series name.
func GetCharactersBySearchSeries(series string, limit int, offset int) ([]Character, error) {
	rows := db.PoolQueryRows(charactersBySeries, series, limit, offset)
	return handleRows(rows)
}

// GetCharacterImages gets the basic information on the character images.
func GetCharacterImages(id int, limit int, offset int, nsfw string, crop string) ([]CharacterImage, error) {
	rows := db.PoolQueryRows(characterImagesByIDOffsetLimit, id, limit, offset, nsfw, crop)
	return handleBasicImage(rows)
}
