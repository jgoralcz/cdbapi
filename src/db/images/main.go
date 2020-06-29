package images

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetImageByID gets the image information based off the user's input for an ID.
func GetImageByID(id int) string {
	row := db.PoolQueryRow(ImageByID, id)
	return handleRow(row)
}
