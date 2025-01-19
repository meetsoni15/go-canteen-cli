package main

import (
	"fmt"
	"strings"

	"github.com/meetsoni15/go-canteen-cli/pkg/colors"
	ownerauth "github.com/meetsoni15/go-canteen-cli/pkg/ownerAuth"
)

func (a *app) ownerAuth() (string, bool) {
	fmt.Println(colors.Highlight("\nYou chose to login as an Owner."))

	// Ask for Owner's username
	fmt.Print(colors.Normal("Enter Owner username: "))
	a.scanner.Scan()
	username := strings.TrimSpace(a.scanner.Text())

	// Ask for Owner's password
	fmt.Print(colors.Normal("Enter Owner password: "))
	a.scanner.Scan()
	password := strings.TrimSpace(a.scanner.Text())

	if ownerauth.DoOwnerAuth(username, password) {
		return colors.Success("\nWelcome, Owner!"), false
	}

	return colors.Error("\nInvalid username or password. Access denied."), false
}
