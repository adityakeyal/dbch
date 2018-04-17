package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

//LoadConfiguration -
func LoadConfiguration() *SystemInfo {

	home, _ := homedir.Dir()

	basePath := path.Join(home, ".dbch", "env.json")

	data, err := ioutil.ReadFile(basePath)

	if err != nil {
		return &SystemInfo{}
	}

	var systemInfo SystemInfo

	json.Unmarshal(data, &systemInfo)

	return &systemInfo

}

func init() {

}
