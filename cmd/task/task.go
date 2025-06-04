package task

import (
	"github.com/wriverarincon/command"
)

type TaskCommand struct {
	registry command.Registry
}

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

func NewTaskCommand() command.Command {
	return &TaskCommand{}
}

func (t *TaskCommand) Execute(args []string) error {
	return nil
}

func (t *TaskCommand) Metadata() command.MetaData {
	return taskMetaData
}

func AddSubcommands(registry *command.Registry) {
	subCommands := map[string]handler{
		"create": {[]string{"task", "create"}, NewCreateSubcommand()},
	}
	for _, v := range subCommands {
		registry.New(v.path, v.fn, nil)
	}
}
