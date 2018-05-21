package cmd

import (
	"github.com/beevik/etree"
)

//xmlReplacer - Used to replace the contents in an XMLby XPath. The method takes the below parameters:
// 1. XML file content in bytes
// 2. Values to be replaced, identified by XPath
// 3. Value to be replaced
// Returns - The replaced XML in bytes
//Example:
// XPATH: beans/bean[id=mqQueueConnectionFactory]/property[name=brokerURL]/value
// Candidate:
// a. Search for beans root tag
// b. Followed by a child bean tag with id=mqQueueConnectionFactory
// c. From within the bean tag, search for a property child tag with attribute of name=brokerURL
// d. Finally a child of value which needs to be replaced
// This uses the "github.com/beevik/etree" API to identify and replace the elements
func xmlReplacer(file []byte, sourceXpath string, value string) []byte {

	doc := etree.NewDocument()
	err := doc.ReadFromBytes(file)
	if err != nil {
		panic(err)
	}

	for _, elem := range doc.FindElements(sourceXpath) {
		elem.SetText(value)
	}

	data, _ := doc.WriteToBytes()

	return data
}
