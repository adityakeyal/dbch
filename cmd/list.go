package cmd

import (
	"flag"

	"github.com/adityakeyal/dbch/command"
)

var l = list{}

var newlist = &command.Command{
	Name:    "mylist",
	Use:     "Test Command",
	Short:   "Short Desc",
	Long:    `Long Desc`,
	Execute: l.execute,
}

type list struct {
}

func (list *list) execute(args []string) {
	panic("Not yet defined")
}

type listFlag struct {
	all  *bool
	name *string
}

func (list *list) parseArguments(args []string) listFlag {
	var flags listFlag

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	flags.name = listCommand.String("name", "", "Name of the command")
	flags.all = listCommand.Bool("all", false, "List All")
	listCommand.Parse(args[1:])

	return flags
}

func init() {
	rootCmd.AddCommand(newlist)
}
