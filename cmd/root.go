package cmd

import (
	"os"

	"github.com/adityakeyal/dbch/command"
)

//var cfgFile string

type root struct {
	defaultCommand string
}

var r = root{"exec"}

var rootCmd = &command.Command{
	Use:     "Test Command",
	Short:   "Short Desc",
	Long:    `Long Desc`,
	Name:    "root",
	Execute: root_execute,
}

func root_execute(args []string) {

	panic("Nothing defined")

	//loop over all child commands and check if any of them satisfy the sub command or not
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	args := os.Args[1:]

	if len(args) > 0 {

		subCommand := args[0]

		isExecuted := false

		for _, x := range rootCmd.SubCommand {

			if subCommand == x.Name {
				x.Execute(args)
				isExecuted = true
				break
			}

		}

		if !isExecuted {
			panic("Invalid option")
		}

	} else {
		r.checkDefaultCommand(args)
	}

}

func (r *root) checkDefaultCommand(args []string) {
	for _, x := range rootCmd.SubCommand {
		if r.defaultCommand == x.Name {
			x.Execute(args)
			break
		}
	}
}
