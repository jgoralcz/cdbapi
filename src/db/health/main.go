package health

import (
	"github.com/jgoralcz/cdbapi/src/db"
)

// GetNow gets the current timestamp from the database.
func GetNow() (Health, error) {
	var health = Health{}
	err := db.Get(&health, now)
	return health, err
}
