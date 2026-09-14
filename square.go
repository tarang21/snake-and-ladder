package main

type Square struct {
	SquareNumber int
}

func NewSquare(squareNumber int) *Square {
	return &Square{squareNumber}
}
