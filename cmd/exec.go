//Exec - This is the default command. This is used to run
//the command which changes the property and xml files based on an environment

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/adityakeyal/gocli/command"
)

var e = exec{}

var execCmd = &command.Command{
	Name:  "exec",
	Use:   "",
	Short: "Replace file content defined for environment",
	Long: `This command will replace all environment properties identified by 
	  - provided the environment
	  - default environment
	`,
	Execute: e.execute,
	SubHelp: func() string {

		return `
	If an environment is provided then the values are replaced for that environment else the default will be run.
	To know details of the environment check the list command.
		`
	},
}

type exec struct {
}

func (exec *exec) visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		systemInfo := LoadConfiguration()
		environment := ""
		if len(os.Args) > 2 {
			environment = os.Args[2]
		}
		for _, fileDetails := range systemInfo.GetEnvironment(environment).ReplaceInfo {
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

//The code where the actual logic lies.
//This is used to do a lot of stuff like :
//  Loop over all folders and determine where the files are present
func (exec *exec) execute(args []string) {

	cwd, _ := os.Getwd()

	filepath.Walk(cwd, exec.visit)

}

func init() {
	command.RootCmd.AddDefaultCommand(execCmd)
}
