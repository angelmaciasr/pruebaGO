package game

import (
	"fmt"
	"pruebaGO/printer"
)

// Card struct with a color field of type int
type Card struct {
	value int
	color int
}

const (
	NUMBER_OF_COLORS = 4
)

func getColorName(c int) string {

	if(c < 0 || c >= NUMBER_OF_COLORS) {
		return fmt.Sprintf("Undefined Color %d", c)
	}

	colors := []string{
		"Red", // 0
		"Blue", // 1
		"Green", // 2
		"Yellow", // 3
	}

	return colors[c]
}

func (c Card) Play(game *Game) bool{

	// Put card on the table
	game.cardOnTable = c
	game.cardsPlayed = append(game.cardsPlayed, game.cardOnTable)

	printer.PrintCardPlay(game.players[game.playerOnTurn].PlayerName, c.ToString())

	// Check game over
	if game.players[game.playerOnTurn].isGameOver() {
		printer.PrintGameOver(game.players[game.playerOnTurn].PlayerName)
		return true
	}

	// Next player
	switch c.value {
		case -1: {game.playerOnTurn = (game.playerOnTurn +2) % game.numberOfPlayers}// Jump
		default: {// Move to next player
			game.playerOnTurn = (game.playerOnTurn + 1) % game.numberOfPlayers
		}
	}

	return false
}


func (c Card) ToString() string {
	var value string
	if c.value == -1 {
		value = "Jump"
	} else {
		value = fmt.Sprintf("%d", c.value)
	}
	return fmt.Sprintf("%s %s", getColorName(c.color), value)
}
