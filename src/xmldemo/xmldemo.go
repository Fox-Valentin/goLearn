package xmldemo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func getAttrVale(attr []xml.Attr, name string) string {
	for _, v := range attr {
		if v.Name.Local == name {
			return v.Value
		}
	}
	return ""
}
func Readxml() {
	content, err := ioutil.ReadFile("./xmldemo/xmldemo.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	var t xml.Token
	var isBd bool
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if isBd {
				if name == "message" {
					fmt.Println(getAttrVale(token.Attr, "Include"))
				}
			} else {
				if name == "body" {
					isBd = true
				}
			}
		case xml.EndElement:
			if isBd {
				if token.Name.Local == "body" {
					isBd = false
				}
			}
		}
	}
}
