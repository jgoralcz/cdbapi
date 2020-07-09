package series

var seriesSearch = `
SELECT id, name, description, alternate_name, image_url_cdn, nsfw, is_game, is_western
FROM (
	SELECT id, name, description, alternate_name, image_url_cdn, nsfw, is_game, is_western
	FROM waifu_schema.series_table
	WHERE (
		name ILIKE '%' || $1 || '%' OR levenshtein(name, $1) <= 2
		OR (alternate_name ILIKE '%' || $1 || '%' AND alternate_name IS NOT NULL)
	)
	AND (
		('false' = $3 AND nsfw = FALSE)
		OR ('true' = $3 AND nsfw = TRUE)
		OR ('false' != $3 AND 'true' != $3)
	)
	AND (
		('false' = $4 AND is_western = FALSE)
		OR ('true' = $4 AND is_western = TRUE)
		OR ('false' != $4 AND 'true' != $4)
	)
	AND (
		('false' = $5 AND is_game = FALSE)
		OR ('true' = $5 AND is_game = TRUE)
		OR ('false' != $5 AND 'true' != $5)
	)
	ORDER BY
		CASE
		WHEN name ILIKE $1 THEN 0
		WHEN name ILIKE $1 || '%' THEN 1
		WHEN name ILIKE '%' || $1 || '%' THEN 2
		WHEN alternate_name ILIKE $1 THEN 3
		WHEN alternate_name ILIKE $1 || '%' THEN 4
		WHEN levenshtein(name, $1) <= 1 THEN 5
		ELSE 6 END, name, alternate_name
	LIMIT $2
) wt
ORDER BY
	CASE
	WHEN wt.name ILIKE $1 THEN 0
	WHEN $1 ILIKE ANY (
		SELECT UNNEST(string_to_array(wt.name, ' ')) AS name
	) THEN 1
	WHEN wt.name ILIKE $1 || '%' THEN 2
	WHEN wt.name ILIKE '%' || $1 || '%' THEN 3
	WHEN wt.alternate_name ILIKE $1 THEN 4
	WHEN wt.alternate_name ILIKE $1 || '%' THEN 5
	WHEN levenshtein(wt.name, $1) <= 1 THEN 6
	ELSE 7 END, wt.name, wt.alternate_name
LIMIT $2;
`

var seriesRandom = `
SELECT id, name, description, alternate_name, image_url_cdn, nsfw, is_game, is_western
FROM waifu_schema.series_table
WHERE (
	('false' = $2 AND nsfw = FALSE)
	OR ('true' = $2 AND nsfw = TRUE)
	OR (
		'false' != $2 AND 'true' != $2 AND nsfw IS NOT NULL
	)
)
AND (
	('false' = $3 AND is_western = FALSE)
	OR ('true' = $3 AND is_western = TRUE)
	OR (
		'false' != $3 AND 'true' != $3 AND is_western IS NOT NULL
	)
)
AND (
	('false' = $4 AND is_game = FALSE)
	OR ('true' = $4 AND is_game = TRUE)
	OR (
		'false' != $4 AND 'true' != $4 AND is_game IS NOT NULL
	)
)
ORDER BY random()
LIMIT $1;
`

var seriesByID = `
SELECT id, name, description, alternate_name, image_url_cdn, nsfw, is_game, is_western
FROM waifu_schema.series_table
WHERE id = $1;
`
