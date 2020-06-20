package routes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jgoralcz/go_cdbapi/src/db/characters"
	"github.com/jgoralcz/go_cdbapi/src/helpers"
)

// CharacterByIDHandler handles routes when a user searches by a character ID.
// Returns 400 if the id is not an int, and returns 404 if the id is not found.
func CharacterByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strID := params["id"]

	id, err := strconv.Atoi(strID)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{ "error": "Must have a valid id parameter" }`))
		return
	}

	json := characters.GetCharacterByID(id)

	if string(json) == "[]" {
		w.WriteHeader(404)
		w.Write([]byte("{ \"error\": \"Could not find a character with id " + strID + "\" }"))
		return
	}

	if json == nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(json)
}

// CharacterNameHandler handles the logic for searching for a character with the
// user's specified parameters that act as filters: limit (int), nsfw (true or false),
// western (true or false), game (true or false)
func CharacterNameHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	initLimit := query.Get("limit")
	nsfw := query.Get("nsfw")
	western := query.Get("western")
	game := query.Get("game")
	name := query.Get("name")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	if name == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{ "error": "Must have a valid name query parameter"}`))
		return
	}

	json := characters.SearchCharacter(name, limit, isNSFW, isWestern, isGame)

	if json == nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(json)
}

// CharacterRandomHandler handles the logic for a random character request with
// the user's specified parameters that act as filters: limit (int), nsfw (true or false),
// western (true or false), game (true or false)
func CharacterRandomHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	initLimit := query.Get("limit")
	nsfw := query.Get("nsfw")
	western := query.Get("western")
	game := query.Get("game")

	limit := helpers.MaxLimit(initLimit, 1, 20)
	isNSFW := helpers.DefaultBoolean(nsfw)
	isWestern := helpers.DefaultBoolean(western)
	isGame := helpers.DefaultBoolean(game)

	json := characters.GetRandomCharacter(limit, isNSFW, isWestern, isGame)

	if json == nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(json)
}
