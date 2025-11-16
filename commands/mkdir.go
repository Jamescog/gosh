package commands

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type MkdirCommand struct{}

func (c *MkdirCommand) Name() string {
	return "mkdir"
}

func (c *MkdirCommand) Description() string {
	return "create directories"
}

func (c *MkdirCommand) Usage() string {
	return "mkdir [OPTIONS] DIRECTORY..."
}

func (c *MkdirCommand) Examples() []Example {
	return []Example{
		{"mkdir newdir", "create a new directory"},
		{"mkdir -p dir1/dir2/dir3", "create nested directories"},
		{"mkdir -v newdir", "create directory with verbose output"},
		{"mkdir -m 0700 private", "create directory with specific permissions"},
	}
}

func (c *MkdirCommand) Execute(paths []string, flags []string) error {
	return handleMkdir(paths, flags)
}

func init() {
	Register(&MkdirCommand{})
}

func handleMkdir(paths []string, flags []string) error {
	fs := flag.NewFlagSet("mkdir", flag.ContinueOnError)
	createParents := fs.Bool("p", false, "no error if existing, make parent directories as needed")
	verboseOutput := fs.Bool("v", false, "print a message for each created directory")
	permissionMode := fs.String("m", "0755", "set file mode (as in chmod), not a=rwx - umask")

	if err := fs.Parse(flags); err != nil {
		return err
	}

	fileMode := parseFileMode(*permissionMode)

	for _, directory := range paths {
		if err := createDirectory(directory, fileMode, *createParents, *verboseOutput); err != nil {
			return fmt.Errorf("cannot create directory '%s': %v", directory, err)
		}
	}
	return nil
}

func createDirectory(path string, mode os.FileMode, createParents, verbose bool) error {
	if !createParents {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("file exists")
		}
	}

	if err := os.MkdirAll(path, mode); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("mkdir: created directory '%s'\n", path)
	}

	return nil
}

func parseFileMode(modeStr string) os.FileMode {
	mode, err := strconv.ParseUint(modeStr, 8, 32)
	if err != nil {
		return 0755
	}
	return os.FileMode(mode)
}
