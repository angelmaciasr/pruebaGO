package printer

import (
	"fmt"
)


func AskPlayers() []string {
	var numberOfPlayers int
	for {
		fmt.Println("Enter number of players (2-4):")
		fmt.Scanln(&numberOfPlayers)

		if numberOfPlayers >= 2 && numberOfPlayers <= 4 {
			break
		}

		fmt.Println("Invalid number of players. Please enter a number between 2 and 4.")
	}

	var players []string

	for i := 0; i < numberOfPlayers; i++ {
		fmt.Printf("Player %d, enter your name: ", i+1)

		var playerName string
		fmt.Scanln(&playerName)

		fmt.Printf("Player '%s' registered!\n", playerName)

		players = append(players, playerName)
	}
	return players
}


func PrintTurnInit(playerName string, playersCards string, cardOnTable string){
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------------------------")

	fmt.Printf("Player %s's turn\n", playerName)
	fmt.Printf("Player %s's cards: ", playersCards)

	fmt.Printf("\nCard on table: %s\n", cardOnTable)
}

func PrintCardPlay(playerName string, cardString string){
	fmt.Printf("Player %s plays: %s\n", playerName, cardString)
}

func PrintStealCard(playerName string, cardString string){
	fmt.Printf("Player %s has no valid cards to play and steals a card.\n", playerName)
	fmt.Printf("Player %s steals a card: %s\n", playerName, cardString)
}

func PrintGameOver(playerName string){
	fmt.Printf("Player %s wins!\n", playerName)
}