package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
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
	helpers.MarshalJSONFile("/Users/Josh/Documents/GitHub/go_cdbapi/api.json", &dbConfig)

	parsedURL := "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host +
		":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Database + "?pool_max_conns=" + strconv.Itoa(dbConfig.MaxConnections) +
		"&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.ConnectionTimeoutMillis) + "ms&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.IdleTimeoutMillis) + "ms"

	var err error
	pool, err = pgxpool.Connect(context.Background(), parsedURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Printf("%s", dbConfig.User)
}

// PoolQueryRows queries the database pool and retreives multiple rows.
func PoolQueryRows(statement string, params ...interface{}) pgx.Rows {
	rows, err := pool.Query(context.Background(), statement, params...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

// PoolQueryRow queries the database pool and retreives a single row.
func PoolQueryRow(statement string, params ...interface{}) pgx.Row {
	return pool.QueryRow(context.Background(), statement, params...)
}
