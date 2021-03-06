package characters

import (
	"github.com/jgoralcz/go_cdbapi/src/db"
)

// SearchCharacter searches for a character based off the user's input for:
// name, limit (1-20), nsfw (true or false), western (true or false), game (true or false)
func SearchCharacter(name string, limit int, nsfw string, western string, game string) []byte {
	rows := db.PoolQueryRows(SearchCharacterQuery, name, limit, nsfw, western, game)
	return Handler(rows)
}

// GetRandomCharacter retrieves a random character from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomCharacter(limit int, nsfw string, western string, game string) []byte {
	rows := db.PoolQueryRows(RandomCharacterQuery, limit, nsfw, western, game)
	return Handler(rows)
}

// GetCharacterByID gets the character information based off the user's input for an ID.
func GetCharacterByID(id int) []byte {
	rows := db.PoolQueryRows(CharacterByID, id)
	return Handler(rows)
}
