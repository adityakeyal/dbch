package cmd

import (
	"github.com/beevik/etree"
)

func xmlReplacer(file []byte, sourceXpath string, value string) []byte {

	doc := etree.NewDocument()
	err := doc.ReadFromBytes(file)
	if err != nil {
		panic(err)
	}

	//beans/bean[id=mqQueueConnectionFactory]/property[name=brokerURL]/value

	for _, elem := range doc.FindElements(sourceXpath) {
		elem.SetText(value)
	}

	data, _ := doc.WriteToBytes()

	return data
}
