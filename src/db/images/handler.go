package images

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Image is a data structure that models a character. Be sure to check for nulls.
type Image struct {
	ID             int         `json:"id" example:"187840"`
	CharacterID    int         `json:"character_id" example:"10739"`
	ImageURL       null.String `json:"image_url" swaggertype:"string" example:" https://cdn.bongo.best/characters/10739/6ae0800a-e3d8-43e2-b783-9eeb595fa30d.jpg"`
	ImageURLCrop   null.String `json:"image_url_crop" swaggertype:"string" example:"https://cdn.bongo.best/characters/187840/ab9b6e58-a055-4702-879a-7f3300c771f5.jpg"`
	Nsfw           null.Bool   `json:"nsfw" swaggertype:"boolean" example:"false"`
	Width          null.Float  `json:"width" swaggertype:"number" example:"2420"`
	Height         null.Float  `json:"height" swaggertype:"number" example:"2670"`
	FileType       null.String `json:"file_type" swaggertype:"string" example:"png"`
	BufferSize     null.Int    `json:"buffer_size" swaggertype:"number" example:"402123"`
	BufferSizeCrop null.Int    `json:"buffer_size_crop" swaggertype:"number" example:"72288"`
}

func handleRow(row pgx.Row) string {
	i := new(Image)

	err := row.Scan(
		&i.ID, &i.CharacterID, &i.ImageURL, &i.ImageURLCrop, &i.Nsfw,
		&i.Width, &i.Height, &i.FileType,
		&i.BufferSize, &i.BufferSizeCrop,
	)

	if err != nil && err != sql.ErrNoRows {
		return "{}"
	}

	if err != nil {
		return ""
	}

	imagesJSON, marshalErr := json.Marshal(i)
	if marshalErr != nil {
		log.Println(marshalErr)
		return ""
	}

	return string(imagesJSON)
}
