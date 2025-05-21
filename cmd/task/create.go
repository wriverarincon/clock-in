package task

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/wriverarincon/command"
)

type taskCreate struct{}

var titleFlag = command.NewFlag("title", "t", "Title of the task", "", true)
var bodyFlag = command.NewFlag("body", "b", "Body of the task", "", true)
var startFlag = command.NewFlag("start-day", "", "Start day of the task, defaults to today", time.Now().Format("2006-01-02"), false)
var endFlag = command.NewFlag("end-day", "", "End day of the task, defaults to today", time.Now().Format("2006-01-02"), false)

var createMetaData = command.MetaData{
	Name:             "create",
	ShortDescription: "Creates a task",
	LongDescription:  "Creates a task with the defined options",
	Flags:            []command.Flag{titleFlag, bodyFlag, startFlag, endFlag},
}

func (c taskCreate) Execute(args []string) error {
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
		return err
	}
	fmt.Printf("got %v", title)
	return nil
}

func (c taskCreate) Metadata() command.MetaData {
	return createMetaData
}
