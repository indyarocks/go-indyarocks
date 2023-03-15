package main

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

func main() {
	playerChoice := ""
	playerValue := -1

	computerChoice := rand.Intn(3)

	clearScreen()

	fmt.Print("Please enter your choice -> ")
	reader := bufio.NewReader(os.Stdin)
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	lowerPlayerChoice := strings.ToLower(playerChoice)

	fmt.Println()
	if lowerPlayerChoice == "rock" {
		playerValue = ROCK
	} else if lowerPlayerChoice == "paper" {
		playerValue = PAPER
	} else if lowerPlayerChoice == "scissors" {
		playerValue = SCISSORS
	}
	fmt.Println("You choose: ", playerChoice, "Your value: ", playerValue)
	fmt.Println("Computer choose:", computerChoice)
}

func clearScreen() {
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
