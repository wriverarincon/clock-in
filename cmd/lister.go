package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/wriverarincon/clock-in/internals/repository/commands"
)

type HelpCommand struct {
	Registry *commands.Registry
}

func (p *HelpCommand) Execute(args []string) {
	r := p.Registry
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.AlignRight)
	fmt.Println("COMMANDS")
	for k, v := range r.Commands {
		shortDescription := v.MetaData.ShortDescription
		fmt.Fprintln(writer, strings.Join([]string{k, shortDescription}, "\t")+"\t")
	}
	writer.Flush()
}
