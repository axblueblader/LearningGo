package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// ExitCode : code for exit loop
	ExitCode int = -1
	// HumanIndex : index for human player
	HumanIndex = 1
	// BotIndex : index for bot player
	BotIndex = 2
	// WinCode : code for win game
	WinCode = 1
	// LoseCode : code for lose game
	LoseCode = 2
	// DrawCode : code for draw gameDrawCode, LoseCode},
	DrawCode = 0
	// RockCode : represents rock choice
	RockCode = 0
	// PaperCode : represents paper choice
	PaperCode = 1
	// ScissorCode : represents scissor choice
	ScissorCode = 2
)

var weaponName = [...]string{"Rock", "Paper", "Scissor"}

// PlayerModel contains name and choice of player
type PlayerModel struct {
	choice int
	name   string
}

// RockPaperScissorGame : is the controller, recieves data from input, returns data for outputting
type RockPaperScissorGame struct {
	theHuman PlayerModel
	theBot   PlayerModel
	winner   int
}

func main() {
	var newGame RockPaperScissorGame
	fmt.Println("Welcome to program")
	initializePlayers(&newGame)
	mainGameLoop(&newGame)
	fmt.Println("Thanks for playing")
}

func initializePlayers(theGame *RockPaperScissorGame) {
	var playerName string
	for {
		fmt.Println("Enter your name: ")
		_, err := fmt.Scanf("%s", &playerName)
		if err != nil {
			fmt.Println("An error occured, please re-enter:")
			continue
		}
		break
	}
	theGame.theHuman = PlayerModel{name: playerName, choice: 0}
	theGame.theBot = PlayerModel{name: "BOT", choice: 0}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func getInputFromUser() (int, error) {
	var playerChoice int
	fmt.Printf("Choose [%d]-rock,[%d]-paper,[%d]-scissor,[%d]-exit: ",
		RockCode, PaperCode, ScissorCode, ExitCode)
	_, err := fmt.Scanf("%d", &playerChoice)
	if err != nil || invalidPlayerChoice(playerChoice) {
		return 0, &errorString{s: "Input error"}
	}
	return playerChoice, nil
}

func mainGameLoop(theGame *RockPaperScissorGame) {
	for {
		playerChoice, err := getInputFromUser()
		if err != nil {
			fmt.Printf("%s, please re-enter\n", err.Error())
			continue
		}
		if playerChoice == ExitCode {
			break
		} else {
			theGame.Run(playerChoice)
			showChoices(*theGame)
			showResults(*theGame)
		}
	}
}

var choices = [...]int{0, 1, 2, -1}

func invalidPlayerChoice(playerChoice int) bool {
	var validity = true
	for _, choice := range choices {
		if choice == playerChoice {
			validity = false
		}
	}
	return validity
}

func (player *PlayerModel) setChoice(choice int) {
	player.choice = choice
}

func generateRandomChoice() int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(seed)
	return randomer.Intn(2)
}

// Run : set choices and set winner
func (theGame *RockPaperScissorGame) Run(playerChoice int) {
	theGame.theHuman.setChoice(playerChoice)
	theGame.theBot.setChoice(generateRandomChoice())
	theGame.setWinner(getHumanResult(theGame.theHuman, theGame.theBot))
}

// RoundResult : result of round based on player's choices
// Usage:
// RoundResult[playerA.choice][playerB.choice]
var RoundResult = [][]int{
	[]int{DrawCode, LoseCode, WinCode},
	[]int{WinCode, DrawCode, LoseCode},
	[]int{LoseCode, WinCode, DrawCode},
}

func getHumanResult(humanPlayer, botPlayer PlayerModel) int {
	return RoundResult[humanPlayer.choice][botPlayer.choice]
}

func (theGame *RockPaperScissorGame) setWinner(roundResult int) {
	switch roundResult {
	case WinCode:
		theGame.winner = HumanIndex
	case LoseCode:
		theGame.winner = BotIndex
	case DrawCode:
		theGame.winner = DrawCode
	}
}

func showChoices(theGame RockPaperScissorGame) {
	fmt.Printf("%s's choice: %s | Bot's choice: %s\n",
		theGame.theHuman.name, // need to call Get method here
		weaponName[theGame.theHuman.choice],
		weaponName[theGame.theBot.choice])
}

func showResults(theGame RockPaperScissorGame) {
	switch theGame.winner {
	case HumanIndex:
		fmt.Printf("%s won the game \n", theGame.theHuman.name)
	case BotIndex:
		fmt.Println("Bot won the game")
	case DrawCode:
		fmt.Println("The game was draw")
	}
}
