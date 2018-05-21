//Package cmd -
// Entry point in the cmd package. This is invoked by the main() of the
// main.go application.
////////////////
package cmd

import (
	"github.com/adityakeyal/gocli/command"
)

//Execute - Start the execution of Command
//This loads the gocli command chain. All the registered commands are
//now made available by the gocli toolkit
func Execute() {
	command.Execute()
}
