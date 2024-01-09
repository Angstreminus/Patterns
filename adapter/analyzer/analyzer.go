package analyzer

import "fmt"

type CustomerAnalyzer interface {
	SendXmlData()
}

type XMLDoc struct {
}

func (doc XMLDoc) SendXmlData() {
	fmt.Println("XML Sending")
}
