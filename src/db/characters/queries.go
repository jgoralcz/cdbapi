package characters

var characterSearchShort = `
SELECT wt.id, name, description, image_url, image_url_clean,
	nsfw, series, series_id, husbando,
	count, position, last_edit_by, last_edit_date,
	nicknames, spoiler_nicknames,
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
			UNION
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_appears_in_id AS series_id
				FROM waifu_schema.series_appears_in_series wsais
				WHERE wsais.series_id = wt.series_id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
FROM (
	SELECT t1.name, (
		SELECT
			CASE t1.nsfw WHEN TRUE then TRUE
				ELSE wsst.nsfw
			END
	) AS nsfw, wsst.name AS series, series_id, husbando, t1.image_url, t1.image_url_clean_discord,
		t1.image_url_clean, t1.description, t1.id, t1.last_edit_by, t1.last_edit_date,
		COALESCE(array_remove(array_agg(DISTINCT(wscn.nickname)), NULL), '{}') AS nicknames,
		COALESCE(array_remove(array_agg(DISTINCT(CASE WHEN wscn.is_spoiler = TRUE THEN wscn.nickname ELSE NULL END)), NULL), '{}') AS spoiler_nicknames,
		count, position
	FROM (
		SELECT name, nsfw, series_id, husbando,
			image_url, image_url_clean_discord, image_url_clean,
			description, wt.id, last_edit_by, last_edit_date
		FROM (
			SELECT id
			FROM (
				SELECT id
				FROM waifu_schema.waifu_table wt
				WHERE tsv @@ to_tsquery(lower(f_unaccent($1)) || ':*')
			) ws
			UNION
			SELECT id
			FROM (
				SELECT character_id AS id
				FROM waifu_schema.character_nicknames
				WHERE tsv @@ to_tsquery(lower(f_unaccent($1)) || ':*')
			) wscn
		) ws
		JOIN waifu_schema.waifu_table wt ON wt.id = ws.id
		LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = ws.id
		ORDER BY
			CASE
				WHEN name_lower LIKE f_unaccent(lower($1)) THEN 0
				WHEN nickname_lower LIKE f_unaccent(lower($1)) THEN 1
				WHEN name_lower LIKE f_unaccent(lower($1)) || '%' THEN 2
				WHEN nickname_lower LIKE f_unaccent(lower($1)) || '%' THEN 3
				WHEN name_lower LIKE '%' || f_unaccent(lower($1)) || '%' THEN 4
				WHEN nickname_lower LIKE '%' || f_unaccent(lower($1)) || '%' THEN 5
			ELSE 6 END, wt.name
	) t1
	JOIN waifu_schema.series_table wsst ON wsst.id = t1.series_id
	LEFT JOIN mv_rank_claim_waifu mv ON mv.waifu_id = t1.id
	LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = t1.id
	WHERE (
		('false' = $3 AND t1.nsfw = FALSE)
		OR ('true' = $3 AND t1.nsfw = TRUE)
		OR t1.nsfw IS NOT NULL
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
	GROUP BY t1.name, t1.series_id, t1.nsfw, wsst.name, wsst.nsfw, husbando, t1.image_url, t1.image_url_clean_discord, t1.image_url_clean, t1.description, t1.id, t1.last_edit_by, t1.last_edit_date, count, position
) wt
ORDER BY
	CASE
		WHEN f_unaccent(lower(name)) LIKE f_unaccent(lower($1)) THEN 0
		WHEN f_unaccent(lower(name)) LIKE f_unaccent(lower($1)) || '%' THEN 1
		WHEN f_unaccent(lower(name)) ILIKE '%' || f_unaccent(lower($1)) || '%' THEN 2
		WHEN '%' || f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(nicknames) ), ' ')) ) THEN 3
	ELSE 4 END, name
LIMIT $2;
`

var characterSearch = `
SELECT wt.id, name, description, image_url, image_url_clean,
	nsfw, series, series_id, husbando,
	count, position, last_edit_by, last_edit_date,
	nicknames, spoiler_nicknames,
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
			UNION
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_appears_in_id AS series_id
				FROM waifu_schema.series_appears_in_series wsais
				WHERE wsais.series_id = wt.series_id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
FROM (
	SELECT t1.name, (
		SELECT
			CASE t1.nsfw WHEN TRUE then TRUE
				ELSE wsst.nsfw
			END
	) AS nsfw, wsst.name AS series, series_id, husbando, t1.image_url, t1.image_url_clean_discord,
		t1.image_url_clean, t1.description, t1.id, t1.last_edit_by, t1.last_edit_date,
		COALESCE(array_remove(array_agg(DISTINCT(wscn.nickname)), NULL), '{}') AS nicknames,
		COALESCE(array_remove(array_agg(DISTINCT(CASE WHEN wscn.is_spoiler = TRUE THEN wscn.nickname ELSE NULL END)), NULL), '{}') AS spoiler_nicknames,
		count, position
	FROM (
		SELECT name, nsfw, series_id, husbando,
			image_url, image_url_clean_discord, image_url_clean,
			description, wt.id, last_edit_by, last_edit_date
		FROM (
			SELECT id
        FROM (
          SELECT wt.id
          FROM waifu_schema.waifu_table wt
          WHERE name_lower like '%' || f_unaccent(lower($1)) || '%'
        ) ws
        UNION
        SELECT id
        FROM (
          SELECT wt2.id, (similarity(f_unaccent(lower($1)), wt2.name_lower)) as score
          FROM waifu_schema.waifu_table wt2
          WHERE wt2.name_lower % f_unaccent(lower($1))
          ORDER BY score DESC
          limit 10
        ) ws2
        UNION
        SELECT id
        FROM (
          SELECT wscn.character_id AS id
          FROM waifu_schema.character_nicknames wscn
          JOIN waifu_schema.waifu_table wt ON wt.id = wscn.character_id
          WHERE nickname_lower like '%' || f_unaccent(lower($1)) || '%'
        ) wscn
        UNION
        SELECT id
        FROM (
          SELECT wscn2.character_id AS id, (similarity(f_unaccent(lower($1)), wscn2.nickname_lower)) as score
          FROM waifu_schema.character_nicknames wscn2
          WHERE wscn2.nickname_lower % f_unaccent(lower($1))
          ORDER BY score DESC
          limit 10
        ) wscn2
		) ws
		JOIN waifu_schema.waifu_table wt ON wt.id = ws.id
		LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = ws.id
		ORDER BY
			CASE
				WHEN name_lower LIKE f_unaccent(lower($1)) THEN 0
				WHEN nickname_lower LIKE f_unaccent(lower($1)) THEN 1
				WHEN name_lower LIKE f_unaccent(lower($1)) || '%' THEN 2
				WHEN nickname_lower LIKE f_unaccent(lower($1)) || '%' THEN 3
				WHEN name_lower LIKE '%' || f_unaccent(lower($1)) || '%' THEN 4
				WHEN nickname_lower LIKE '%' || f_unaccent(lower($1)) || '%' THEN 5
			ELSE 6 END, wt.name
	) t1
	JOIN waifu_schema.series_table wsst ON wsst.id = t1.series_id
	LEFT JOIN mv_rank_claim_waifu mv ON mv.waifu_id = t1.id
	LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = t1.id
	WHERE (
		('false' = $3 AND t1.nsfw = FALSE)
		OR ('true' = $3 AND t1.nsfw = TRUE)
		OR t1.nsfw IS NOT NULL
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
	GROUP BY t1.name, t1.series_id, t1.nsfw, wsst.name, wsst.nsfw, husbando, t1.image_url, t1.image_url_clean_discord, t1.image_url_clean, t1.description, t1.id, t1.last_edit_by, t1.last_edit_date, count, position
) wt
ORDER BY
	CASE
		WHEN f_unaccent(lower(name)) LIKE f_unaccent(lower($1)) THEN 0
		WHEN f_unaccent(lower(name)) LIKE f_unaccent(lower($1)) || '%' THEN 1
		WHEN f_unaccent(lower(name)) ILIKE '%' || f_unaccent(lower($1)) || '%' THEN 2
		WHEN '%' || f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(nicknames) ), ' ')) ) THEN 3
	ELSE 4 END, name
LIMIT $2;
`

var characterRandom = `
SELECT *
FROM (
	SELECT wt.id, name, description, image_url, image_url_clean, nsfw, 
	series, series_id, husbando, count, position, last_edit_by, last_edit_date,
	COALESCE(array_remove(array_agg(DISTINCT(wscn.nickname)), NULL), '{}') AS nicknames,
	COALESCE(array_remove(array_agg(DISTINCT(CASE WHEN wscn.is_spoiler = TRUE THEN wscn.nickname ELSE NULL END)), NULL), '{}') AS spoiler_nicknames,
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
			UNION
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_appears_in_id AS series_id
				FROM waifu_schema.series_appears_in_series wsais
				WHERE wsais.series_id = wt.series_id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
	FROM (
		SELECT ws.name, (
		SELECT
			CASE ws.nsfw WHEN TRUE then TRUE
				ELSE wsst.nsfw
			END
		) AS nsfw, wsst.name AS series, wsst.id AS series_id, ws.husbando,
		ws.image_url, ws.image_url_clean_discord, ws.image_url_clean, ws.description,
		ws.id, ws.last_edit_by, ws.last_edit_date, count, position
		
		FROM mv_random_waifu_series mvws
		JOIN waifu_schema.waifu_table ws ON ws.id = mvws.id
		JOIN waifu_schema.series_table wsst ON wsst.id = ws.series_id
		LEFT JOIN mv_rank_claim_waifu mv ON mv.waifu_id = ws.id
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
		LIMIT 1000
	) wt
	LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = wt.id
	GROUP BY wt.id, name, nsfw, series, series_id, husbando, image_url, image_url_clean_discord,
	image_url_clean, description, last_edit_by, last_edit_date, count, position
) t1
ORDER BY random()
LIMIT $1;
`

var characterByID = `
SELECT wt.id, name, description, image_url, image_url_clean, nsfw, 
	series, series_id, husbando, count, position, last_edit_by, last_edit_date,
	COALESCE(array_remove(array_agg(DISTINCT(wscn.nickname)), NULL), '{}') AS nicknames,
	COALESCE(array_remove(array_agg(DISTINCT(CASE WHEN wscn.is_spoiler = TRUE THEN wscn.nickname ELSE NULL END)), NULL), '{}') AS spoiler_nicknames,
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
			UNION
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_appears_in_id AS series_id
				FROM waifu_schema.series_appears_in_series wsais
				WHERE wsais.series_id = wt.series_id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
FROM (
	SELECT ws.name, (
		SELECT
			CASE ws.nsfw WHEN TRUE then TRUE
				ELSE wsst.nsfw
			END
	) AS nsfw, wsst.name AS series, wsst.id AS series_id, ws.husbando,
	ws.image_url, ws.image_url_clean_discord, ws.image_url_clean, ws.description,
	ws.id, ws.last_edit_by, ws.last_edit_date, count, position
	
	FROM waifu_schema.waifu_table ws
	JOIN waifu_schema.series_table wsst ON wsst.id = ws.series_id
	LEFT JOIN mv_rank_claim_waifu mv ON mv.waifu_id = ws.id
	WHERE ws.id = $1
) wt
LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = wt.id
GROUP BY wt.id, name, nsfw, series, series_id, husbando, image_url, image_url_clean_discord,
image_url_clean, description, last_edit_by, last_edit_date, count, position;
`

var charactersBySeries = `
SELECT wt.id, name, description, image_url, image_url_clean,
	nsfw, series, series_id, husbando,
	count, position, last_edit_by, last_edit_date,
	nicknames, spoiler_nicknames,
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
			UNION
			SELECT wsst.name AS series, wsst.id AS series_id, wsst.nsfw AS nsfw, wsst.is_game AS game, wsst.is_western AS western
			FROM (
				SELECT series_appears_in_id AS series_id
				FROM waifu_schema.series_appears_in_series wsais
				WHERE wsais.series_id = wt.series_id
			) wsai
			JOIN waifu_schema.series_table wsst ON wsst.id = wsai.series_id
		) item
	) AS appears_in
  FROM (
    SELECT name, nsfw, series, husbando, image_url, image_url_clean_discord, w.series_id,
      image_url_clean, description, w.id, last_edit_by, last_edit_date, count, position,
      COALESCE(array_remove(array_agg(DISTINCT(wscn.nickname)), NULL), '{}') AS nicknames,
      COALESCE(array_remove(array_agg(DISTINCT(CASE WHEN wscn.is_spoiler = TRUE THEN wscn.nickname ELSE NULL END)), NULL), '{}') AS spoiler_nicknames,
			COALESCE(array_remove(array_agg(DISTINCT(w.nickname)), NULL), '{}') AS series_nicknames
		FROM (
      SELECT ws.name, (
        SELECT
          CASE ws.nsfw WHEN TRUE then TRUE
            ELSE wsst.nsfw
          END
      ) AS nsfw, wsst.name AS series, wsst.id AS series_id, ws.husbando,
        ws.image_url, ws.image_url_clean_discord, ws.image_url_clean, ws.description,
        ws.id, ws.last_edit_by, ws.last_edit_date, wsst.nickname
      FROM (
        SELECT wsst.id, name, nsfw, wssn.nickname
        FROM (
          SELECT wsst.id
          FROM waifu_schema.series_table wsst
          LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
          WHERE (
            f_unaccent(wsst.name) ILIKE '%' || f_unaccent($1) || '%'
            OR f_unaccent(nickname) ILIKE '%' || f_unaccent($1) || '%'
            -- OR nickname_lower % f_unaccent(lower($1))
            -- OR name_lower % f_unaccent(lower($1))
          )
        ) ws
        JOIN waifu_schema.series_table wsst ON wsst.id = ws.id
        LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = ws.id
        ORDER BY
          CASE
            WHEN name_lower LIKE lower(f_unaccent($1)) THEN 0
            WHEN nickname_lower LIKE lower(f_unaccent($1)) THEN 1
            WHEN name LIKE lower(f_unaccent($1)) || '%' THEN 2
            WHEN nickname_lower LIKE lower(f_unaccent($1)) || '%' THEN 3
            WHEN name_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 4
            WHEN nickname_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 5
          ELSE 6 END, wsst.name
        LIMIT 100
      ) wsst
      JOIN waifu_schema.waifu_table ws ON ws.series_id = wsst.id
      UNION
      SELECT ws.name, (
        SELECT
          CASE ws.nsfw WHEN TRUE then TRUE
            ELSE wsst.nsfw
          END
      ) AS nsfw, wsst.name AS series, wsst.id AS series_id, ws.husbando,
        ws.image_url, ws.image_url_clean_discord, ws.image_url_clean, ws.description,
        ws.id, ws.last_edit_by, ws.last_edit_date, wsst.nickname
        FROM (
          SELECT wsst.id, name, nsfw, wssn.nickname
          FROM (
            SELECT wsst.id
            FROM waifu_schema.series_table wsst
            LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
            WHERE (
              f_unaccent(wsst.name) ILIKE '%' || f_unaccent($1) || '%'
              OR f_unaccent(nickname) ILIKE '%' || f_unaccent($1) || '%'
              -- OR nickname_lower % f_unaccent(lower($1))
              -- OR name_lower % f_unaccent(lower($1))
            )
          ) ws
          JOIN waifu_schema.series_table wsst ON wsst.id = ws.id
          LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = ws.id
          ORDER BY
            CASE
              WHEN name_lower LIKE lower(f_unaccent($1)) THEN 0
              WHEN nickname_lower LIKE lower(f_unaccent($1)) THEN 1
              WHEN name LIKE lower(f_unaccent($1)) || '%' THEN 2
              WHEN nickname_lower LIKE lower(f_unaccent($1)) || '%' THEN 3
              WHEN name_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 4
              WHEN nickname_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 5
            ELSE 6 END, wsst.name
          LIMIT 100
      ) wsst
      -- character subseries
      JOIN waifu_schema.appears_in wsai ON wsai.series_id = wsst.id

      -- character data
      JOIN waifu_schema.waifu_table ws ON ws.id = wsai.waifu_id
      UNION
      SELECT ws.name, (
        SELECT
          CASE ws.nsfw WHEN TRUE then TRUE
            ELSE wsst.nsfw
          END
      ) AS nsfw, wsst.name AS series, wsst.id AS series_id, ws.husbando,
        ws.image_url, ws.image_url_clean_discord, ws.image_url_clean, ws.description,
        ws.id, ws.last_edit_by, ws.last_edit_date, wsst.nickname
      FROM (
        SELECT wsst.id, name, nsfw, wssn.nickname
        FROM (
          SELECT wsst.id
          FROM waifu_schema.series_table wsst
          LEFT JOIN waifu_schema.series_nicknames wssn ON wssn.series_id = wsst.id
          WHERE (
            f_unaccent(wsst.name) ILIKE '%' || f_unaccent($1) || '%'
            OR f_unaccent(nickname) ILIKE '%' || f_unaccent($1) || '%'
            -- OR nickname_lower % f_unaccent(lower($1))
            -- OR name_lower % f_unaccent(lower($1))
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
        ORDER BY
          CASE
            WHEN name_lower LIKE lower(f_unaccent($1)) THEN 0
            WHEN nickname_lower LIKE lower(f_unaccent($1)) THEN 1
            WHEN name LIKE lower(f_unaccent($1)) || '%' THEN 2
            WHEN nickname_lower LIKE lower(f_unaccent($1)) || '%' THEN 3
            WHEN name_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 4
            WHEN nickname_lower LIKE '%' || lower(f_unaccent($1)) || '%' THEN 5
          ELSE 6 END, wsst.name
        LIMIT 100
      ) wsst
      -- subseries
      JOIN waifu_schema.series_appears_in_series wssais ON (wssais.series_id = wsst.id OR wssais.series_appears_in_id = wsst.id)

      -- character data
      JOIN waifu_schema.waifu_table ws ON ws.series_id = wssais.series_id
    ) w
    LEFT JOIN mv_rank_claim_waifu mv ON mv.waifu_id = w.id
    LEFT JOIN waifu_schema.character_nicknames wscn ON wscn.character_id = w.id
    GROUP BY name, nsfw, series, series_id, husbando, image_url, image_url_clean_discord,
      image_url_clean, description, w.id, last_edit_by, last_edit_date, count, position
  ) wt
  ORDER BY
    CASE
      WHEN f_unaccent(series) ILIKE f_unaccent($1) THEN 0
      WHEN f_unaccent($1) ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(series_nicknames) ), ' ')) ) THEN 1
      WHEN f_unaccent(series) ILIKE f_unaccent($1) || '%' THEN 2
      WHEN f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(series_nicknames) ), ' ')) ) THEN 3
      WHEN f_unaccent(series) ILIKE '%' || f_unaccent($1) || '%' THEN 4
      WHEN '%' || f_unaccent($1) || '%' ILIKE ANY ( SELECT UNNEST( string_to_array(f_unaccent( UNNEST(series_nicknames) ), ' ')) ) THEN 5
    ELSE 6 END, series, name
	LIMIT $2 
	OFFSET $3;
`

var characterImagesByIDOffsetLimit = `
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
	OFFSET $3;
`
