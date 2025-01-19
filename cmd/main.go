package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/meetsoni15/go-canteen-cli/pkg/colors"
	"github.com/meetsoni15/go-canteen-cli/pkg/logger"
)

type app struct {
	log *slog.Logger
}

func main() {
	var app = app{
		log: logger.New(),
	}
	app.startScreen()

}

func (a *app) startScreen() {
	// print welcome message
	a.intro()
	// create channel that indicates user wants to quit
	doneChan := make(chan struct{})

	// run user input go routine
	go a.readUserInput(doneChan, os.Stdin)

	// block go routine
	<-doneChan

	// close channel
	close(doneChan)

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

func (a *app) readUserInput(doneChan chan struct{}, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for {
		res, done := a.checkOptions(scanner)
		if done {
			doneChan <- struct{}{}
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
