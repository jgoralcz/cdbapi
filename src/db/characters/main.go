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

	var characters = []Character{}
	err := db.Select(&characters, query, name, limit, nsfw, western, game)
	return characters, err
}

// GetCharactersBySearchSeries gets character information by series name.
func GetCharactersBySearchSeries(series string, limit int, offset int) ([]Character, error) {
	var characters = []Character{}
	err := db.Select(&characters, charactersBySeries, series, limit, offset)
	return characters, err
}

// GetRandomCharacter retrieves a random character from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomCharacter(limit int, nsfw string, western string, game string) ([]Character, error) {
	var characters = []Character{}
	err := db.Select(&characters, characterRandom, limit, nsfw, western, game)
	return characters, err
}

// GetCharacterByID gets the character information based off the user's input for an ID.
func GetCharacterByID(id int) (Character, error) {
	var character = Character{}
	err := db.Get(&character, characterByID, id)
	return character, err
}

// GetCharacterImages gets the basic information on the character images.
func GetCharacterImages(id int, limit int, offset int, nsfw string, crop string) ([]CharacterImage, error) {
	var characterImage = []CharacterImage{}
	err := db.Select(&characterImage, characterImagesByIDOffsetLimit, id, limit, offset, nsfw, crop)
	return characterImage, err
}
