package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adityakeyal/gocli/command"
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
		for _, fileDetails := range System.Environments[0].ReplaceInfo {
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
	// cwd, _ := os.Getwd()
	cwd := "D:/code/igv/nrithai/BR-GV-FOR-FIFO-FROM-THAGMO-PROD-REL-14-06032018/console/igv-console/CONSOLE-INF"

	filepath.Walk(cwd, exec.visit)

}

func init() {
	command.RootCmd.AddDefaultCommand(execCmd)
}
