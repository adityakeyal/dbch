package command

//Command used to decide which child to call
type Command struct {
	Name       string
	Use        string
	Short      string
	Long       string
	Execute    func(args []string)
	SubCommand []*Command
}

//AddCommand - Command
func (rc *Command) AddCommand(cmd *Command) {
	rc.SubCommand = append(rc.SubCommand, cmd)
}
