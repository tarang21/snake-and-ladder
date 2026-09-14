package main

func main() {
	ladder1 := NewLadder(*NewSquare(4), *NewSquare(14))
	ladder2 := NewLadder(*NewSquare(9), *NewSquare(31))
	ladder3 := NewLadder(*NewSquare(21), *NewSquare(42))
	ladder4 := NewLadder(*NewSquare(28), *NewSquare(84))
	ladder5 := NewLadder(*NewSquare(51), *NewSquare(67))
	ladder6 := NewLadder(*NewSquare(71), *NewSquare(91))

	ladders := []*Ladder{ladder1, ladder2, ladder3, ladder4, ladder5, ladder6}

	snake1 := NewSnake(*NewSquare(16), *NewSquare(6))
	snake2 := NewSnake(*NewSquare(47), *NewSquare(26))
	snake3 := NewSnake(*NewSquare(62), *NewSquare(19))
	snake4 := NewSnake(*NewSquare(87), *NewSquare(24))
	snake5 := NewSnake(*NewSquare(93), *NewSquare(73))
	snake6 := NewSnake(*NewSquare(98), *NewSquare(78))

	snakes := []*Snake{snake1, snake2, snake3, snake4, snake5, snake6}

	board := NewBoard(ladders, snakes)

	red := NewCoin("Red")
	blue := NewCoin("Blue")
	green := NewCoin("Green")

	coins := []*Coin{red, blue, green}

	game := *NewGame(*board, coins)
	game.Play()
}
