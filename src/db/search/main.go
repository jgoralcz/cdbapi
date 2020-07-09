package search

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetResults gets the necessary results to perform a search.
func GetResults(name string, limit int, nsfw string, western string, game string) string {
	row := db.PoolQueryRow(searchAll, name, limit, nsfw, western, game)
	return handleRow(row)
}
