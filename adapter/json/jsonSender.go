package json

import "fmt"

type JsonDoc struct {
}

func (jd JsonDoc) ConvertToXML() {
	fmt.Println("<xml></xml>")
}

type JsonXMLAdapter struct {
	jsonDoc *JsonDoc
}

func (adapter JsonXMLAdapter) SendXmlData() {
	adapter.jsonDoc.ConvertToXML()
	fmt.Println("Sendind XML")
}
