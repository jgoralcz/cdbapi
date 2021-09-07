package main

import "github.com/jgoralcz/cdbapi/src/server"

// @title Bongo Bot Character Database
// @version 1.0
// @description Bongo Bot Character Database is an API used to get characters, series, and images stored from Bongo Bot.
// It is mainly composed of GET requests which retrieve characters and series.
// @termsOfService https://github.com/jgoralcz/cdbapi
// @contact.name jgoralcz
// @contact.url https://github.com/jgoralcz/cdbapi
// @BasePath /
func main() {
	server.Run()
}
