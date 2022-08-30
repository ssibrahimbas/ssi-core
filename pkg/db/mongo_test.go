package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func TestNewMongo(t *testing.T) {
	mdb, err := NewMongo("mongodb://localhost:27017", "test")
	if err != nil {
		t.Fatal(err)
	}
	if err := mdb.c.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}
}

func TestNewMongoWithWrongURI(t *testing.T) {
	mdb, err := NewMongo("mongod://localst:27011", "test")
	if mdb != nil {
		t.Fatal("mdb is not nil")
	}
	if err == nil {
		t.Fatal("err should not nil")
	}
}

func TestMongoDB_GetCollection(t *testing.T) {
	mdb, err := NewMongo("mongodb://localhost:27017", "test")
	if err != nil {
		t.Fatal(err)
	}
	col := mdb.GetCollection("test")
	if col == nil {
		t.Fatal("collection is nil")
	}
}
