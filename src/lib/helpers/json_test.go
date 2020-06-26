package helpers

import (
	"testing"
)

type mock struct {
	test string
}

func TestMarshalJSONFilePanic(t *testing.T) {
	var m mock

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MarshalJSONFile did not panic.")
		} else {
			t.Logf("MarshalJSONFile (panic) was successful")
		}
	}()
	MarshalJSONFile("test", &m)
}
