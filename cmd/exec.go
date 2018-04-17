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
		systemInfo := LoadConfiguration()
		for _, fileDetails := range systemInfo.GetDefaultEnvironment().ReplaceInfo {
			if fileDetails.Filename == info.Name() {
				data, _ := ioutil.ReadFile(path)
				for _, change := range fileDetails.Changes {

					switch fileDetails.FileType {
					case "prop":
						data = propReplacer(data, change.Source, change.Target)
					case "xml":
						data = xmlReplacer(data, change.Source, change.Target)
					default:
					}

				}
				ioutil.WriteFile(path, data, info.Mode())
			}
		}
	}
	return nil
}

func (exec *exec) execute(args []string) {
	cwd, _ := os.Getwd()

	filepath.Walk(cwd, exec.visit)

}

func init() {
	command.RootCmd.AddDefaultCommand(execCmd)
}
