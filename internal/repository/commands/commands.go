package commands

import (
	"fmt"
	"os"
)

type Handler interface {
	Execute(args []string)
}

type SubHandler func(args []string)

type MetaData struct {
	ShortDescription string
	LongDescription  string
	AvailableOptions [][]string
}

type command struct {
	handler    Handler
	subCommand map[string]SubHandler
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

func (cr *Registry) NewCommand(name string, handler Handler, meta MetaData) {
	cr.Commands[name] = command{
		handler:    handler,
		subCommand: make(map[string]SubHandler),
		MetaData:   meta,
		flags:      nil,
	}
}

func (cr *Registry) AddSubCommand(parent string, name string, handler SubHandler) {
	cr.Commands[parent].subCommand[name] = handler
}

func (cr *Registry) Execute() {
	if len(os.Args) < 2 {
		cr.Commands["help"].handler.Execute(nil)
		return
	}
	if cmd, ok := cr.Commands[os.Args[1]]; ok {
		cmd.handler.Execute(os.Args)
	} else {
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func (cr *Registry) ExecuteSubCommand(args []string) {
	c := cr.Commands[args[1]]
	if ok := c.subCommand[args[2]]; ok != nil {
		ok(args[3:])
		return
	}
	fmt.Printf("Subcommand %v was not found", args[2])
}
