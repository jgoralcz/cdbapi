package series

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetSearchSeries searches for a series based off the user's input for:
// name, limit (1-20), nsfw (true or false), western (true or false), game (true or false)
func GetSearchSeries(name string, limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(seriesSearch, name, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetRandomSeries retrieves a random series from the database based off the user's input of:
// limit (1-20), nsfw (true or false), western (true or false), game (true or false)
// and returns all matching criteria.
func GetRandomSeries(limit int, nsfw string, western string, game string) string {
	rows := db.PoolQueryRows(seriesRandom, limit, nsfw, western, game)
	return handleRows(rows)
}

// GetSeriesByID gets the series information based off the user's input for an ID.
func GetSeriesByID(id int) string {
	row := db.PoolQueryRow(seriesByID, id)
	return handleRow(row)
}