package routes

import (
	"github.com/gorilla/mux"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
	"github.com/jgoralcz/go_cdbapi/src/middleware"
	"github.com/urfave/negroni"
)

// Routes is a function that binds with http to handle particular routes
// It also includes special middleware when a panic occurs, logging each request,
// and specifying a default header.
func Routes() *negroni.Negroni {
	router := mux.NewRouter()

	router.HandleFunc("/{characters/random:characters/random\\/?}", CharacterRandomHandler).Methods("GET")
	router.HandleFunc("/{characters:characters\\/?}", CharacterNameHandler).Methods("GET")
	router.HandleFunc("/characters/{id}", CharacterByIDHandler).Methods("GET")
	router.HandleFunc("/characters/{id}/", CharacterByIDHandler).Methods("GET")

	n := negroni.New()

	// error middleware
	recovery := negroni.NewRecovery()
	env := helpers.GetEnvOrDefault("ENV", "LOCAL")
	if env != "LOCAL" {
		recovery.PrintStack = false
	}

	n.Use(recovery)
	n.Use(negroni.HandlerFunc(middleware.CommonHeaders))
	n.Use(negroni.HandlerFunc(middleware.Logging))
	n.UseHandler(router)

	return n
}
