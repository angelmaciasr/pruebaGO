package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	playerID int
	playerName string
	playerCards []Card
}


func (p *Player) AddCard(card Card) {
	p.playerCards = append(p.playerCards, card)
	fmt.Printf("Player %s received card: %s\n", p.playerName, card.PrintCard())
}

func (p *Player) Play(game *Game) bool {
	fmt.Printf("Player %s's turn\n", p.playerName)
	fmt.Printf("Player %s's cards: ", p.playerName)

	for _, card := range p.playerCards {
		fmt.Printf("| %s |", card.PrintCard())
	}
	fmt.Printf("\nCard on table: %s\n", game.cardOnTable.PrintCard())

	

	// Simulate player playing a card
	return p.playTurn(game)
}


func (p *Player) playTurn(game *Game) bool{
	cards := p.getValidCards(game)

	if len(cards) == 0 {
		fmt.Printf("Player %s has no valid cards to play and steals a card.\n", p.playerName)

		p.stealCard(game)

		if isValidCard(p.playerCards[len(p.playerCards)-1], game.cardOnTable) {
			// Play the card
			p.playCard(game, len(p.playerCards)-1)
		}

		return false //if Player had to steal a card, didnt finished
	} else {

		var nCardPlayed = cards[rand.Intn(len(cards))]
		var finished = p.playCard(game, nCardPlayed)

		return finished
	}
}


func (p *Player) playCard(game *Game, cardIndex int) bool {
	card := p.playerCards[cardIndex]

	// Remove the card from player's hand
	if cardIndex == len(p.playerCards)-1  {
		p.playerCards = p.playerCards[:len(p.playerCards)-1] // Remove the last card
	}else{
		p.playerCards = append(p.playerCards[:cardIndex], p.playerCards[cardIndex+1:]...)
	}

	// Execute card actions
	return card.Play(game)
}



func (p *Player) getValidCards(game *Game) []int {
	validCards := make([]int, 0)

	for i := range p.playerCards {
		if isValidCard(p.playerCards[i], game.cardOnTable) {
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
	p.playerCards = append(p.playerCards, card)
	fmt.Printf("Player %s steals a card: %s\n", p.playerName, card.PrintCard())

	if len(game.cardsDeck) == 0 {
		game.cardsDeck = game.cardsPlayed
		game.cardsPlayed = make([]Card, 0) // Reset the played cards
	}
}


func (p *Player) isGameOver() bool {
	return len(p.playerCards) == 0
}