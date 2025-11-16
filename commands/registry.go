package commands

import (
	"fmt"
	"sort"
)

// Registry manages all available shell commands
type Registry struct {
	commands map[string]Command
}

// NewRegistry creates a new command registry
func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]Command),
	}
}

// Register adds a command to the registry
func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
}

// Get retrieves a command by name
func (r *Registry) Get(name string) (Command, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}

// List returns all registered commands sorted by name
func (r *Registry) List() []Command {
	cmds := make([]Command, 0, len(r.commands))
	for _, cmd := range r.commands {
		cmds = append(cmds, cmd)
	}
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Name() < cmds[j].Name()
	})
	return cmds
}

// Exists checks if a command is registered
func (r *Registry) Exists(name string) bool {
	_, exists := r.commands[name]
	return exists
}

// Execute runs a command by name with the given arguments
func (r *Registry) Execute(name string, paths []string, flags []string) error {
	cmd, exists := r.Get(name)
	if !exists {
		return fmt.Errorf("command not found: %s", name)
	}
	return cmd.Execute(paths, flags)
}

// DefaultRegistry is the global command registry
var DefaultRegistry = NewRegistry()

// Register adds a command to the default registry
func Register(cmd Command) {
	DefaultRegistry.Register(cmd)
}
