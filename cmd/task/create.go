package task

import "github.com/wriverarincon/command"

type taskCreate struct{}

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
	return nil
}

func (c taskCreate) Metadata() command.MetaData {
	return command.MetaData{}
}

func Init() {
	return
}
