package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_populateSourceDB(t *testing.T) {

	recCount := 100

	collection := populateSourceDB(recCount)

	nr, _ := collection.Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}
