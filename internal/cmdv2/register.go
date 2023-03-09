package cmdv2

import (
	"fmt"
)

var registeredCommands = []Command{}

// RegisterCommand adds a new Command type to the list of available commands in the CLI.
func RegisterCommand(command Command) {
	for _, registeredCommand := range registeredCommands {
		if registeredCommand.Type() == command.Type() {
			err := fmt.Errorf("command type already registered: %s", command.Type())
			panic(err)
		}
	}

	registeredCommands = append(registeredCommands, command)
}

// RegisteredCommands returns all registered Command types in the CLI.
func RegisteredCommands() []Command {
	return registeredCommands
}
