package helpers

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// MarshalJSONFile takes in a filename path and struct reference to
// parse out the contents of the json file into the struct.
func MarshalJSONFile(filename string, structRef interface{}) []byte {
	plan, _ := ioutil.ReadFile(filename)
	err := json.Unmarshal(plan, structRef)

	if err != nil {
		log.Error(err)
		log.Error("Problem with file: %s", filename)
		panic(err)
	}

	return plan
}
