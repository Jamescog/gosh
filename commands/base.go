package commands

// Command defines the interface that all shell commands must implement.
// This enables automatic registration, help generation, and uniform execution.
type Command interface {
	// Name returns the command name as used in the shell
	Name() string

	// Description returns a short one-line description of what the command does
	Description() string

	// Usage returns the usage syntax (e.g., "ls [OPTIONS] [PATH]")
	Usage() string

	// Examples returns a list of example usages with descriptions
	Examples() []Example

	// Execute runs the command with the given paths and flags
	Execute(paths []string, flags []string) error
}

// Example represents a usage example for a command
type Example struct {
	Command     string
	Description string
}
