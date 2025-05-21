package task

import (
	"time"

	"github.com/wriverarincon/command"
)

type TaskCommand struct{}

type handler struct {
	path []string
	fn   command.Command
}

var titleFlag = command.NewFlag("title", "t", "Title of the task", "", true)
var bodyFlag = command.NewFlag("body", "b", "Body of the task", "", true)
var startFlag = command.NewFlag("start-day", "", "Start day of the task, defaults to today", time.Now().Format("2006-01-02"), false)
var endFlag = command.NewFlag("end-day", "", "End day of the task, defaults to today", time.Now().Format("2006-01-02"), false)

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
