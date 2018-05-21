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
	"github.com/adityakeyal/dbch/cmd"
)

//Entry point of the application.
//This delegates the call to the Execute() method of the root.go.
//This is by convention of go-cli
//Once all packages in the cmd folder are loaded, then the "Command" is setup and all
//registered commands are setup
func main() {
	cmd.Execute()

}
