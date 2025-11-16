package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct{}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Description() string {
	return "exit the shell"
}

func (c *ExitCommand) Usage() string {
	return "exit"
}

func (c *ExitCommand) Examples() []Example {
	return []Example{
		{"exit", "exit the shell"},
	}
}

func (c *ExitCommand) Execute(paths []string, flags []string) error {
	fmt.Println("Exiting gosh...")
	os.Exit(0)
	return nil
}

func init() {
	Register(&ExitCommand{})
}
