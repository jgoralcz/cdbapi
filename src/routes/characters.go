package routes

import (
	"log"
	"net/http"

	"github.com/jgoralcz/go_cdbapi/src/db"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
)

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
}

func CharacterRandomHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	initLimit := query.Get("limit")
	nsfw := query.Get("nsfw")
	western := query.Get("western")
	game := query.Get("game")

	limit := helpers.MaxLimit(initLimit, 1, 20)

	json := db.GetRandomCharacter(limit, nsfw, western, game)

	log.Println(json)

	w.Write(json)
}
