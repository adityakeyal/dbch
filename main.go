// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/adityakeyal/dbch/cmd"
)

func main() {

	slcB, _ := json.Marshal(cmd.Replace)
	fmt.Println(string(slcB))

	jsonStr := `[{"Filename":"xenos-jdbc.properties","Changes":[{"Source":"jdbc.components.url.GLOBAL","Target":"jdbc:oracle:thin:@localhost:1521:XE"},{"Source":"jdbc.components.userName.GLOBAL","Target":"gvth_dev_global"},{"Source":"jdbc.components.password.GLOBAL","Target":"gvth_dev_global"}]},{"Filename":"GMO-jdbc.properties","Changes":[{"Source":"jdbc.components.url.GLOBAL","Target":"jdbc:oracle:thin:@localhost:1521:XE"},{"Source":"jdbc.components.userName.GLOBAL","Target":"gvth_dev_gmo_txn"},{"Source":"jdbc.components.password.GLOBAL","Target":"gvth_dev_gmo_txn"}]},{"Filename":"NRI-jdbc.properties","Changes":[{"Source":"jdbc.components.url.GLOBAL","Target":"jdbc:oracle:thin:@localhost:1521:XE"},{"Source":"jdbc.components.userName.GLOBAL","Target":"gvth_dev_nri_txn"},{"Source":"jdbc.components.password.GLOBAL","Target":"gvth_dev_nri_txn"}]}]`
	var xx []cmd.FileDetails
	json.Unmarshal([]byte(jsonStr), &xx)
	fmt.Println(xx)

	//cmd.Execute()
}
