package commands

import (
	"fmt"
	"os"
)

func NewRegistry() *Registry {
	return &Registry{
		Commands: make(map[string]Command),
	}
}

func (cr *Registry) NewCommand(name string, handler Handler, meta MetaData) {
	cr.Commands[name] = Command{
		handler:     handler,
		SubCommands: make(map[string]Command),
		MetaData:    meta,
	}
}

func (cr *Registry) AddSubCommand(parent string, name string, handler Handler, meta MetaData) {
	c := Command{
		handler:  handler,
		MetaData: meta,
	}
	cr.Commands[parent].SubCommands[name] = c
}

func (cr *Registry) Execute() {
	args := os.Args
	if len(args) < 2 {
		cr.Commands["help"].handler(nil)
		return
	}
	if cmd, ok := cr.Commands[args[1]]; ok {
		cmd.handler(args)
	} else {
		fmt.Println("Unknown command:", args[1])
	}
}

func (cr *Registry) ExecuteSubCommand(args []string) {
	c := cr.Commands[args[1]]
	if h := c.SubCommands[args[2]].handler; h != nil {
		h(args[3:])
		return
	}
	fmt.Printf("Subcommand %v was not found", args[2])
}
