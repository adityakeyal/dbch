////////////////
// Entry point of the application
////////////////
package cmd

import (
	"github.com/adityakeyal/gocli/command"
)

//Execute - Start the execution of Command
//This just delgates the call to the main command
func Execute() {
	command.Execute()
}
