package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"

	Data "github.com/BenAnderson72/DataReconciler/data"
	"github.com/mjarkk/mongomock"
	"github.com/r3labs/diff/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_populateTargetDB0(t *testing.T) {

	collection := populateTargetDB0(recCount)

	nr, _ := collection.Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}

func corruptTargetDB(file_sourceDB string, recCount int, collection *mongomock.Collection) {

	records := getSourceData(file_sourceDB)

	// Corrupt recCount random records on target DB
	txs := []string{}
	n := 0
	for n < recCount {
		i := rand.IntN(1000)
		tid := records[i][3]
		// make sure they are unique
		for slices.Contains(txs, tid) {
			i = rand.IntN(1000)
			tid = records[i][3]
		}

		pt := Data.PaymentType{}
		err := collection.FindFirst(&pt, bson.M{"transaction_id": tid})

		if err != nil {
			log.Fatal(err)
		}

		pt.Amount = 0
		pt.Reference += " CORRUPT"

		err = collection.ReplaceFirst(bson.M{"transaction_id": tid}, pt)

		if err != nil {
			log.Fatal(err)
		}

		txs = append(txs, tid)
		n++
	}

}

// This method populates the target mongoDB from the source DB (source_db.csv)
func Test_populateTargetDB(t *testing.T) {

	collection := populateTargetDB(file_sourceDB)

	nr, _ := collection.Count(bson.M{})

	// Check record counts
	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}

// This method populates the target mongoDB from the source DB (source_db.csv)
func Test_reconcile(t *testing.T) {

	collection := populateTargetDB(file_sourceDB)

	nr, _ := collection.Count(bson.M{})

	// Check record counts
	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

	corruptTargetDB(file_sourceDB, 10, &collection)

	records := getSourceData(file_sourceDB)

	// px := Data.PaymentType{}
	// _ = collection.FindFirst(&px, bson.M{"transaction_id": txs[0]})
	// fmt.Print(px.Amount)

	for _, rec := range records {
		pt := Data.PaymentType{}

		// Check Transaction IDs exist
		err := collection.FindFirst(&pt, bson.M{"transaction_id": rec[3]})

		// if txs[0] == rec[3] {
		// 	pt.Sender_Account += " MESSED WITH"
		// }

		if strings.HasSuffix(pt.Reference, "CORRUPT") {
			fmt.Println("here")
		}

		if err != nil {
			log.Fatal(err)
		}

		ps := Data.LoadPayment(rec)

		changelog, _ := diff.Diff(ps, pt)

		if len(changelog) != 0 {
			t.Errorf("%v", changelog)
		}

		// if !reflect.DeepEqual(ps, pt) {
		// 	t.Errorf("expected (%v) got (%v)", ps, pt)
		// }

	}

}

var recCount int = 1000
var file_sourceDB string = "./data/source_db.csv"

func Test_populateSourceDB(t *testing.T) {
	populateSourceDB(recCount, file_sourceDB)
}

// // a successful case
// func TestShouldUpdateStats(t *testing.T) {
// 	mock, err := pgxmock.NewPool()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer mock.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").
// 		WillReturnResult(pgxmock.NewResult("UPDATE", 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").
// 		WithArgs(2, 3).
// 		WillReturnResult(pgxmock.NewResult("INSERT", 1))
// 	mock.ExpectCommit()

// 	// now we execute our method
// 	if err = recordStats(mock.Conn(), 2, 3); err != nil {
// 		t.Errorf("error was not expected while updating: %s", err)
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// // a failing test case
// func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
// 	mock, err := pgxmock.NewPool()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer mock.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").
// 		WillReturnResult(pgxmock.NewResult("UPDATE", 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").
// 		WithArgs(2, 3).
// 		WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectRollback()

// 	// now we execute our method
// 	if err = recordStats(mock.Conn(), 2, 3); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
