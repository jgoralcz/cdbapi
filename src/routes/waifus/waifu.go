package waifus

import (
	"errors"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("oops"))
	// query := r.URL.Query()
	// name := query.Get("name")
	// if name == "" {
	// 	name = "Guest"
	// }
	// log.Printf("Received request for %s\n", name)
	// w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}
