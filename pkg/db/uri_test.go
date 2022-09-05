package db

import "testing"

func TestCalcMongoUri(t *testing.T) {
	uri := CalcMongoUri(UriParams{
		Host: "localhost",
		Port: "27017",
		User: "user",
		Pass: "pass",
		Db:   "db",
	})
	if uri != "mongodb://user:pass@localhost:27017/db" {
		t.Error("Wrong uri", uri)
	}
}
