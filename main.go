package main

import (
	"github.com/wriverarincon/clock-in/cmd"
	"github.com/wriverarincon/clock-in/internal/repository/commands"
)

func main() {
	// Create the registry
	registry := commands.NewRegistry()

	init := &cmd.InitCommand{}
	help := &cmd.HelpCommand{Registry: registry}
	task := &cmd.TaskCommand{Registry: registry}

	registry.NewCommand("init", init.Init, cmd.InitMetaData)
	registry.NewCommand("help", help.Init, cmd.HelpMetaData)
	registry.NewCommand("task", task.Init, cmd.TaskMetaData)
	registry.Execute()
}
