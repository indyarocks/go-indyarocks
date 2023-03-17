package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func (g *Game) Rounds() {
	//	use select to process input in channels
	//	  Print information to screen
	//	Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "window") {
		//windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		//	linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("======================")
	fmt.Println("Let the game begin. Best of 3 wins!")
	fmt.Println("===================================")
}

func (g *Game) PlayRound() bool {
	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("-----")

	fmt.Print("Please enter rock, paper or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	//<-g.DisplayChan

	switch computerValue {
	case ROCK:
		fmt.Println("Computer choose: ROCK")
		break
	case PAPER:
		fmt.Println("Computer choose: PAPER")
		break
	case SCISSORS:
		fmt.Println("Computer choose: SCISSORS")
		break
	default:
	}

	fmt.Println()

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice"
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
}

func (g *Game) PrintSummary() {
	fmt.Println("Game over!")
	fmt.Println("Here is the result:")
	if g.Round.PlayerScore > g.Round.ComputerScore {
		fmt.Println("Player wins!")
	} else {
		fmt.Println("Computer wins!")
	}
	fmt.Printf("Computer wins %d/3\n", g.Round.ComputerScore)
	fmt.Printf("Player wins %d/3\n", g.Round.PlayerScore)
}
