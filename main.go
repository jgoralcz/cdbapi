package main

import "github.com/jgoralcz/cdbapi/src/server"

// @title Character Database (Go CDBAPI)
// @version 1.0
// @description Character Database is a simple example of golang using Echo and interacting with a PostgreSQL database.
// It is mainly composed of GET requests which retrieve characters and series.
// @termsOfService http://swagger.io/terms/
// @contact.name jgoralcz
// @contact.url https://github.com/jgoralcz/cdbapi
// @BasePath /api/v1
func main() {
	server.Run()
}
