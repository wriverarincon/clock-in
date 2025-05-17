package commands

import (
	"fmt"
	"os"
)

type Handler interface {
	Execute(args []string)
}

type MetaData struct {
	ShortDescription string
	LongDescription  string
	AvailableOptions [][]string
}

type command struct {
	handler    Handler
	subCommand Handler
	MetaData   MetaData
	flags      []string
}

type Registry struct {
	Commands map[string]command
}

type option string

func NewRegistry() *Registry {
	return &Registry{
		Commands: make(map[string]command),
	}
}

func (cr *Registry) NewCommand(name string, handler Handler, subCommand Handler, metadata MetaData, flags []string) {
	c := command{
		handler:    handler,
		subCommand: subCommand,
		MetaData:   metadata,
		flags:      flags,
	}
	cr.Commands[name] = c
}

func (cr *Registry) Execute() {
	if len(os.Args) < 2 {
		cr.Commands["help"].handler.Execute(nil)
		return
	}
	if cmd, ok := cr.Commands[os.Args[1]]; ok {
		cmd.handler.Execute(os.Args[2:])
	} else {
		fmt.Println("Unknown command:", os.Args[1])
	}
}
