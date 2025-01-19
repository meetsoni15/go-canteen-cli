package main

import (
	"fmt"
	"strings"

	"github.com/meetsoni15/go-canteen-cli/pkg/colors"
)

func (a *app) intro() {
	fmt.Println(colors.Title("Canteen Management System"))
	fmt.Println(colors.Title("============"))
}

func (a *app) introOptions() {
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
