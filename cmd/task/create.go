package task

import "github.com/wriverarincon/command"

type taskCreate struct{}

func (c taskCreate) Execute(args []string) error {
	return nil
}

func (c taskCreate) Metadata() command.MetaData {
	return command.MetaData{}
}

func Init() {
	return
}
