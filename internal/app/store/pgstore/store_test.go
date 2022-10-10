package pgstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://postgres:qwe123@localhost:5432/restapi_test_gopher?sslmode=disable"
	}

	os.Exit(m.Run())
}
