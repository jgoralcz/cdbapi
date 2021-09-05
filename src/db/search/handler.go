package search

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Search is the combination of all the possible searches to reduce server load.
type Search struct {
	Characters []character `json:"characters"`
	Series     []series    `json:"series"`
}

type character struct {
	ID           int         `json:"id" example:"7960"`
	Name         string      `json:"name" example:"Marie"`
	OriginalName null.String `json:"original_name" swaggertype:"string" example:"null"`
	ImageURL     null.String `json:"image_url" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	ImageURLCrop null.String `json:"image_url_crop" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	Nsfw         bool        `json:"nsfw" example:"false"`
	SeriesNsfw   bool        `json:"series_nsfw" example:"false"`
	Game         bool        `json:"game" example:"false"`
	Western      bool        `json:"western" example:"false"`
	Series       null.String `json:"series" swaggertype:"string" example:"Persona 4"`
	SeriesID     int         `json:"series_id" example:"2240"`
}

type series struct {
	ID      int    `json:"id" example:"4731"`
	Name    string `json:"name" example:"Flight Rising"`
	Nsfw    bool   `json:"nsfw" example:"false"`
	Game    bool   `json:"game" example:"true"`
	Western bool   `json:"western" example:"true"`
}

func handleRow(row pgx.Row) string {
	s := new(Search)

	err := row.Scan(&s.Characters, &s.Series)

	if err != nil && err != sql.ErrNoRows {
		return "{}"
	}

	if err != nil {
		return ""
	}

	searchJSON, marshalErr := json.Marshal(s)
	if marshalErr != nil {
		log.Println(marshalErr)
		return ""
	}

	return string(searchJSON)
}
