package images

import (
	"gopkg.in/guregu/null.v3"
)

// Image is a data structure that models a character. Be sure to check for nulls.
type Image struct {
	ID             int         `db:"image_id" json:"id" example:"187840"`
	CharacterID    int         `db:"character_id" json:"character_id" example:"10739"`
	ImageURL       null.String `db:"image_url_path_extra" json:"image_url" swaggertype:"string" example:" https://cdn.bongo.best/characters/10739/6ae0800a-e3d8-43e2-b783-9eeb595fa30d.jpg"`
	ImageURLCrop   null.String `db:"image_url_clean_path_extra" json:"image_url_crop" swaggertype:"string" example:"https://cdn.bongo.best/characters/187840/ab9b6e58-a055-4702-879a-7f3300c771f5.jpg"`
	Nsfw           null.Bool   `json:"nsfw" swaggertype:"boolean" example:"false"`
	Width          null.Float  `json:"width" swaggertype:"number" example:"2420"`
	Height         null.Float  `json:"height" swaggertype:"number" example:"2670"`
	FileType       null.String `db:"file_type" json:"file_type" swaggertype:"string" example:"png"`
	BufferSize     null.Int    `db:"buffer_length" json:"buffer_size" swaggertype:"number" example:"402123"`
	BufferSizeCrop null.Int    `db:"buffer_length_clean" json:"buffer_size_crop" swaggertype:"number" example:"72288"`
}
