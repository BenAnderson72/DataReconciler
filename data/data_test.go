package Data

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/r3labs/diff/v3"
)

func Test_GenSourceData(t *testing.T) {

	pmnt := GenPayment()

	fmt.Printf("%+v", pmnt)

	// s := GetSpot("Rest Bay (Porthcawl)")

	// if s.Name != "Rest Bay (Porthcawl)" {
	// 	t.Errorf("invalid name : %s", s.Name)
	// }

}

func Test_CompareDE1(t *testing.T) {

	pmnt1 := GenPayment()

	pmnt2 := GenPayment()

	if !reflect.DeepEqual(pmnt1, pmnt2) {
		t.Errorf("expected (%v) got (%v)", pmnt1, pmnt2)
	}

}

func Test_CompareDE2(t *testing.T) {

	pmnt1 := paymentType{
		Sender_Account:   "1234",
		Receiver_Account: "5678",
		TransactionID:    "ABC123",
		Amount:           410.20,
		// Currency:      fake.Currency().Currency(),
		Currency:  "GBP",
		Reference: "REF1",
	}

	pmnt2 := pmnt1

	// pmnt2.Sender_Account="12345"

	if !reflect.DeepEqual(pmnt1, pmnt2) {
		t.Errorf("expected (%v) got (%v)", pmnt1, pmnt2)
	}

}

func Test_CompareDiff1(t *testing.T) {

	pmnt1 := paymentType{
		Sender_Account:   "1234",
		Receiver_Account: "5678",
		TransactionID:    "ABC123",
		Amount:           410.20,
		// Currency:      fake.Currency().Currency(),
		Currency:  "GBP",
		Reference: "REF1",
	}

	pmnt2 := pmnt1

	pmnt2.Sender_Account = "12345"

	changelog, _ := diff.Diff(pmnt1, pmnt2)

	fmt.Printf("%d", len(changelog))
}
