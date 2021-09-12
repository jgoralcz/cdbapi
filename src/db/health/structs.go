package health

import "time"

// Health is a data structure that models the health of the database with a now() check
type Health struct {
	Now time.Time `db:"now" json:"now" example:"2021-09-11T03:32:27.409955Z"`
}
