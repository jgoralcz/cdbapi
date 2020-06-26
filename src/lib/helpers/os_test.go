package helpers

import (
	"os"
	"testing"
)

func TestGetEnvOrDefaultSetEnv(t *testing.T) {
	os.Setenv("FOO", "bar")
	expected := "bar"

	value := GetEnvOrDefault("FOO", "test")

	if value != expected {
		t.Errorf("GetEnvOrDefault(\"FOO\", \"test\") failed. Expected %v, Received %v", expected, value)
	} else {
		t.Logf("GetEnvOrDefault(\"FOO\", \"test\") passed.")
	}
}

func TestGetEnvOrDefaultNoEnv(t *testing.T) {
	expected := "fail"

	fail := GetEnvOrDefault("BATMAN", "fail")

	if fail != expected {
		t.Errorf("GetEnvOrDefault(\"BATMAN\", \"fail\") failed. Expected %v, Received %v", fail, expected)
	} else {
		t.Logf("GetEnvOrDefault(\"BATMAN\", \"fail\") passed.")
	}
}
