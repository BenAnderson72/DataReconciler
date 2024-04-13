package main

import (
	"fmt"
	"log"

	"github.com/mjarkk/mongomock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"username"`
	Email string             `bson:"email"`
}

func main() {
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
