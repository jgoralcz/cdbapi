package main

import "github.com/jgoralcz/cdbapi/src/server"

// @title Character Database
// @version 1.0
// @description Character Database is an API used to get characters, series, and images stored from Bongo Bot.
// It is mainly composed of GET requests which retrieve characters and series.
// @termsOfService http://swagger.io/terms/
// @contact.name jgoralcz
// @contact.url https://github.com/jgoralcz/cdbapi
// @BasePath /v1
func main() {
	server.Run()
}
