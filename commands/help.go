package commands

import (
	"fmt"
	"strings"

	"github.com/Jamescog/gosh/utils"
)

type HelpCommand struct{}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "display help information about commands"
}

func (c *HelpCommand) Usage() string {
	return "help [COMMAND]"
}

func (c *HelpCommand) Examples() []Example {
	return []Example{
		{"help", "list all available commands"},
		{"help ls", "show detailed help for ls command"},
	}
}

func (c *HelpCommand) Execute(paths []string, flags []string) error {
	if len(paths) == 0 {
		return showAllCommands()
	}
	return showCommandHelp(paths[0])
}

func init() {
	Register(&HelpCommand{})
}

func showAllCommands() error {
	fmt.Println(utils.Blue("Available commands:"))
	fmt.Println()

	commands := DefaultRegistry.List()
	maxNameLen := 0
	for _, cmd := range commands {
		if len(cmd.Name()) > maxNameLen {
			maxNameLen = len(cmd.Name())
		}
	}

	for _, cmd := range commands {
		padding := strings.Repeat(" ", maxNameLen-len(cmd.Name())+2)
		fmt.Printf("  %s%s%s\n", utils.Green(cmd.Name()), padding, cmd.Description())
	}

	fmt.Println()
	fmt.Println("Use 'help <command>' for more information about a specific command.")
	return nil
}

func showCommandHelp(cmdName string) error {
	cmd, exists := DefaultRegistry.Get(cmdName)
	if !exists {
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	fmt.Println(utils.Blue("Command:"), utils.Green(cmd.Name()))
	fmt.Println()
	fmt.Println(utils.Blue("Description:"))
	fmt.Printf("  %s\n", cmd.Description())
	fmt.Println()
	fmt.Println(utils.Blue("Usage:"))
	fmt.Printf("  %s\n", cmd.Usage())

	examples := cmd.Examples()
	if len(examples) > 0 {
		fmt.Println()
		fmt.Println(utils.Blue("Examples:"))
		for _, ex := range examples {
			fmt.Printf("  %s\n", utils.Green(ex.Command))
			fmt.Printf("    %s\n", ex.Description)
		}
	}

	return nil
}
