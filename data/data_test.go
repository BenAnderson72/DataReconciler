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

func Test_LoadPayment(t *testing.T) {

	values := []string{"2024-04-14 07:20:17.044182591 +0100 BST", "9277785356448308", "6517378067341210", "10362433-0fb8-475a-9587-36983e3f5cc6", "404.59", "USD", "REF 5135"}

	p1 := LoadPayment(values)

	// t.Logf("Got (%v)", p1)

	if p1.Time.String() != values[0] {
		t.Errorf("Fail p1.Time %v", p1)
	}

	if p1.Sender_Account != values[1] {
		t.Errorf("Fail p1.Sender_Account %v", p1)
	}

	if p1.Receiver_Account != values[2] {
		t.Errorf("Fail p1.Receiver_Account %v", p1)
	}

	if p1.TransactionID != values[3] {
		t.Errorf("Fail p1.TransactionID %v", p1)
	}

	if fmt.Sprintf("%.2f", p1.Amount) != values[4] {
		t.Errorf("Fail p1.Amount %v", p1)
	}

	if p1.Currency != values[5] {
		t.Errorf("Fail p1.Currency %v", p1)
	}

	if p1.Reference != values[6] {
		t.Errorf("Fail p1.Reference %v", p1)
	}

}
