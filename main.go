package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jamescog/gosh/commands"
	"github.com/Jamescog/gosh/utils"
)

func main() {
	utils.InitColor()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(utils.Green("-> gosh:"))
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "" {
			continue
		}

		parsed := utils.ParseCommand(input)
		command := parsed.Command

		if !commands.DefaultRegistry.Exists(command) {
			fmt.Printf(utils.Red("command not found: %s\n"), command)
			fmt.Println("Type 'help' to see available commands.")
			continue
		}

		if err := commands.DefaultRegistry.Execute(command, parsed.Paths, parsed.Flags); err != nil {
			fmt.Printf(utils.Red("%s: %v\n"), command, err)
		}
	}
}
