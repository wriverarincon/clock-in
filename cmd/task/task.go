package task

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/wriverarincon/command"
)

type TaskCommand struct{}

type handler struct {
	path []string
	fn   command.Command
}

var titleFlag = command.Flag{
	Name:         "title",
	Shorthand:    "t",
	Description:  "Title of the task",
	DefaultValue: "",
	Required:     true,
}

var bodyFlag = command.Flag{
	Name:         "body",
	Shorthand:    "b",
	Description:  "Body of the task",
	DefaultValue: "",
	Required:     true,
}

var startFlag = command.Flag{
	Name:         "start-day",
	Shorthand:    "",
	Description:  "Start day of the task, defaults to today",
	DefaultValue: time.Now().Format("2006-01-02"),
}

var endFlag = command.Flag{
	Name:         "end-day",
	Shorthand:    "",
	Description:  "End day of the task, defaults to today",
	DefaultValue: time.Now().Format("2006-01-02"),
}

var TaskMetaData = command.MetaData{
	ShortDescription: "Create task",
	LongDescription:  "Creates and stores a task for the day and [OPTIONS] specified",
	Flags:            []command.Flag{titleFlag, bodyFlag, startFlag, endFlag},
}

func New() command.Command {
	return &TaskCommand{}
}

func (t *TaskCommand) Execute(args []string) error {
	return nil
}

func (t *TaskCommand) Metadata() command.MetaData {
	return TaskMetaData
}

func Setup(registry *command.Registry) {
	subCommands := map[string]handler{
		"create": {[]string{"task", "create"}, taskCreate{}},
	}
	for _, v := range subCommands {
		registry.New(v.path, v.fn)
	}
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
