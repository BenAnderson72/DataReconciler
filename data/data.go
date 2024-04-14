package Data

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/jaswdr/faker/v2"
)

// Payment represents a single payment
type PaymentType struct {
	Time             time.Time `json:"timestamp" diff:"-"`
	Sender_Account   string    `json:"sender_account"`
	Receiver_Account string    `json:"receiver_account"`
	TransactionID    string    `json:"transaction_id" bson:"transaction_id"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Reference        string    `json:"description"`
}

// LoadPayment loads a single payment from a slice of strings
func LoadPayment(values []string) PaymentType {

	amt, _ := strconv.ParseFloat(values[4], 64)
	time, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", values[0])

	return PaymentType{
		Time:             time,
		Sender_Account:   values[1],
		Receiver_Account: values[2],
		TransactionID:    values[3],
		Amount:           amt,
		Currency:         values[5],
		Reference:        values[6],
	}

}

// Reduce Floats to fixed precision
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// GenPayment generates a single payment
func GenPayment() PaymentType {

	// Randomise
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fake := faker.NewWithSeed(r)

	// get last (t0) and next (t1) midnight timestamps
	loc, _ := time.LoadLocation("Europe/London")

	year, month, day := time.Now().In(loc).Date()
	t0 := time.Date(year, month, day, 0, 0, 0, 0, loc)
	t1 := t0.AddDate(0, 0, 1)

	pmnt := PaymentType{
		Time:             fake.Time().TimeBetween(t0, t1),
		Sender_Account:   fake.Payment().CreditCardNumber(),
		Receiver_Account: fake.Payment().CreditCardNumber(),
		TransactionID:    fake.UUID().V4(),
		Amount:           roundFloat(fake.Float64(99, 0, 666), 2),
		// Currency:      fake.Currency().Currency(),
		Currency:  fake.RandomStringElement([]string{"GBP", "USD"}),
		Reference: fmt.Sprintf("REF %d", fake.Int16Between(0, 9999)),
	}

	// fake.Struct().Fill(&pmnt)

	return pmnt
}
