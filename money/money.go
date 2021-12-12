package money

import "errors"

type Money struct {
	amount   float32
	currency string
}

func New(amount float32, currency string) (*Money, error) {
	if len(currency) == 0 {
		return nil, errors.New("null exception")
	}
	money := &Money{
		amount:   amount,
		currency: currency,
	}
	return money, nil
}

func (m *Money) Add(arg Money) (*Money, error) {
	if m.currency != arg.currency {
		return nil, errors.New("currency is not match")
	}
	money := &Money{
		amount:   m.amount + arg.amount,
		currency: m.currency,
	}
	return money, nil
}
