package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/wriverarincon/clock-in/internal/repository/commands"
)

type TaskCommand struct {
	Registry *commands.Registry
}

var TaskMetaData = commands.MetaData{
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

func (t *TaskCommand) createTask(args []string) {
	var (
		title string
		body  string
		from  string
		until string
	)

	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.StringVar(&title, "title", "", "task title")
	fs.StringVar(&title, "t", "", "task title (shorthand)")
	fs.StringVar(&body, "body", "", "task body")
	fs.StringVar(&body, "b", "", "task body (shorthand)")
	fs.StringVar(&from, "from", "", "task start time")
	fs.StringVar(&from, "f", "", "task start time (shorthand)")
	fs.StringVar(&until, "until", "", "task end time")
	fs.StringVar(&until, "u", "", "task end time (shorthand)")

	if err := fs.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}
}

func (t *TaskCommand) Execute(args []string) {
	t.Registry.AddSubCommand(args[1], "create", t.createTask)

	t.Registry.ExecuteSubCommand(args)
}
