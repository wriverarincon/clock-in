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

var (
	title string
	body  string
	from  string
	until string
)

func (t *TaskCommand) createTask(args []string) {
	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.StringVar(&title, "title", "", "task title")
	fs.StringVar(&body, "body", "", "task body")

	if err := fs.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}
}

func (t *TaskCommand) Execute(args []string) {
	t.Registry.AddSubCommand(args[1], "create", t.createTask)

	t.Registry.ExecuteSubCommand(args)
	fmt.Printf("Task %v created\n", title)
}
