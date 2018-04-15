package cmd

import (
	"strings"
)

func propReplacer(file []byte, source string, value string) []byte {

	filecontent := string(file)

	newfilecontent := ""

	lines := strings.Split(filecontent, "\n")
	for _, val := range lines {

		if strings.Index(val, "=") == -1 {
			newfilecontent += val + "\n"
			//continue
		} else {

			keyVals := strings.Split(val, "=")
			key := strings.TrimSpace(keyVals[0])

			if key == source {
				//fmt.Println("Found the key with value " + keyVals[1])
				newfilecontent += key + "=" + value + "\n"
			} else {
				newfilecontent += val + "\n"
			}
		}

	}

	return []byte(newfilecontent)
}
