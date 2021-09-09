package series

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v3"
)

// Series is a data structure that models a series. Be sure to check for nulls.
type Series struct {
	ID          int         `json:"id" example:"4731"`
	Name        string      `json:"name" example:"Flight Rising"`
	Description null.String `json:"description" swaggertype:"string" example:"Flight Rising is a free Virtual Pet/breeding sim focusing on dragons. It launched on June 9th, 2013, following a wildly successful Kickstarter campaign."`
	ImageURL    null.String `json:"image_url" swaggertype:"string" example:"https://cdn.bongo.best/series/4731/e1dcff95-fa7c-4995-92f2-b9dc9840ffe1.png"`
	ReleaseDate null.String `json:"release_date" swaggertype:"string" example:"null"`
	Nsfw        bool        `json:"nsfw" example:"false"`
	Game        bool        `json:"game" example:"true"`
	Western     bool        `json:"western" example:"true"`
	Nicknames   []string    `json:"nicknames"`
}

func handleRows(rows pgx.Rows) ([]Series, error) {
	series := []Series{}

	for rows.Next() {
		s := new(Series)
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&s.ImageURL,
			&s.ReleaseDate,
			&s.Nsfw,
			&s.Game,
			&s.Western,
			&s.Nicknames,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		series = append(series, *s)
	}

	rowsErr := rows.Err()
	if rowsErr != nil {
		log.Error(rowsErr)
		return nil, rowsErr
	}

	return series, nil
}

func handleRow(row pgx.Row) (*Series, error) {
	s := new(Series)

	err := row.Scan(
		&s.ID,
		&s.Name,
		&s.Description,
		&s.ImageURL,
		&s.ReleaseDate,
		&s.Nsfw,
		&s.Game,
		&s.Western,
		&s.Nicknames,
	)

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error(err)
		return s, err
	}

	return s, nil
}
