package main

import "fmt"

func main() {
	players := configureGame()
	StartGame(players)
}


func configureGame() []Player{

	// number Of Players
	var numberOfPlayers int

	for {
		fmt.Println("Enter number of players (2-4):")
		fmt.Scanln(&numberOfPlayers)

		if numberOfPlayers >= 2 && numberOfPlayers <= 4 {
			break
		}

		fmt.Println("Invalid number of players. Please enter a number between 2 and 4.")
	}
	
	var players []Player

	for i := 1; i <= numberOfPlayers; i++ {
		fmt.Printf("Player %d, enter your name: ", i)
		var playerName string
		fmt.Scanln(&playerName)
		fmt.Printf("Player '%s' registered!\n", playerName)

		player := Player{
			playerID: i,
			playerName: playerName,
			playerCards: nil, // Initialize with no cards
		}
		players = append(players, player)
	}

	return players
}