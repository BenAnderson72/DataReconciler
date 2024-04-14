package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_populateTargetDB0(t *testing.T) {

	recCount := 100

	collection := populateTargetDB0(recCount)

	nr, _ := collection.Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}

func Test_populateTargetDB(t *testing.T) {

	collection := populateTargetDB(file_sourceDB)

	nr, _ := collection.Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	if nr != uint64(1000) {
		t.Errorf("Found %d items", nr)
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
