package characters

import (
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Character is a data structure that models a character. Be sure to check for nulls.
type Character struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	OriginalName null.String `json:"original_name"`
	Origin       null.String `json:"origin"`
	ImageURL     null.String `json:"image_url"`
	Nsfw         null.Bool   `json:"nsfw"`
	SeriesNsfw   null.Bool   `json:"series_nsfw"`
	Game         null.Bool   `json:"game"`
	Western      null.Bool   `json:"western"`
	Series       null.String `json:"series"`
	SeriesID     null.Int    `json:"series_id"`
	Age          null.Int    `json:"age"`
	DateOfBirth  null.String `json:"date_of_birth"`
	HipCM        null.Float  `json:"hip_cm"`
	BustCM       null.Float  `json:"bust_cm"`
	WeightKG     null.Float  `json:"weight_kg"`
	HeightCM     null.Float  `json:"height_cm"`
	BloodType    null.String `json:"blood_type"`
	AppearsIn    []AppearsIn `json:"appears_in"`
}

// AppearsIn is a data structure that models which series a particular character
// belongs to. It also includes convenient metadata such the series name,
// whether it is nsfw, a game, and/or a western.
type AppearsIn struct {
	Series   string    `json:"series"`
	SeriesID int       `json:"series_id"`
	Nsfw     null.Bool `json:"nsfw"`
	Game     null.Bool `json:"game"`
	Western  null.Bool `json:"western"`
}

// Handler handles rows from a database query and populates them into a Character struct.
// It then parses them into json to send back.
func Handler(rows pgx.Rows) string {
	if rows == nil {
		return "[]"
	}

	characters := []Character{}
	for rows.Next() {
		c := new(Character)
		err := rows.Scan(
			&c.ID, &c.Name, &c.Description, &c.OriginalName, &c.Origin,
			&c.ImageURL, &c.Nsfw, &c.SeriesNsfw, &c.Game, &c.Western, &c.Series,
			&c.SeriesID, &c.Age, &c.DateOfBirth, &c.HipCM, &c.BustCM, &c.WeightKG,
			&c.HeightCM, &c.BloodType, &c.AppearsIn,
		)
		if err != nil {
			log.Println(err)
			break
		}

		if c.AppearsIn == nil {
			c.AppearsIn = []AppearsIn{}
		}

		characters = append(characters, *c)
	}

	charactersJSON, marshalErr := json.Marshal(characters)
	if marshalErr != nil {
		log.Println(marshalErr)
		return ""
	}

	rowsErr := rows.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return ""
	}

	return string(charactersJSON)
}
