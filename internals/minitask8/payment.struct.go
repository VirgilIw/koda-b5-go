package minitask8

type Checkout interface {
	Payment()
}

type PaymentBank struct{}
type PaymentOnline struct{}

func (p PaymentBank) Payment() string {
	return "bank"
}

func (p PaymentOnline) Payment() string {
	return "online"
}
