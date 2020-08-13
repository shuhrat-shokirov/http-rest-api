package sqlstore

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M)  {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl =="" {
 		databaseUrl = "host=localhost dbname=restapi_test user=user password=pass sslmode=disable"
	}

	os.Exit(m.Run())
}
