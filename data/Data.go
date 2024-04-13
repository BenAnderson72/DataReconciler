package Data

import (
	"math"
	"time"

	"github.com/jaswdr/faker/v2"
)

// {"transaction_id":"TXN000001","amount":361.5,"currency":"GBP","sender_account":"650BXA7BITIAXD6OWN0FUQ","receiver_account":"LXKVL3PY29NNTMO1B8HG0O","transaction_date":"2023-11-10","payment_reference":"Invoice 00056"}

type paymentType struct {
	Time             time.Time `json:"timestamp"`
	Sender_Account   string    `json:"sender_account"`
	Receiver_Account string    `json:"receiver_account"`
	TransactionID    string    `json:"transaction_id"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Reference        string    `json:"description"`
}

// Deal with Floats not fixing dps
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GenPaymentSource() paymentType {

	fake := faker.New()

	// get last (t0) and next (t1) midnight timestamps
	loc, _ := time.LoadLocation("Europe/London")
	year, month, day := time.Now().In(loc).Date()
	t0 := time.Date(year, month, day, 0, 0, 0, 0, loc)
	t1 := t0.AddDate(0, 0, 1)

	pmnt := paymentType{
		Time:             fake.Time().TimeBetween(t0, t1),
		Sender_Account:   fake.Payment().CreditCardNumber(),
		Receiver_Account: fake.Payment().CreditCardNumber(),
		TransactionID:    fake.UUID().V4(),
		Amount:           roundFloat(fake.Float64(2, 0, 666), 2),
		// Currency:      fake.Currency().Currency(),
		Currency:  fake.RandomStringElement([]string{"GBP", "USD"}),
		Reference: fake.Lorem().Text(20),
	}

	// fake.Struct().Fill(&pmnt)

	return pmnt
}
