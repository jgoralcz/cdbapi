package characters

import (
	"database/sql"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Character is a data structure that models a character.
type Character struct {
	ID               int         `json:"id" example:"7960"`
	Name             string      `json:"name" example:"Marie"`
	Description      null.String `json:"description" swaggertype:"string" example:"Marie is an assistant of the Velvet Room in Persona 4 Golden. When Marie is first met, she seems to be a very cold and antisocial individual. She is sullen, cranky, sarcastic, irritable, very foul-mouthed and often prone to mood swings. She will not hesitate to voice out her opinion or express her thoughts, regardless of how it would make everyone feel. Marie sometimes expresses her thoughts in poems which often deal with depressing themes like farewells and existentialism, questioning her own origin. Besides this, Marie is into fashion, loves nature (another major theme in her poems) and is apparently concerned about her figure, which is why she only eats healthy food."`
	ImageURL         null.String `json:"image_url" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	ImageURLCrop     null.String `json:"image_url_crop" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	Nsfw             null.Bool   `json:"nsfw" swaggertype:"boolean" example:"false"`
	Game             null.Bool   `json:"game" swaggertype:"boolean" example:"false"`
	Western          null.Bool   `json:"western" swaggertype:"boolean" example:"false"`
	Series           string      `json:"series" swaggertype:"string" example:"Persona 4"`
	SeriesID         int         `json:"series_id" example:"2240"`
	Husbando         null.Bool   `json:"husbando" swaggertype:"boolean" nullable:"true" example:"false"`
	Count            null.Int    `json:"claims" swaggertype:"integer" example:"5"`
	Position         null.Int    `json:"rank" swaggertype:"integer" example:"5"`
	LastEditBy       null.String `json:"last_edit_by" swaggertype:"string" example:"304478893010583552"`
	LastEditDate     null.String `json:"last_edit_date" swaggertype:"string"`
	Nicknames        []string    `json:"nicknames"`
	SpoilerNicknames []string    `json:"spoiler_nicknames"`
	AppearsIn        []AppearsIn `json:"appears_in"`
}

// AppearsIn is a data structure that models which series a particular character
// belongs to. It also includes convenient metadata such the series name,
// whether it is nsfw, a game, and/or a western.
type AppearsIn struct {
	Series   string `json:"series" example:"Persona 4"`
	SeriesID int    `json:"series_id" example:"2240"`
	Nsfw     bool   `json:"nsfw" swaggertype:"boolean" example:"false"`
	Game     bool   `json:"game" swaggertype:"boolean" example:"true"`
	Western  bool   `json:"western" swaggertype:"boolean" example:"false"`
}

// CharacterImage is a minimal set of data for the additional images on a character
type CharacterImage struct {
	CharacterID  int         `json:"character_id" example:"7960"`
	ImageID      int         `json:"image_id" example:"1234"`
	ImageURL     null.String `json:"image_url" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	ImageURLCrop null.String `json:"image_url_crop" swaggertype:"string" example:"https://cdn.bongo.best/characters/7960/82736d1f-fa95-4f6e-ae78-f9422f065202_thumb.png"`
	Nsfw         null.Bool   `json:"nsfw" swaggertype:"boolean" example:"false"`
}

func handleRows(rows pgx.Rows) string {
	rowsErr := rows.Err()
	if rows == nil && rowsErr == nil {
		return "[]"
	}

	characters := []Character{}
	for rows.Next() {
		c := new(Character)
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.ImageURL,
			&c.ImageURLCrop,
			&c.Nsfw,
			&c.Game,
			&c.Western,
			&c.Series,
			&c.SeriesID,
			&c.Husbando,
			&c.Count,
			&c.Position,
			&c.LastEditBy,
			&c.LastEditDate,
			&c.Nicknames,
			&c.SpoilerNicknames,
			&c.AppearsIn,
		)

		if err != nil {
			log.Error(err)
			break
		}

		if c.AppearsIn == nil {
			c.AppearsIn = []AppearsIn{}
		}

		characters = append(characters, *c)
	}

	charactersJSON, marshalErr := json.Marshal(characters)
	if marshalErr != nil {
		log.Error(marshalErr)
		return ""
	}

	if rowsErr != nil {
		log.Error(rowsErr)
		return ""
	}

	return string(charactersJSON)
}

func handleRow(row pgx.Row) string {
	c := new(Character)
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.Description,
		&c.ImageURL,
		&c.ImageURLCrop,
		&c.Nsfw,
		&c.Game,
		&c.Western,
		&c.Series,
		&c.SeriesID,
		&c.Husbando,
		&c.Count,
		&c.Position,
		&c.LastEditBy,
		&c.LastEditDate,
		&c.Nicknames,
		&c.SpoilerNicknames,
		&c.AppearsIn,
	)

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		return "{}"
	}

	if err != nil {
		log.Error(err)
		return ""
	}

	if c.AppearsIn == nil {
		c.AppearsIn = []AppearsIn{}
	}

	characterJSON, marshalErr := json.Marshal(c)
	if marshalErr != nil {
		log.Error(marshalErr)
		return ""
	}

	return string(characterJSON)
}

func handleBasicImage(rows pgx.Rows) string {
	if rows == nil {
		return "[]"
	}

	images := []CharacterImage{}
	for rows.Next() {
		i := new(CharacterImage)
		err := rows.Scan(&i.CharacterID, &i.ImageID, &i.ImageURL, &i.ImageURLCrop, &i.Nsfw)
		if err != nil {
			log.Error(err)
			break
		}

		images = append(images, *i)
	}

	imageJSON, marshalErr := json.Marshal(images)
	if marshalErr != nil {
		log.Error(marshalErr)
		return ""
	}

	rowsErr := rows.Err()
	if rowsErr != nil {
		log.Error(rowsErr)
		return ""
	}

	return string(imageJSON)
}
