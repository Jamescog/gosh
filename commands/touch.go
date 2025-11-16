package commands

import (
	"fmt"
	"os"
	"time"
)

type TouchCommand struct{}

func (c *TouchCommand) Name() string {
	return "touch"
}

func (c *TouchCommand) Description() string {
	return "create or update file timestamps"
}

func (c *TouchCommand) Usage() string {
	return "touch FILE..."
}

func (c *TouchCommand) Examples() []Example {
	return []Example{
		{"touch file.txt", "create or update file.txt"},
		{"touch file1.txt file2.txt", "create or update multiple files"},
	}
}

func (c *TouchCommand) Execute(paths []string, flags []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("missing file operand")
	}

	for _, file := range paths {
		if err := touchFile(file); err != nil {
			return fmt.Errorf("cannot touch '%s': %v", file, err)
		}
	}
	return nil
}

func init() {
	Register(&TouchCommand{})
}

func touchFile(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		return file.Close()
	}

	now := time.Now()
	return os.Chtimes(path, now, now)
}
