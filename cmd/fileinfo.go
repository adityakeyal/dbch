package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

//Something -
func Something() {
	slcB, _ := json.Marshal(System)
	x := string(slcB)
	fmt.Println(x)

	home, _ := homedir.Dir()
	fmt.Println(home)
	basePath := path.Join(home, ".dbch")
	os.Mkdir(basePath, os.ModePerm)
	filePath := path.Join(basePath, "env.json")
	ioutil.WriteFile(filePath, slcB, os.ModePerm)
	// json.Unmarshal([]byte(jsonStr), &xx)
	// fmt.Println(xx)

}
