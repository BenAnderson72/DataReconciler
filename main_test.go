package main

import (
	"log"
	"math/rand/v2"
	"slices"
	"testing"

	Data "github.com/BenAnderson72/DataReconciler/data"
	"github.com/mjarkk/mongomock"
	"go.mongodb.org/mongo-driver/bson"
)

var recCount int = 1000
var file_sourceDB string = "./data/source_db.csv"

// func Test_populateTargetDB0(t *testing.T) {

// 	collection := populateTargetDB0(recCount)

// 	nr, _ := collection.Count(bson.M{})
// 	// fmt.Printf("Found %d items", nr)

// 	if nr != uint64(recCount) {
// 		t.Errorf("Found %d items", nr)
// 	}

// }

// corruptSourceDB corrupts the source DB (csv file!) by recCount # of records to give us something to reconcile
// it returns the ids of the corrupt records
func corruptTargetDB(file_sourceDB string, recCount int, collection *mongomock.Collection) []string {

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

	return txs

}

// This test method populates the target mongoDB from the source DB (source_db.csv)
func Test_populateTargetDB(t *testing.T) {

	collection := populateTargetDB(file_sourceDB)

	nr, _ := collection.Count(bson.M{})

	// Check record counts
	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}

// This test method populates the target mongoDB from the source DB (source_db.csv)
// corrupts the source DB (csv file!) by recCount # of records
// reconciles the records by count and data (using data.diff) and reports diffs to a mock mongo store
// The abiding test is to ensure that the diffs are expected
func Test_reconcile(t *testing.T) {

	targetCollection := populateTargetDB(file_sourceDB)

	// Corrupt 10 records on target DB
	corruptedTxs := corruptTargetDB(file_sourceDB, 10, &targetCollection)

	sourceRecords := getSourceData(file_sourceDB)

	reconCol := reconcileRecords(sourceRecords, targetCollection)

	nr, _ := reconCol.Count(bson.M{})

	// Reconcile record counts
	if nr != uint64(len(corruptedTxs)) {
		t.Errorf("Found %d items", nr)
	}

	// Check all out txids are found
	for _, txid := range corruptedTxs {
		pr := PaymentRecon{}

		// Check Transaction IDs exist
		err := reconCol.FindFirst(&pr, bson.M{"transaction_id": txid})

		if err != nil {
			t.Fatal(err)
		}

	}

}

// This test method populates the target mongoDB from the source DB (source_db.csv)
// func Test_reconcile0(t *testing.T) {

// 	targetCollection := populateTargetDB(file_sourceDB)

// 	nr, _ := targetCollection.Count(bson.M{})

// 	// Reconcile record counts
// 	if nr != uint64(recCount) {
// 		t.Errorf("Found %d items", nr)
// 	}

// 	// Corrupt 10 records on target DB
// 	corruptTargetDB(file_sourceDB, 10, &targetCollection)

// 	sourceRecords := getSourceData(file_sourceDB)

// 	for _, rec := range sourceRecords {
// 		pt := Data.PaymentType{}

// 		// Check Transaction IDs exist
// 		err := targetCollection.FindFirst(&pt, bson.M{"transaction_id": rec[3]})

// 		// if strings.HasSuffix(pt.Reference, "CORRUPT") {
// 		// 	fmt.Println("here")
// 		// }

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		ps := Data.LoadPayment(rec)

// 		changelog, _ := diff.Diff(ps, pt)

// 		if len(changelog) != 0 {
// 			t.Errorf("%v", changelog)
// 		}

// 	}

// }

// This test method populates the source DB (source_db.csv) with fake data
func Test_populateSourceDB(t *testing.T) {
	populateSourceDB(recCount, file_sourceDB)
}
