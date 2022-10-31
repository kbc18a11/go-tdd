package src

type Dollar struct {
	Amount int
}

func (dollar *Dollar) New(amount int) {
	dollar.Amount = amount
}

func (dollar *Dollar) Times(multiplier int) Dollar {
	return Dollar{dollar.Amount * multiplier}
}
