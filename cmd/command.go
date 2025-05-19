package cmd

import "github.com/wriverarincon/command"

type customCommand struct {
	command  *command.Command
	metadata *command.MetaData
}

func NewCommand() *customCommand {
	return &customCommand{}
}
