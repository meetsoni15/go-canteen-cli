package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/meetsoni15/go-canteen-cli/pkg/colors"
	"github.com/meetsoni15/go-canteen-cli/pkg/logger"
)

type app struct {
	log     *slog.Logger
	scanner *bufio.Scanner
	// create channel that indicates user wants to quit
	doneChan chan struct{}
}

func main() {
	var app = app{
		log:      logger.New(),
		scanner:  bufio.NewScanner(os.Stdin),
		doneChan: make(chan struct{}),
	}
	app.startScreen()

}

func (a *app) startScreen() {
	// print welcome message
	a.intro()

	// run user input go routine
	go a.readUserInput()

	// block go routine
	<-a.doneChan

	// close channel
	close(a.doneChan)

	// print bye
	fmt.Print("GoodBye")
}

func (a *app) intro() {
	fmt.Println(colors.Title("Canteen Management System"))
	fmt.Println(colors.Title("============"))
	fmt.Println(colors.Normal("How would you like to login - as a customer or as owner??"))
	fmt.Println(colors.Highlight(strings.Repeat(" ", 15) + "1. Customer"))
	fmt.Println(colors.Highlight(strings.Repeat(" ", 15) + "2. Owner"))
	fmt.Println("Your choice:")
	fmt.Println(colors.Quit("Enter q for quit."))
	a.prompt()
}

func (a *app) prompt() {
	fmt.Print("-> ")
}

func (a *app) readUserInput() {
	for {
		res, done := a.checkOptions(a.scanner)
		if done {
			a.doneChan <- struct{}{}
			return
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

	return "msg", false
}
