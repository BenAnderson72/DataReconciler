package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	Data "github.com/BenAnderson72/DataReconciler/data"

	"github.com/mjarkk/mongomock"
)

func populateTargetDB0(recCount int) mongomock.Collection {
	db := mongomock.NewDB()
	collection := db.Collection("transactions")

	n := 0
	for n < recCount {
		pmnt := Data.GenPayment()

		err := collection.Insert(pmnt)
		if err != nil {
			log.Fatal(err)
		}
		n++
	}

	return *collection

}

// type PgxIface interface {
// 	Begin(context.Context) (pgx.Tx, error)
// 	Close(context.Context) error
// }

// func recordStats(db PgxIface, userID, productID int) (err error) {
// 	tx, err := db.Begin(context.Background())
// 	if err != nil {
// 		return
// 	}
// 	defer func() {
// 		switch err {
// 		case nil:
// 			err = tx.Commit(context.Background())
// 		default:
// 			_ = tx.Rollback(context.Background())
// 		}
// 	}()
// 	sql := "UPDATE products SET views = views + 1"
// 	if _, err = tx.Exec(context.Background(), sql); err != nil {
// 		return
// 	}
// 	sql = "INSERT INTO product_viewers (user_id, product_id) VALUES ($1, $2)"
// 	if _, err = tx.Exec(context.Background(), sql, userID, productID); err != nil {
// 		return
// 	}
// 	return
// }

func populateSourceDB(recCount int, file_sourceDB string) {

	csvFile, err := os.Create(file_sourceDB)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	n := 0
	for n < recCount {
		pmnt := Data.GenPayment()

		w.Write([]string{pmnt.Time.String(), pmnt.Sender_Account, pmnt.Receiver_Account, pmnt.TransactionID, fmt.Sprintf("%.2f", pmnt.Amount), pmnt.Currency, pmnt.Reference})

		n++
	}

	w.Flush()

}

func populateTargetDB(file_sourceDB string) mongomock.Collection {

	csvFile, err := os.Open(file_sourceDB)

	if err != nil {
		log.Fatalf("Error while reading the file: %s", err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("Error while reading records: %s", err)
	}

	db := mongomock.NewDB()
	collection := db.Collection("transactions")

	for _, rec := range records {
		pmnt := Data.LoadPayment(rec)
		err := collection.Insert(pmnt)
		if err != nil {
			log.Fatal(err)
		}

	}
	return *collection
}

// func main0() {
// 	db := mongomock.NewDB()
// 	collection := db.Collection("users")
// 	err := collection.Insert(User{
// 		ID:    primitive.NewObjectID(),
// 		Name:  "test",
// 		Email: "example@example.org",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// nr, _ := db.Collection("users").Count(bson.M{})
// 	// fmt.Printf("Found %d items", nr)

// 	user := User{}
// 	err = collection.FindFirst(&user, bson.M{"username": "test"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Found user: %+v\n", user)
// 	// After exit the database data is gone
// }
