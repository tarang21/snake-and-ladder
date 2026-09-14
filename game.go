package main

import "fmt"

type Game struct {
	Board *Board
	Coins []*Coin
	Dice  *Dice
}

func NewGame(board Board, coins []*Coin) *Game {
	return &Game{&board, coins, NewDice()}
}

func (g *Game) Play() {
	for {
		gameOver := false
		for _, coin := range g.Coins {
			fmt.Println(coin.Color, "playing.")
			diceNumber := g.Dice.RollDice()
			coin.CurrentPosition.SquareNumber += diceNumber
			if coin.CurrentPosition.SquareNumber > 100 {
				coin.CurrentPosition.SquareNumber -= diceNumber
				fmt.Println(coin.Color, "need exact 100, stays at", coin.CurrentPosition.SquareNumber)
				continue
			}
			g.applyJumps(coin)
			fmt.Println("Color", coin.Color, "final position", coin.CurrentPosition.SquareNumber)
			if coin.CurrentPosition.SquareNumber == 100 {
				fmt.Println("Color", coin.Color, "Won!")
				gameOver = true
				break
			}
		}
		if gameOver == true {
			break
		}
	}
}

func (g *Game) applyJumps(coin *Coin) {
	jumps := g.Board.Jumps
	for {
		moved := false
		for _, jump := range jumps {
			if coin.CurrentPosition.SquareNumber == jump.start().SquareNumber {
				coin.CurrentPosition.SquareNumber = jump.end().SquareNumber
				fmt.Println(coin.Color, "jumped to", coin.CurrentPosition.SquareNumber)
				moved = true
				break // one jump, then check the new square from the top
			}
		}
		if !moved {
			return
		}
	}
}
