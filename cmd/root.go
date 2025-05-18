package cmd

import "github.com/wriverarincon/clock-in/internal/repository/commands"

type InitCommand struct{}

var InitMetaData = commands.MetaData{
	ShortDescription: "Initializes the app",
	LongDescription: `The "init" command initializes the app, retrieving the tasks completed the previos day,
						  along with basic analytics facts.`,
	Flags: [][]string{
		{"-Q, --quiet", "Initializes the app in quiet mode, not returning any output."},
	},
}

// InitCommand starts the app and prints a custom message for the user
func (init *InitCommand) Init(args []string) {
	// TODO: Add custom message with user name
	println("Up and Running!")
	// TODO: Add a state changer: stopped > [run this command] > running
	// IF APP.STATE == "STOPPED" THEN APP.STATE.START()
}
