package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/wriverarincon/clock-in/internal/repository/commands"
)

type HelpCommand struct {
	Registry *commands.Registry
}

var HelpMetaData = commands.MetaData{
	ShortDescription: "Prints this message",
	LongDescription:  "Prints this message along with all the available commands.",
}

func printCommands(commands map[string]commands.Command) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Println("CORE COMMANDS")
	for k, v := range commands {
		shortd := v.MetaData.ShortDescription
		fmt.Fprintln(writer, "  "+strings.Join([]string{k, shortd}, "\t")+"\t")
	}
	writer.Flush()
}

func printSubCommands(subCommands [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Println("SUB COMMANDS")
	for _, v := range subCommands {
		fmt.Fprintln(writer, "  "+strings.Join([]string{v[0], v[1]}, "\t")+"\n")
	}
	writer.Flush()
}

func printFlags(flags [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Println("FLAGS")
	for _, v := range flags {
		fmt.Fprintln(writer, "  "+strings.Join([]string{v[0], v[1]}, "\t"))
	}
	writer.Flush()
}

func (p *HelpCommand) Execute(args []string) {
	r := p.Registry
	if len(args) >= 1 {
		printSubCommands(r.Commands[args[0]].MetaData.SubCommands)
		printFlags(r.Commands[args[0]].MetaData.Flags)

	} else {
		fmt.Println("USAGE\n  <app> <command> <subcommand> [flags]\n")
		printCommands(r.Commands)
	}
}
