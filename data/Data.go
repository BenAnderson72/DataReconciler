package Data

import (
	"fmt"
	"time"

	"github.com/jaswdr/faker/v2"
)

type paymentType struct {
	Time          time.Time `json:"timestamp"`
	Payer         string    `json:"payer"`
	Payee         string    `json:"payee"`
	TransactionID string    `json:"id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Description   string    `json:"description"`
}

func GenPaymentSource() {

	fake := faker.New()

	loc, _ := time.LoadLocation("Europe/London")

	year, month, day := time.Now().In(loc).Date()
	t0 := time.Date(year, month, day, 0, 0, 0, 0, loc)
	t1 := t0.AddDate(0, 0, 1)

	pmnt := paymentType{
		Time:          fake.Time().TimeBetween(t0, t1),
		Payer:         fake.Person().Name(),
		Payee:         fake.Person().Name(),
		TransactionID: fake.Hash().MD5(),
		Amount:        fake.Float(2, -6666666, 6666666),
		Currency:      fake.Currency().Currency(),
		Description:   fake.Lorem().Text(20),
	}

	// fake.Struct().Fill(&pmnt)

	fmt.Printf("%+v", pmnt)

	// fake.Payment().CreditCardNumber()

	// fake.Person().Name()
	// // Lucy Cechtelar

	// fake.Address().Address()
	// // 426 Jordy Lodge

	// fake.Lorem().Text(100)

	// return ""
}
