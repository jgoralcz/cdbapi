package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func MarshalJSONFile (filename string, structRef interface{}) []byte {
	plan, _ := ioutil.ReadFile(filename)
	err := json.Unmarshal(plan, structRef)

	if err != nil {
		log.Printf("%s", err)
		log.Fatalf("Could not find file: %s", filename)
	}
	
	return plan
}
