// package db

// import (
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/jackc/pgx/v4"
// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// // const poolQuery = async (query, paramsArray) => {
// //   const client = await pool.connect();
// //   try {
// //     const result = await client.query(query, paramsArray);

// //     if (!result || !result.rows || !result.rowCount) return undefined;

// //     return result.rows;
// //   } finally {
// //     client.release();
// //   }
// // };

// // module.exports = {
// //   poolQuery,
// // };

// // type App struct {
// // 	Repo Repository
// // }

// // type Repository interface {
// // 	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
// // 	Query(query string, args ...interface{}) (pgx.Rows, error)
// // 	QueryRow(query string, args ...interface{}) pgx.Row
// // 	Close() error
// // }

// // type PGRepository struct {
// // 	*pgxpool.Pool
// // }

// // type ConnConfig pgx.Connconfig {

// // }

// var pool, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

// var path = "/Users/Josh/Documents/GitHub/go_cdbapi/api.json"

// func init() {
// 	plan, _ := ioutil.ReadFile("/Users/Josh/Documents/GitHub/go_cdbapi/api.json")
// 	var data interface{}
// 	err := json.Unmarshal(plan, &data)

// 	if err != nil {
// 		log.Fatalf("Could not find file: %s", path)
// 	}
// }

// // if err != nil {
// // 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// // 	os.Exit(1)
// // }

// func PoolQuery(statement string, params ...interface{}) pgx.Row {
// 	tx, err := pool.Begin(context.Background())
// 	if err != nil {
// 		return nil
// 	}
// 	defer tx.Rollback(context.Background())

// 	res := tx.QueryRow(context.Background(), statement, params...)
// 	if err != nil {
// 		return nil
// 	}

// 	err = tx.Commit(context.Background())
// 	if err != nil {
// 		// TODO: check if this is neede
// 		tx.Rollback(context.Background())
// 		return nil
// 	}

// 	return res
// }

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

var Pool *pgxpool.Pool

func init() {
	var dbConfig dbConfig
	helpers.MarshalJSONFile("/Users/Josh/Documents/GitHub/go_cdbapi/api.json", &dbConfig)

	parsedURL := "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host +
		":" + strconv.Itoa(dbConfig.Port) + "/" + dbConfig.Database + "?pool_max_conns=" + strconv.Itoa(dbConfig.MaxConnections) +
		"&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.ConnectionTimeoutMillis) + "ms&pool_max_conn_lifetime=" + strconv.Itoa(dbConfig.IdleTimeoutMillis) + "ms"

	var err error
	Pool, err = pgxpool.Connect(context.Background(), parsedURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Printf("%s", dbConfig.User)
}

func PoolQueryRows(statement string, params ...interface{}) pgx.Rows {
	rows, err := Pool.Query(context.Background(), statement, params...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

func PoolQueryRow(statement string, params ...interface{}) pgx.Row {
	// tx, err := Pool.Begin(context.Background())
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }
	// defer tx.Rollback(context.Background())

	row := Pool.QueryRow(context.Background(), statement, params...)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }

	// err = tx.Commit(context.Background())
	// if err != nil {
	// 	// TODO: check if this is needed
	// 	tx.Rollback(context.Background())
	// 	log.Println(err)
	// 	return nil
	// }
	// log.Print(err)
	log.Print(row)

	return row
}
