package commands

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

type Handler func(args []string)
