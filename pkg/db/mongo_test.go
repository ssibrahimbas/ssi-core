package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func TestMongoTransformId(t *testing.T) {
	mdb, err := NewMongo("mongodb://localhost:27017", "test")
	if err != nil {
		t.Fatal(err)
	}
	id := mdb.transformId("5c9a2d9e7b4d1e0001f0d6d4")
	if !primitive.IsValidObjectID(id.Hex()) {
		t.Fatal("id is not valid")
	}
}
