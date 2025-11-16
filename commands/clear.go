package commands

import "fmt"

type ClearCommand struct{}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) Description() string {
	return "clear the terminal screen"
}

func (c *ClearCommand) Usage() string {
	return "clear"
}

func (c *ClearCommand) Examples() []Example {
	return []Example{
		{"clear", "clear the screen"},
	}
}

func (c *ClearCommand) Execute(paths []string, flags []string) error {
	fmt.Print("\033[2J\033[H")
	return nil
}

func init() {
	Register(&ClearCommand{})
}
