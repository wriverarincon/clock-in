package cmd

type InitCommand struct{}

// InitCommand starts the app and prints a custom message for the user
func (init *InitCommand) Execute(args []string) {
	// TODO: Add custom message with user name
	println("Up and Running!")
	// TODO: Add a state changer: stopped > [run this command] > running
	// IF APP.STATE == "STOPPED" THEN APP.STATE.START()
}
