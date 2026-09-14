package main

type Snake struct {
	Start *Square
	End   *Square
}

func NewSnake(startSquare Square, endSquare Square) *Snake {
	return &Snake{&startSquare, &endSquare}
}

func (s *Snake) start() *Square {
	return s.Start
}

func (s *Snake) end() *Square {
	return s.End
}
