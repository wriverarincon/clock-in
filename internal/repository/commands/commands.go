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
	SubCommands      [][]string
	Flags            [][]string
}

type Registry struct {
	Commands map[string]Command
}

type Command struct {
	handler     Handler
	SubCommands map[string]Command
	MetaData    MetaData
}

type SubHandler func(args []string)

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

func (cr *Registry) AddSubCommand(parent string, name string, handler SubHandler) {
	cr.Commands[parent].SubCommands[name] = handler
}

func (cr *Registry) Execute() {
	args := os.Args
	if len(args) < 2 {
		cr.Commands["help"].handler(nil)
		return
	}
	fmt.Println(args[1], cr.Commands)
	if cmd, ok := cr.Commands[args[1]]; ok {
		cmd.handler(args[2:])
	} else {
		fmt.Println("Unknown command:", args[1])
	}
}

func (cr *Registry) ExecuteSubCommand(args []string) {
	c := cr.Commands[args[1]]
	if ok := c.SubCommands[args[2]]; ok != nil {
		ok(args[3:])
		return
	}
	fmt.Printf("Subcommand %v was not found", args[2])
}
