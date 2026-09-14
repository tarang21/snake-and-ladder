package main

type Coin struct {
	CurrentPosition *Square
	Color           string
}

func NewCoin(color string) *Coin {
	return &Coin{&Square{0}, color}
}
