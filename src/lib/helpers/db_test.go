package helpers

import (
	"testing"
)

func TestGenerateParsedURLFromConfig(t *testing.T) {
	expected := "postgres://user:pass@host:1234/db?pool_max_conns=5&pool_max_conn_lifetime=5000ms&pool_max_conn_idle_time=2000ms"

	dbConfig := DbConfig{
		User:                    "user",
		Host:                    "host",
		Database:                "db",
		Password:                "pass",
		MaxConnections:          5,
		ConnectionTimeoutMillis: 5000,
		IdleTimeoutMillis:       2000,
		Port:                    1234,
	}

	received := GenerateParsedURLFromConfig(dbConfig)

	if received != expected {
		t.Errorf("generateParsedURLFromConfig(dbConfig) failed. Expected %v, Received %v", expected, received)
	} else {
		t.Logf("generateParsedURLFromConfig(dbConfig) passed.")
	}
}
