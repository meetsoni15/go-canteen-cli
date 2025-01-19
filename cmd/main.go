package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type app struct {
	scanner *bufio.Scanner
	// create channel that indicates user wants to quit
	doneChan chan struct{}
}

func main() {
	var app = app{
		scanner:  bufio.NewScanner(os.Stdin),
		doneChan: make(chan struct{}),
	}
	app.startScreen()
}

func (a *app) startScreen() {
	// print welcome message
	a.intro()
	a.introOptions()

	// run user input go routine
	go a.readUserInput()

	// block go routine
	<-a.doneChan

	// close channel
	close(a.doneChan)

	// print bye
	fmt.Print("GoodBye")
}

func (a *app) readUserInput() {
	for {
		res, done := a.checkOptions(a.scanner)
		if done {
			a.doneChan <- struct{}{}
			return
		}

		if !done {
			a.introOptions()
		}

		fmt.Println(res)
		a.prompt()
	}
}

func (a *app) checkOptions(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// get user input
	input := scanner.Text()

	// check to see if user wants to quit
	if strings.EqualFold(input, "q") {
		return "", true
	}

	switch input {
	case "1":
	case "2":
		return a.ownerAuth()
	default:
		return "Invalid Choice", false
	}

	return "", false
}
