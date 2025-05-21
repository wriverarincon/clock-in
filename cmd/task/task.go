package task

import (
	"github.com/wriverarincon/command"
)

type TaskCommand struct{}

type handler struct {
	path []string
	fn   command.Command
}

var taskMetaData = command.MetaData{
	Name:             "task",
	ShortDescription: "Create task",
	LongDescription:  "Creates and stores a task for the day and [OPTIONS] specified",
	Flags:            []command.Flag{},
}

func New() command.Command {
	return &TaskCommand{}
}

func (t *TaskCommand) Execute(args []string) error {
	return nil
}

func (t *TaskCommand) Metadata() command.MetaData {
	return taskMetaData
}

func Setup(registry *command.Registry) func() {
	return func() {
		subCommands := map[string]handler{
			"create": {[]string{"task", "create"}, taskCreate{}},
		}
		for _, v := range subCommands {
			registry.New(v.path, v.fn, nil)
		}
	}
}
