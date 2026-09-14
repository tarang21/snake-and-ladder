package main

type Ladder struct {
	Start *Square
	End   *Square
}

func NewLadder(startSquare Square, endSquare Square) *Ladder {
	return &Ladder{&startSquare, &endSquare}
}

func (l *Ladder) start() *Square {
	return l.Start
}

func (l *Ladder) end() *Square {
	return l.End
}
