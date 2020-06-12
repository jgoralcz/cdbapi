package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/jgoralcz/go_cdbapi/src/db"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
	"github.com/jgoralcz/go_cdbapi/src/middleware"
	"github.com/jgoralcz/go_cdbapi/src/routes/waifus"
	"github.com/urfave/negroni"
)

func Routes() *negroni.Negroni {
	router := mux.NewRouter()

	router.HandleFunc("/characters", waifus.Handler).Methods("GET")

	n := negroni.New()

	// error middleware
	recovery := negroni.NewRecovery()
	env := helpers.GetEnvOrDefault("ENV", "LOCAL")
	if env != "LOCAL" {
		recovery.PrintStack = false
	}

	n.Use(recovery)
	n.Use(negroni.HandlerFunc(middleware.LoggingMiddleware))
	n.UseHandler(router)

	test := db.GetRandomCharacter(5)
	log.Printf(test)

	return n
}
