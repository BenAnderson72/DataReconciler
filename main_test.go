package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_populateTargetDB(t *testing.T) {

	recCount := 100

	collection := populateTargetDB(recCount)

	nr, _ := collection.Count(bson.M{})
	// fmt.Printf("Found %d items", nr)

	if nr != uint64(recCount) {
		t.Errorf("Found %d items", nr)
	}

}

// go test -run ^Test_populateSourceDB$ github.com/BenAnderson72/DataReconciler
func Test_populateSourceDB(t *testing.T) {
	populateSourceDB(1000, "./data/source_db.csv")
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
