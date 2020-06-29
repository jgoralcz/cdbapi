package characters

// CharacterSearch is the query that searches for a character based off the user's input
// for the name, limit, nsfw, western, and game params;
var CharacterSearch = `
SELECT id, name, description, original_name, origin, image_url, image_url_clean AS image_url_crop, nsfw, "seriesNsfw",
	is_game AS game, is_western AS western, series, series_id, age, date_of_birth, hip_cm,
	bust_cm, weight_kg, height_cm, blood_type,
	(
		SELECT json_agg(item)
		FROM (
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_id
				FROM waifu_schema.appears_in wsai
				WHERE wsai.waifu_id = wt.id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
  FROM (
		SELECT ws.id, ws.name, ws.description, ws.original_name, ws.origin, ws.image_url, ws.image_url_clean, ws.nsfw,
			wsst.nsfw AS "seriesNsfw", wsst.is_game, wsst.is_western, wsst.name AS series, ws.series_id, ws.age, ws.date_of_birth,
			ws.hip_cm, ws.bust_cm, ws.weight_kg, ws.height_cm, ws.blood_type
		FROM waifu_schema.waifu_table ws
		JOIN waifu_schema.series_table wsst ON wsst.id = ws.series_id
    WHERE (
			ws.name ILIKE '%' || $1 || '%' OR levenshtein(ws.name, $1) <= 2
      OR (ws.original_name ILIKE '%' || $1 || '%' AND ws.original_name IS NOT NULL)
			OR (ws.romaji_name ILIKE '%' || $1 || '%' AND ws.romaji_name IS NOT NULL)
		)
		AND (
			('false' = $3 AND ws.nsfw = FALSE)
			OR ('true' = $3 AND ws.nsfw = TRUE)
			OR ws.nsfw IS NOT NULL
		)
		AND (
			('false' = $3 AND wsst.nsfw = FALSE)
			OR ('true' = $3 AND wsst.nsfw = TRUE)
			OR wsst.nsfw IS NOT NULL
		)
		AND (
			('false' = $4 AND wsst.is_western = FALSE)
			OR ('true' = $4 AND wsst.is_western = TRUE)
			OR wsst.is_western IS NOT NULL
		)
		AND (
			('false' = $5 AND wsst.is_game = FALSE)
			OR ('true' = $5 AND wsst.is_game = TRUE)
			OR wsst.is_game IS NOT NULL
		)
    ORDER BY
      CASE
      WHEN ws.name ILIKE $1 THEN 0
      WHEN ws.name ILIKE $1 || '%' THEN 1
      WHEN ws.name ILIKE '%' || $1 || '%' THEN 2
      WHEN ws.romaji_name ILIKE $1 THEN 3
      WHEN ws.romaji_name ILIKE $1 || '%' THEN 4
      WHEN ws.original_name ILIKE $1 THEN 5
      WHEN ws.original_name ILIKE $1 || '%' THEN 6
      WHEN levenshtein(ws.name, $1) <= 1 THEN 7
      ELSE 8 END, ws.name, ws.romaji_name, ws.original_name
    LIMIT $2
  ) wt
  ORDER BY
    CASE
    WHEN wt.name ILIKE $1 THEN 0
    WHEN wt.original_name ILIKE $1 THEN 1
    WHEN $1 ILIKE ANY (
      SELECT UNNEST(string_to_array(wt.name, ' ')) AS name
    ) THEN 2
    WHEN wt.name ILIKE $1 || '%' THEN 3
    WHEN wt.name ILIKE '%' || $1 || '%' THEN 4
    WHEN wt.original_name ILIKE $1 THEN 5
    WHEN wt.original_name ILIKE $1 || '%' THEN 6
    WHEN levenshtein(wt.name, $1) <= 1 THEN 7
    ELSE 8 END, wt.name, wt.original_name
  LIMIT $2;
`

// CharacterRandom gets a random character based off the user's input for limit, nsfw, western, and game.
var CharacterRandom = `
SELECT ws.id, ws.name, ws.description, ws.original_name, ws.origin, ws.image_url, ws.image_url_clean AS image_url_crop, 
	ws.nsfw, wsst.nsfw AS "seriesNsfw", wsst.is_game AS game, wsst.is_western AS western, wsst.name AS series,
	ws.series_id, ws.age, ws.date_of_birth, ws.hip_cm, ws.bust_cm, ws.weight_kg, ws.height_cm, ws.blood_type,
	(
		SELECT json_agg(item)
		FROM (
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
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
	OR (
		'false' != $2 AND 'true' != $2 AND ws.nsfw IS NOT NULL
	)
)
AND (
	('false' = $2 AND wsst.nsfw = FALSE)
	OR ('true' = $2 AND wsst.nsfw = TRUE)
	OR (
		'false' != $2 AND 'true' != $2 AND wsst.nsfw IS NOT NULL
	)
)
AND (
	('false' = $3 AND wsst.is_western = FALSE)
	OR ('true' = $3 AND wsst.is_western = TRUE)
	OR (
		'false' != $3 AND 'true' != $3 AND wsst.is_western IS NOT NULL
	)
)
AND (
	('false' = $4 AND wsst.is_game = FALSE)
	OR ('true' = $4 AND wsst.is_game = TRUE)
	OR (
		'false' != $4 AND 'true' != $4 AND wsst.is_game IS NOT NULL
	)
)
AND r > (
		SELECT MAX(r)
		FROM waifu_schema.waifu_table
	) * random()
ORDER BY r
LIMIT $1;
`

// CharacterByID is the query to find a character based on their ID.
var CharacterByID = `
SELECT ws.id, ws.name, ws.description, ws.original_name, ws.origin, ws.image_url, ws.image_url_clean AS image_url_crop,
	ws.nsfw, wsst.nsfw AS "seriesNsfw", wsst.is_game AS game, wsst.is_western AS western, wsst.name AS series,
	ws.series_id, ws.age, ws.date_of_birth, ws.hip_cm, ws.bust_cm, ws.weight_kg, ws.height_cm, ws.blood_type,
	(
		SELECT json_agg(item)
		FROM (
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
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
WHERE ws.id = $1;
`

// CharacterImagesByIDOffsetLimit is the query to find the images for a character with an offset and limit.
var CharacterImagesByIDOffsetLimit = `
	SELECT waifu_id, image_id, image_url_path_extra, image_url_clean_path_extra, nsfw
	FROM waifu_schema.waifu_table_images
	WHERE waifu_id = $1
		AND (
			('false' = $4 AND nsfw = FALSE)
			OR ('true' = $4 AND nsfw = TRUE)
			OR ('false' != $4 AND 'true' != $4)
		)
		AND (
			('true' = $5 AND image_url_clean_path_extra IS NOT NULL)
			OR 'true' != $5
		)
	LIMIT $2
	OFFSET $3
`
