package cmd

import (
	"strings"

	"github.com/clbanning/mxj"
)

func xmlReplacer(file []byte, sourceXpath string, value string) []byte {

	mv, _ := mxj.NewMapXml([]byte(file))
	var path = ""
	var subkeys = make([]string, 0, 0)

	finalValue := ""

	for i, k := range strings.Split(sourceXpath, "/") {
		if k == "" {
			continue
		}
		idxProp := strings.Index(k, "[")
		if idxProp > -1 {
			prop := k[idxProp+1 : len(k)-1]
			k = k[:idxProp]
			//add the - for attributes
			prop = "-" + strings.Replace(prop, "=", ":", -1)
			subkeys = append(subkeys, prop)
		}
		finalValue = k
		if i > 0 {
			k = "*"
		}
		path += k + "."
	}

	finalValue += ":" + value
	path = path[:len(path)-1]
	mv.UpdateValuesForPath(finalValue, path, subkeys...)
	data, _ := mv.XmlIndent("", "    ")
	return data
}
