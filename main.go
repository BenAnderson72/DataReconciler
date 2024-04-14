package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	Data "github.com/BenAnderson72/DataReconciler/data"

	"github.com/mjarkk/mongomock"
)

// populateTargetDB0 populates the target database with the given number of records !DEPRECATED!
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

// populateSourceDB populates the source DB (csv file!) with the given number of records
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

// populateTargetDB copies the contents of the target DB with the contents of the source DB
func populateTargetDB(file_sourceDB string) mongomock.Collection {

	records := getSourceData(file_sourceDB)

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
