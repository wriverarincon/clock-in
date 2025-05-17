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

	initMetaData := commands.MetaData{
		ShortDescription: "Initializes the app",
		LongDescription: `The "init" command initializes the app, retrieving the tasks completed the previos day,
						  along with basic analytics facts.`,
		Flags: [][]string{
			{"-Q, --quiet", "Initializes the app in quiet mode, not returning any output."},
		},
	}
	helpMetaData := commands.MetaData{
		ShortDescription: "Prints this message",
		LongDescription:  "Prints this message along with all the available commands.",
	}
	createTaskMetaData := commands.MetaData{
		ShortDescription: "Create task",
		LongDescription:  "Creates and stores a task for the day and [OPTIONS] specified",
		SubCommands: [][]string{
			{"create", "create a new task"},
		},
		Flags: [][]string{
			{"-t, --title", "Title of the task"},
			{"-b, --body", "Body of the task"},
			{"--start-day", "Determines the start day of the task, defaults to today"},
			{"--end-day", "Determines the end day for the task, defaults to today"},
		},
	}

	registry.NewCommand("init", init, initMetaData)
	registry.NewCommand("help", help, helpMetaData)
	registry.NewCommand("task", task, createTaskMetaData)
	registry.Execute()
}
