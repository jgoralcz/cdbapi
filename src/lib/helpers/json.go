package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// MarshalJSONFile takes in a filename path and struct reference to
// parse out the contents of the json file into the struct.
func MarshalJSONFile(filename string, structRef interface{}) []byte {
	plan, _ := ioutil.ReadFile("filename")
	err := json.Unmarshal(plan, structRef)

	if err != nil {
		log.Printf("%s", err)
		log.Printf("Problem with file: %s", filename)
		panic(err)
	}

	return plan
}
