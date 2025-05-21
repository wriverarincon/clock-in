package task

import "github.com/wriverarincon/command"

type SubcommandCreate struct{}

func (c SubcommandCreate) Execute(args []string) error {
	return nil
}

func (c SubcommandCreate) Metadata() command.MetaData {
	return command.MetaData{}
}

func Init() {
	return
}
