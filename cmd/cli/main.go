package main

import (
	"fmt"
	"os"

	"github.com/pedro3g/bvm/handlers"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a command. Available commands are:")
		fmt.Println(" - list: List all available versions")
		return
	}

	mainArg := args[0]

	if mainArg == "list" {
		handlers.ListVersions(true)
	} else {
		fmt.Printf("Command '%s' not available\n", mainArg)
	}
}
