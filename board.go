package main

type Board struct {
	Squares []*Square
	Ladders []*Ladder
	Snakes  []*Snake
	Jumps   []Jumps
}

func NewBoard(Ladders []*Ladder, Snakes []*Snake) *Board {
	var Squares []*Square
	for i := 1; i <= 100; i++ {
		Squares = append(Squares, &Square{i})
	}
	Jumps := allJumps(Ladders, Snakes)
	return &Board{Squares, Ladders, Snakes, Jumps}
}

func allJumps(ladders []*Ladder, snakes []*Snake) []Jumps {
	var Jumps []Jumps
	for _, ladder := range ladders {
		Jumps = append(Jumps, ladder)
	}
	for _, snake := range snakes {
		Jumps = append(Jumps, snake)
	}
	return Jumps
}
