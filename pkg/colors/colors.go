package colors

import "github.com/fatih/color"

var Title = color.New(color.FgCyan, color.Bold).SprintFunc()
var Highlight = color.New(color.FgGreen).SprintFunc()
var Normal = color.New(color.FgWhite).SprintFunc()
var Quit = color.New(color.FgRed).SprintFunc()
var Success = color.New(color.FgGreen).SprintFunc()
var Error = color.New(color.FgRed).SprintFunc()
