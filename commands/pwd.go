package commands

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Description() string {
	return "print the current working directory"
}

func (c *PwdCommand) Usage() string {
	return "pwd"
}

func (c *PwdCommand) Examples() []Example {
	return []Example{
		{"pwd", "display current directory path"},
	}
}

func (c *PwdCommand) Execute(paths []string, flags []string) error {
	currDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}
	fmt.Println(currDir)
	return nil
}

func init() {
	Register(&PwdCommand{})
}
