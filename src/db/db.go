package db

import (
	"context"
	"strconv"

	"github.com/georgysavva/scany/pgxscan"
	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jgoralcz/cdbapi/src/lib/helpers"
)

type dbConfig struct {
	User                    string
	Host                    string
	Database                string
	Password                string
	MaxConnections          int
	ConnectionTimeoutMillis int
	IdleTimeoutMillis       int
	Port                    int
}

var pool *pgxpool.Pool

func init() {
	var dbConfig dbConfig
	helpers.MarshalJSONFile("/usr/go/api.json", &dbConfig)

	parsedURL := generateParsedURLFromConfig(dbConfig)

	var err error
	pool, err = pgxpool.Connect(context.Background(), parsedURL)

	if err != nil {
		log.Fatal("Unable to connect to database\n", err)
	}

	log.Info("db user ", dbConfig.User, " logged in")
}

func generateParsedURLFromConfig(dbConfig dbConfig) string {
	if dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Host == "" || dbConfig.Database == "" {
		return ""
	}

	return "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host +
		":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Database + "?pool_max_conns=" + strconv.Itoa(dbConfig.MaxConnections) +
		"&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.ConnectionTimeoutMillis) + "ms&pool_max_conn_idle_time=" + strconv.Itoa(dbConfig.IdleTimeoutMillis) + "ms"
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
