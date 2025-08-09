package game

import (
	"fmt"
	"math/rand"
	"pruebaGO/printer"
)

type Player struct {
	PlayerID   int
	PlayerName string
	PlayerCards []Card
}


func (p *Player) AddCard(card Card) {
	p.PlayerCards = append(p.PlayerCards, card)
}

func (p *Player) Play(game *Game) bool {
	// Simulate player playing a card
	return p.playTurn(game)
}


func (p *Player) playTurn(game *Game) bool{
	cards := p.getValidCards(game)

	if len(cards) == 0 {
		p.stealCard(game)

		if isValidCard(p.PlayerCards[len(p.PlayerCards)-1], game.cardOnTable) {
			// Play the card
			p.playCard(game, len(p.PlayerCards)-1)
		}

		return false //if Player had to steal a card, didnt finished
	} else {

		var nCardPlayed = cards[rand.Intn(len(cards))]
		var finished = p.playCard(game, nCardPlayed)

		return finished
	}
}


func (p *Player) playCard(game *Game, cardIndex int) bool {
	card := p.PlayerCards[cardIndex]

	// Remove the card from player's hand
	if cardIndex == len(p.PlayerCards)-1  {
		p.PlayerCards = p.PlayerCards[:len(p.PlayerCards)-1] // Remove the last card
	}else{
		p.PlayerCards = append(p.PlayerCards[:cardIndex], p.PlayerCards[cardIndex+1:]...)
	}

	// Execute card actions
	return card.Play(game)
}



func (p *Player) getValidCards(game *Game) []int {
	validCards := make([]int, 0)

	for i := range p.PlayerCards {
		if isValidCard(p.PlayerCards[i], game.cardOnTable) {
			validCards = append(validCards, i)
		}
	}

	return validCards
}

func isValidCard(card Card, topCard Card) bool {
	return card.color == topCard.color || card.value == topCard.value
}

func (p *Player) stealCard(game *Game) {
	card := game.cardsDeck[0]
	game.cardsDeck = game.cardsDeck[1:] // Remove the card from the deck
	p.PlayerCards = append(p.PlayerCards, card)
	
	printer.PrintStealCard(p.PlayerName, card.ToString())

	if len(game.cardsDeck) == 0 {
		game.cardsDeck = game.cardsPlayed
		game.cardsPlayed = make([]Card, 0) // Reset the played cards
	}
}


func (p *Player) isGameOver() bool {
	return len(p.PlayerCards) == 0
}


func (p *Player) PrintPlayersCards() string {
	var res string
	for _, card := range p.PlayerCards {
		res = fmt.Sprintf("| %s |", card.ToString())
	}

	return res
}