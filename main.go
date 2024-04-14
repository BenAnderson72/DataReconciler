package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	Data "github.com/BenAnderson72/DataReconciler/data"
	"github.com/r3labs/diff/v3"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mjarkk/mongomock"
)

// populateSourceDB populates the source DB (csv file!) with the given number of records. It will overwrite any existing data
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

		w.Write([]string{fmt.Sprintf("%d", n), pmnt.Time.String(), pmnt.Sender_Account, pmnt.Receiver_Account, pmnt.TransactionID, fmt.Sprintf("%.2f", pmnt.Amount), pmnt.Currency, pmnt.Reference})

		n++
	}

	w.Flush()

}

// getSourceData reads the source data from a csv file and returns it as a slice of strings
func getSourceData(file_sourceDB string) [][]string {
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

	return records
}

// struct to hold diff data
type PaymentRecon struct {
	TransactionID string         `json:"transaction_id" bson:"transaction_id"`
	Changelog     diff.Changelog `json:"changelog"`
}

// reconcileRecords reconciles the source records with the target collection, it rights the diff data in the PaymentRecon struct to the recon collection & returns it
func reconcileRecords(sourceRecords [][]string, targetCollection mongomock.Collection) mongomock.Collection {

	db := mongomock.NewDB()
	reconCollection := db.Collection("reconciliations")

	for _, rec := range sourceRecords {
		pt := Data.PaymentType{}

		// Check Transaction IDs exist and point target payment to it
		err := targetCollection.FindFirst(&pt, bson.M{"transaction_id": rec[3]})

		if err != nil {
			log.Fatal(err)
		}

		// Load source payment from csv record
		ps := Data.LoadPayment(rec)

		// compare the pair of payments
		changelog, _ := diff.Diff(ps, pt)

		if len(changelog) != 0 {
			pr := PaymentRecon{TransactionID: ps.TransactionID}
			pr.Changelog = changelog

			err := reconCollection.Insert(pr)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

	return *reconCollection

}

// populateTargetDB copies the contents of the target DB with the contents of the source DB
func populateTargetDB(file_sourceDB string) mongomock.Collection {

	records := getSourceData(file_sourceDB)

	db := mongomock.NewDB()
	txCollection := db.Collection("transactions")

	for _, rec := range records {
		pmnt := Data.LoadPayment(rec)
		err := txCollection.Insert(pmnt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return *txCollection
}
