package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adityakeyal/dbch/command"
)

var e = exec{}

var execCmd = &command.Command{
	Name:    "exec",
	Use:     "Test Command",
	Short:   "Short Desc",
	Long:    `Long Desc`,
	Execute: e.execute,
}

type exec struct {
}

func (exec *exec) visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		for _, fileDetails := range DevEnvironment.ReplaceInfo {
			if fileDetails.Filename == info.Name() {
				data, _ := ioutil.ReadFile(path)
				for _, change := range fileDetails.Changes {
					data = propReplacer(data, change.Source, change.Target)
				}
				ioutil.WriteFile(path, data, info.Mode())
			}
		}
	}
	return nil
}

func (exec *exec) execute(args []string) {
	fmt.Println("cwd")
	// cwd, _ := os.Getwd()
	cwd := "D:/code/igv/nrithai/BR-GV-FOR-FIFO-FROM-THAGMO-PROD-REL-14-06032018/console/igv-console/CONSOLE-INF"

	filepath.Walk(cwd, exec.visit)

}

// type execFlag struct {
// 	all  *bool
// 	name *string
// }

// func (exec *exec) parseArguments(args []string) execFlag {
// 	var flags execFlag

// 	execCommand := flag.NewFlagSet("exec", flag.ExitOnError)
// 	flags.name = execCommand.String("name", "", "Name of the command")
// 	flags.all = execCommand.Bool("all", false, "exec All")
// 	execCommand.Parse(args[1:])

// 	return flags
// }

func init() {
	rootCmd.AddCommand(execCmd)
}
