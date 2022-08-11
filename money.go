package money

type Money struct {
	paise int
}

func NewMoney(paise int) Money {
	if paise < 0 {
		panic("paise should be positive")
	}
	return Money{paise: paise}
}

func (money *Money) AddMoney(paise int) {
	if paise < 0 {
		panic("paise should be positive")
	}
	money.paise += paise
}
