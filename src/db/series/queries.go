package series

var seriesSearch = `
SELECT id, name, description, image_url, release_date, nsfw, is_game, is_western, nicknames
  FROM (
    SELECT wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western,
    COALESCE(array_remove(array_agg(DISTINCT(wssn.nickname)), NULL), '{}') AS nicknames
    FROM (
      SELECT wsst.id, wssn.series_id, wssn.nickname
      FROM (
        SELECT wsst.id
        FROM waifu_schema.series_table wsst
        LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
        WHERE (
          f_unaccent(wsst.name) ILIKE '%' || f_unaccent($1) || '%'
          OR f_unaccent(nickname) ILIKE '%' || f_unaccent($1) || '%'
        )
        UNION
        SELECT id
        FROM (
          SELECT wt2.id, (similarity(f_unaccent(lower($1)), wt2.name_lower)) as score
          FROM waifu_schema.series_table wt2
          WHERE wt2.name_lower % f_unaccent(lower($1))
          ORDER BY score DESC
          limit 10
        ) ws2
      ) ws
      JOIN waifu_schema.series_table wsst ON wsst.id = ws.id
      LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = ws.id
      WHERE (
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
          WHEN f_unaccent(wsst.name) ILIKE f_unaccent($1) THEN 0
          WHEN f_unaccent(wssn.nickname) ILIKE f_unaccent($1) THEN 1
          WHEN f_unaccent(wsst.name) ILIKE f_unaccent($1) || '%' THEN 2
          WHEN f_unaccent(wssn.nickname) ILIKE f_unaccent($1) || '%' THEN 3
          WHEN f_unaccent(wsst.name) ILIKE '%' || f_unaccent($1) || '%' THEN 4
          WHEN f_unaccent(wssn.nickname) ILIKE '%' || f_unaccent($1) || '%' THEN 5
        ELSE 6 END, wsst.name
      LIMIT 500
    ) wssn
    JOIN waifu_schema.series_table wsst ON wsst.id = wssn.id
    GROUP BY wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western
  ) t1
  ORDER BY
    CASE
      WHEN f_unaccent(name) ILIKE f_unaccent($1) THEN 0
      WHEN f_unaccent(name) ILIKE f_unaccent($1) || '%' THEN 1
      WHEN f_unaccent($1) ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(nicknames) ), ' ')) ) THEN 2
      WHEN f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(nicknames) ), ' ')) ) THEN 3
      WHEN f_unaccent(name) ILIKE '%' || f_unaccent($1) || '%' THEN 4
      WHEN '%' || f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(nicknames) ), ' ')) ) THEN 5
    ELSE 6 END, name
  LIMIT $2
`

var seriesRandom = `
SELECT wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western,
	COALESCE(array_remove(array_agg(DISTINCT(wssn.nickname)), NULL), '{}') AS nicknames
FROM waifu_schema.series_table wsst
LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
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
GROUP BY wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western
ORDER BY random()
LIMIT $1;
`

var seriesByID = `
SELECT wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western,
	COALESCE(array_remove(array_agg(DISTINCT(wssn.nickname)), NULL), '{}') AS nicknames
FROM waifu_schema.series_table wsst
LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
WHERE wsst.id = $1
GROUP BY wsst.id, wsst.name, description, image_url, release_date, nsfw, is_game, is_western;
`
