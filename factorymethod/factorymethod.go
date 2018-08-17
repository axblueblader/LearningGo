package main

import (
	"fmt"
)

type Document interface {
	Show()
	New() Document
}

type XMLDocument struct {
	xmlData string
}

func (xmlDoc XMLDocument) Show() {
	fmt.Println(xmlDoc.xmlData)
}

func (xmlDoc XMLDocument) New() Document {
	newDoc := XMLDocument{xmlData: "New XML doc"}
	return newDoc
}

type JSONDocument struct {
	jsonData string
}

func (jsonDoc JSONDocument) Show() {
	fmt.Println(jsonDoc.jsonData)
}

func (jsonDoc JSONDocument) New() Document {
	newDoc := JSONDocument{jsonData: "New JSON doc"}
	return newDoc
}

type IApllication interface {
	OpenFile(fileType string) Document
}

type CoreApplication struct {
	DocFactory map[string]Document
}

func (coreApp CoreApplication) OpenFile(fileType string) Document {
	return coreApp.DocFactory[fileType].New()
}

func (coreApp *CoreApplication) AddDocType(typeName string, doc Document) {
	if coreApp.DocFactory != nil {
		coreApp.DocFactory[typeName] = doc
	} else {
		coreApp.DocFactory = make(map[string]Document)
		coreApp.DocFactory[typeName] = doc
	}
}

func main() {
	fmt.Println("Factory Method Demo")
	var coreApp CoreApplication
	coreApp.AddDocType("XML", XMLDocument{})
	coreApp.AddDocType("JSON", JSONDocument{})
	newDoc := coreApp.OpenFile("XML")
	newDoc.Show()
	newDoc = coreApp.OpenFile("JSON")
	newDoc.Show()
}
