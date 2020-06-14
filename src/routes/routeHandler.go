package routes

import (
	"github.com/gorilla/mux"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
	"github.com/jgoralcz/go_cdbapi/src/middleware"
	"github.com/urfave/negroni"
)

// Routes is a function that binds with http to handle particular routes
func Routes() *negroni.Negroni {
	router := mux.NewRouter()

	router.HandleFunc("/characters", CharacterHandler).Methods("GET")
	router.HandleFunc("/characters/random", CharacterRandomHandler).Methods("GET")

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

	return n
}
