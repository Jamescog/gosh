package commands

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Description() string {
	return "change the current directory"
}

func (c *CdCommand) Usage() string {
	return "cd DIRECTORY"
}

func (c *CdCommand) Examples() []Example {
	return []Example{
		{"cd /home", "change to /home directory"},
		{"cd ..", "change to parent directory"},
		{"cd ~", "change to home directory"},
	}
}

func (c *CdCommand) Execute(paths []string, flags []string) error {
	if len(paths) < 1 {
		return fmt.Errorf("missing operand")
	}
	err := os.Chdir(paths[0])
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func init() {
	Register(&CdCommand{})
}
