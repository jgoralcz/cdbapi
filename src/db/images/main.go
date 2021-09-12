package images

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetImageByID gets the image information based off the user's input for an ID.
func GetImageByID(id int) (Image, error) {
	var image = Image{}
	err := db.Get(&image, imageByID, id)
	return image, err
}
