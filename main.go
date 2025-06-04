package main

import (
	"os"

	"github.com/wriverarincon/clock-in/cmd/task"
	"github.com/wriverarincon/command"
)

func main() {
	// Create the registry
	registry := command.NewRegistry()

	registry.New(nil, task.NewTaskCommand(), nil)
	task.AddSubcommands(registry)
	registry.Execute(os.Args[1:])
}
