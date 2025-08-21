package pkg

import "fmt"

type PaymentType interface {
	Pay(amount float64)
}

type Zarinpal struct{}

func (z *Zarinpal) Pay(amount float64) {
	fmt.Println("Paid", amount, "with zarinpal")
}

type Melli struct{}

func (m *Melli) Pay(amount float64) {
	fmt.Println("Paid", amount, "with Melli")
}

// Strategy Pattern
type PaymentContext struct {
	payment PaymentType
}

func (c PaymentContext) setPayment(pt PaymentType) {
	c.payment = pt
}

func (c PaymentContext) Checkout(amount float64) {
	c.payment.Pay(amount)
}

func example() {
	ctx := PaymentContext{}
	ctx.setPayment(&Zarinpal{})
	ctx.Checkout(100_240)
}
