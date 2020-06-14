package db

import (
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
	"gopkg.in/guregu/null.v4"
)

// TODO: select isGame, isWestern, check on isNSFW
// TODO: fix appearsIn
// TODO: insert back in URL when we have our own website
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
	HipCM        null.Int    `json:"hip_cm"`
	BustCM       null.Int    `json:"bust_cm"`
	WeightKG     null.Int    `json:"weight_kg"`
	HeightCM     null.Int    `json:"height_cm"`
	BloodType    null.String `json:"blood_type"`
	AppearsIn    null.String `json:"appears_in"`
}

var defaultByteArray = []byte{'[', ']'}

func GetRandomCharacter(limit int, nsfw string, western string, game string) []byte {
	rows := buildRandomCharacterQuery(limit, nsfw, western, game)

	if rows == nil {
		return defaultByteArray
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

		// TODO: fix null string
		// if c.AppearsIn == null.String {
		// 	c.AppearsIn = []
		// }

		characters = append(characters, *c)
	}

	charactersJSON, marshalErr := json.Marshal(characters)
	if marshalErr != nil {
		log.Println(marshalErr)
		return defaultByteArray
	}

	rowsErr := rows.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
	}

	return charactersJSON
}

var randomCharacterQuery = `
SELECT ws.id, ws.name, ws.description, ws.original_name, ws.origin, ws.image_url, ws.nsfw, wsst.nsfw AS "seriesNsfw",
	wsst.is_game AS game, wsst.is_western AS western, wsst.name AS series, ws.series_id, ws.age, ws.date_of_birth, ws.hip_cm, ws.bust_cm, ws.weight_kg, ws.height_cm, ws.blood_type,
(
	SELECT json_agg(item)
	FROM (
		SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw
		FROM (
			SELECT series_id
			FROM waifu_schema.appears_in wsai
			WHERE wsai.waifu_id = ws.id
		) wsai
		JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
	) item
) AS appears_in
FROM waifu_schema.waifu_table ws
JOIN waifu_schema.series_table wsst ON wsst.id = ws.series_id
WHERE (
	('false' = $2 AND ws.nsfw = FALSE)
	OR ('true' = $2 AND ws.nsfw = TRUE)
	OR ws.nsfw IS NOT NULL
)
AND (
	('false' = $2 AND wsst.nsfw = FALSE)
	OR ('true' = $2 AND wsst.nsfw = TRUE)
	OR wsst.nsfw IS NOT NULL
)
AND (
	('false' = $3 AND wsst.is_western = FALSE)
	OR ('true' = $3 AND wsst.is_western = TRUE)
	OR wsst.is_western IS NOT NULL
)
AND (
	('false' = $4 AND wsst.is_game = FALSE)
	OR ('true' = $4 AND wsst.is_game = TRUE)
	OR wsst.is_game IS NOT NULL
)
AND r > (
		SELECT MAX(r)
		FROM waifu_schema.waifu_table
	) * random()
ORDER BY r
LIMIT $1;
`

func buildRandomCharacterQuery(limit int, nsfw string, western string, game string) pgx.Rows {
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	return PoolQueryRows(randomCharacterQuery, limit, isNSFW, isWestern, isGame)
}
