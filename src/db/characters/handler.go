package characters

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Character is a data structure that models a character. Be sure to check for nulls.
type Character struct {
	ID            int         `json:"id" example:"7960"`
	Name          string      `json:"name" example:"Marie"`
	Description   null.String `json:"description" example:"Marie is an assistant of the Velvet Room in Persona 4 Golden. When Marie is first met, she seems to be a very cold and antisocial individual. She is sullen, cranky, sarcastic, irritable, very foul-mouthed and often prone to mood swings. She will not hesitate to voice out her opinion or express her thoughts, regardless of how it would make everyone feel. Marie sometimes expresses her thoughts in poems which often deal with depressing themes like farewells and existentialism, questioning her own origin. Besides this, Marie is into fashion, loves nature (another major theme in her poems) and is apparently concerned about her figure, which is why she only eats healthy food."`
	OriginalName  null.String `json:"original_name" example:"null"`
	Origin        null.String `json:"origin" example:"VelvetRoom"`
	ImageURL      null.String `json:"image_url" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	ImageURLClean null.String `json:"image_url_clean" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	Nsfw          bool        `json:"nsfw" example:"false"`
	SeriesNsfw    bool        `json:"series_nsfw" example:"false"`
	Game          bool        `json:"game" example:"false"`
	Western       bool        `json:"western" format:"boolean" example:"false"`
	Series        null.String `json:"series" example:"Persona 4"`
	SeriesID      int         `json:"series_id" example:"2240"`
	Age           null.Int    `json:"age" format:"integer" example:"0"`
	DateOfBirth   null.String `json:"date_of_birth" example:"null"`
	HipCM         null.Float  `json:"hip_cm" format:"number" example:"0"`
	BustCM        null.Float  `json:"bust_cm" format:"number" example:"0"`
	WeightKG      null.Float  `json:"weight_kg" format:"number" example:"0"`
	HeightCM      null.Float  `json:"height_cm" format:"number" example:"164"`
	BloodType     null.String `json:"blood_type" example:"null"`
	AppearsIn     []AppearsIn `json:"appears_in"`
}

// AppearsIn is a data structure that models which series a particular character
// belongs to. It also includes convenient metadata such the series name,
// whether it is nsfw, a game, and/or a western.
type AppearsIn struct {
	Series   string `json:"series" example:"Persona 4"`
	SeriesID int    `json:"series_id" example:"2240"`
	Nsfw     bool   `json:"nsfw" format:"boolean" example:"false"`
	Game     bool   `json:"game" format:"boolean" example:"true"`
	Western  bool   `json:"western" format:"boolean" example:"false"`
}

// HandleRows handles rows from a database query and populates them into a Character struct.
// It then parses them into a stringified json to send back.
func HandleRows(rows pgx.Rows) string {
	if rows == nil {
		return "[]"
	}

	characters := []Character{}
	for rows.Next() {
		c := new(Character)
		err := rows.Scan(
			&c.ID, &c.Name, &c.Description, &c.OriginalName, &c.Origin,
			&c.ImageURL, &c.ImageURLClean, &c.Nsfw, &c.SeriesNsfw, &c.Game, &c.Western, &c.Series,
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

// HandleRow handles a row from the database and parses it into a character struct.
// It then parses them into a stringified json to send back.
func HandleRow(row pgx.Row) string {
	c := new(Character)

	err := row.Scan(
		&c.ID, &c.Name, &c.Description, &c.OriginalName, &c.Origin,
		&c.ImageURL, &c.ImageURLClean, &c.Nsfw, &c.SeriesNsfw, &c.Game, &c.Western, &c.Series,
		&c.SeriesID, &c.Age, &c.DateOfBirth, &c.HipCM, &c.BustCM, &c.WeightKG,
		&c.HeightCM, &c.BloodType, &c.AppearsIn,
	)

	if err != nil && err != sql.ErrNoRows {
		return "[]"
	}

	if err != nil {
		return ""
	}

	if c.AppearsIn == nil {
		c.AppearsIn = []AppearsIn{}
	}

	characterJSON, marshalErr := json.Marshal(c)
	if marshalErr != nil {
		log.Println(marshalErr)
		return ""
	}

	return string(characterJSON)
}
