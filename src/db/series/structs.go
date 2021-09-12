package series

import (
	"gopkg.in/guregu/null.v3"
)

// Series is a data structure that models a series. Be sure to check for nulls.
type Series struct {
	ID          int         `json:"id" example:"4731"`
	Name        string      `json:"name" example:"Flight Rising"`
	Description null.String `json:"description" swaggertype:"string" example:"Flight Rising is a free Virtual Pet/breeding sim focusing on dragons. It launched on June 9th, 2013, following a wildly successful Kickstarter campaign."`
	ImageURL    null.String `db:"image_url" json:"image_url" swaggertype:"string" example:"https://cdn.bongo.best/series/4731/e1dcff95-fa7c-4995-92f2-b9dc9840ffe1.png"`
	ReleaseDate null.String `db:"release_date" json:"release_date" swaggertype:"string" example:"null"`
	Nsfw        bool        `json:"nsfw" example:"false"`
	Game        bool        `db:"is_game" json:"game" example:"true"`
	Western     bool        `db:"is_western" json:"western" example:"true"`
	Nicknames   []string    `json:"nicknames"`
}
