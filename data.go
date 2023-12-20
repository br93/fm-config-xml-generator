package main

import "encoding/xml"

type Record struct {
	XMLName  xml.Name         `xml:"record"`
	Booleans []BooleanElement `xml:"boolean"`
	List     List             `xml:"list"`
}

type BooleanElement struct {
	XMLName xml.Name `xml:"boolean"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type Map struct {
	From string `xml:"from,attr"`
	To   string `xml:"to,attr"`
}

type List struct {
	XMLName xml.Name `xml:"list"`
	ID      string   `xml:"id,attr"`
	Maps    []Map    `xml:"record"`
}
