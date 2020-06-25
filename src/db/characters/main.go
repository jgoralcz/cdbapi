package characters

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// SearchCharacter searches for a character based off the user's input for:
// name, limit (1-20), nsfw (true or false), western (true or false), game (true or false)
func SearchCharacter(name string, limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(SearchCharacterQuery, name, limit, nsfw, western, game)
	return HandleRows(rows)
}

// GetRandomCharacter retrieves a random character from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomCharacter(limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(RandomCharacterQuery, limit, nsfw, western, game)
	return HandleRows(rows)
}

// GetCharacterByID gets the character information based off the user's input for an ID.
func GetCharacterByID(id int) string {
	rows := db.PoolQueryRow(CharacterByID, id)
	return HandleRow(rows)
}
