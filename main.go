package main

import (
	"fmt"
	"log"

	Data "github.com/BenAnderson72/DataReconciler/data"

	"github.com/mjarkk/mongomock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"username"`
	Email string             `bson:"email"`
}

func populateSourceDB(recCount int) mongomock.Collection {
	db := mongomock.NewDB()
	collection := db.Collection("transactions")

	n := 0
	for n < recCount {
		pmnt := Data.GenPaymentSource()

		err := collection.Insert(pmnt)
		if err != nil {
			log.Fatal(err)
		}
		n++
	}

	return *collection

	// After exit the database data is gone
}

func main0() {
	db := mongomock.NewDB()
	collection := db.Collection("users")
	err := collection.Insert(User{
		ID:    primitive.NewObjectID(),
		Name:  "test",
		Email: "example@example.org",
	})
	if err != nil {
		log.Fatal(err)
	}

	// nr, _ := db.Collection("users").Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	user := User{}
	err = collection.FindFirst(&user, bson.M{"username": "test"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found user: %+v\n", user)
	// After exit the database data is gone
}
