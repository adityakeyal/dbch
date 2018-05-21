//  list command is used to list the available environment and file changes for the environment
// The options when runnig the command is as below:
// "-name=<Environment Name>" (optional)
// If no name is provided then all the environment are shown
//

package cmd

import (
	"flag"
	"fmt"

	"github.com/adityakeyal/gocli/command"
)

var l = list{}

var newlist = &command.Command{
	Name:    "list",
	Use:     "dbch.exe list -name=<Environment Name>",
	Short:   "List all environments or environment by name",
	Long:    `Command provides an option to list all the environment and the change options for the environment`,
	Execute: l.execute,
}

type list struct {
}

func (list *list) execute(args []string) {

	flags := l.parseArguments(args)

	systemDetails := LoadConfiguration()

	for _, env := range systemDetails.Environments {

		if *flags.name != "" && *flags.name != env.EnvironmentName {
			continue
		}

		defaultTxt := " "
		if env.Default {
			defaultTxt = " (Default)"
		}

		fmt.Println("Environment : " + env.EnvironmentName + defaultTxt)

		for _, fc := range env.ReplaceInfo {

			fmt.Println("\tFile : " + fc.Filename)
			for _, ch := range fc.Changes {
				fmt.Println("\t\t Key: " + ch.Source + " , Value :  " + ch.Target)
			}

		}

	}
}

type listFlag struct {
	name *string
}

func (list *list) parseArguments(args []string) listFlag {
	var flags listFlag

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	flags.name = listCommand.String("name", "", "Name of the Environment (If none specified list all) ")
	listCommand.Parse(args[1:])

	return flags
}

func init() {

	command.RootCmd.AddCommand(newlist)
}
