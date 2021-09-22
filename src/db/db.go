package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jgoralcz/cdbapi/src/lib/helpers"
)

var pool *pgxpool.Pool

func init() {
	var dbConfig helpers.DbConfig
	helpers.MarshalJSONFile("/usr/go/api.json", &dbConfig)

	parsedURL := helpers.GenerateParsedURLFromConfig(dbConfig)

	var err error
	pool, err = pgxpool.Connect(context.Background(), parsedURL)

	if err != nil {
		log.Fatal("Unable to connect to database\n", err)
	}

	log.Info("db user ", dbConfig.User, " logged in")
}

func Get(dest interface{}, statement string, params ...interface{}) error {
	err := pgxscan.Get(context.Background(), pool, dest, statement, params...)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func Select(dest interface{}, statement string, params ...interface{}) error {
	err := pgxscan.Select(context.Background(), pool, dest, statement, params...)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
