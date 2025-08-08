package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	numberOfPlayers int
	players []Player
	cardsPlayed []Card
	cardsDeck []Card
	cardOnTable Card
	playerOnTurn int
}

func StartGame(players []Player){

	var game Game

	fmt.Println("Initializing Game...")

	game.players = players
	game.numberOfPlayers = len(players)
	game.playerOnTurn = 0
	game.cardsPlayed = make([]Card, 0)

	createCardsDeck(&game)

	createPlayersHands(&game.players, &game)

	getFirstCardOnTable(&game)


	fmt.Println("Game Initialized")

	play(&game)

}


func createCardsDeck(game *Game) {
	fmt.Println("Creating Cards Deck...")

	var cards []Card

	// Create cards
	for i := 0; i < NUMBER_OF_COLORS; i++ {
		for j := -1; j < 10; j++ { // Assuming each color has 10 cards
			card := Card{value: j, color: i}
			cards = append(cards, card)
		}
	}
	
	game.cardsDeck = make([]Card, len(cards))
	index := 0
	// Shuffle the cards
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		game.cardsDeck[index] = cards[j]
		index++
		cards = append(cards[:j], cards[j+1:]...) // Remove the card from the deck
	}

	fmt.Println("Cards Deck Created")
}

func createPlayersHands(players *[]Player, game *Game) {
	fmt.Println("Creating Players Hands...")

	// Initialize players hands
	for i := 0; i < len(*players); i++ {
		for j := 0; j < 7; j++ { // Assuming each player starts with 7 cards
			if len(game.cardsDeck) > 0 {
				card := game.cardsDeck[0]
				game.cardsDeck = game.cardsDeck[1:] // Remove the card from the deck

				(*players)[i].AddCard(card) // Add card to player's hand
			}
		}
	}

	fmt.Println("Players Hands Created")
}


func getFirstCardOnTable(game *Game) {
	game.cardOnTable = game.cardsDeck[0] // Set the first card on the table
	game.cardsDeck = game.cardsDeck[1:] // Remove the card from the deck
}	



func play (game *Game){
	// while true
	for {
		fmt.Println("----------------------------------------------------------------------------------------")
		fmt.Println("----------------------------------------------------------------------------------------")

		if game.players[game.playerOnTurn].Play(game){
			return
		}
	}
	
}



