package cmd

import (
	"strings"
)

//propReplacer - Used to replace the values of specific key in property files
//The method takes the below parameters:
// 1. file content in bytes
// 2. key to be replaced
// 3. value to be replaced
// Returns the file in bytes
// Example:
// The method can be invoked by calling it as below:
// Sample File:
//    key1=value1
//    key2=value2
//    key3=value3
//  source : key2
//  value : val2
// Result:
//    key1=value1
//    key2=val2
//    key3=value3
///////////////////////////////////////////////////////////////////////
func propReplacer(file []byte, source string, value string) []byte {

	filecontent := string(file)

	newfilecontent := ""

	lines := strings.Split(filecontent, "\n")
	for _, val := range lines {

		if strings.Index(val, "=") == -1 {
			newfilecontent += val + "\n"
		} else {

			keyVals := strings.Split(val, "=")
			key := strings.TrimSpace(keyVals[0])

			if key == source {
				newfilecontent += key + "=" + value + "\n"
			} else {
				newfilecontent += val + "\n"
			}
		}

	}

	return []byte(newfilecontent)
}
