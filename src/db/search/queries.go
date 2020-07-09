package search

var searchAll = `
SELECT (
	SELECT json_agg(character)
	FROM (
		SELECT id, name, original_name, image_url_clean, nsfw, "seriesNsfw", is_game, is_western, series, series_id
		FROM (
			SELECT wswt.id, wswt.name, wswt.original_name, image_url_clean, wswt.nsfw, wsst.nsfw AS "seriesNsfw", wsst.is_game, wsst.is_western, wswt.series, wswt.series_id
			FROM waifu_schema.waifu_table wswt
			JOIN waifu_schema.series_table wsst ON wsst.id = wswt.series_id
			WHERE (
				wswt.name ILIKE '%' || $1 || '%' OR levenshtein(wswt.name, $1) <= 2
				OR (wswt.original_name ILIKE '%' || $1 || '%' AND wswt.original_name IS NOT NULL)
				OR (wswt.romaji_name ILIKE '%' || $1 || '%' AND wswt.romaji_name IS NOT NULL)
			)
			AND (
				('false' = $3 AND wswt.nsfw = FALSE)
				OR ('true' = $3 AND wswt.nsfw = TRUE)
				OR wswt.nsfw IS NOT NULL
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
				WHEN wswt.name ILIKE $1 THEN 0
				WHEN wswt.name ILIKE $1 || '%' THEN 1
				WHEN wswt.name ILIKE '%' || $1 || '%' THEN 2
				WHEN wswt.romaji_name ILIKE $1 THEN 3
				WHEN wswt.romaji_name ILIKE $1 || '%' THEN 4
				WHEN wswt.original_name ILIKE $1 THEN 5
				WHEN wswt.original_name ILIKE $1 || '%' THEN 6
				WHEN levenshtein(wswt.name, $1) <= 1 THEN 7
				ELSE 8 END, wswt.name, wswt.romaji_name, wswt.original_name
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
		LIMIT $2
	) character
),
(
	SELECT json_agg(series)
	FROM (
		SELECT id, name, alternate_name, image_url_cdn, nsfw, is_game, is_western
		FROM (
			SELECT id, name, alternate_name, image_url_cdn, nsfw, is_game, is_western
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
		LIMIT $2
	) series
)
`
