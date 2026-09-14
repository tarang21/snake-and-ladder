package main

import "math/rand/v2"

type Dice struct {
	Sides int
}

func (d *Dice) RollDice() int {
	return 1 + rand.IntN(6)
}

func NewDice() *Dice {
	return &Dice{Sides: 6}
}
