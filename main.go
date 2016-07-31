package main

import (
    	"encoding/xml"
	"strings"
	"fmt"
)

type PlistArray struct {
	Integer []int  `xml:"integer"`
}

const in = "<key>KEY1</key><string>VALUE OF KEY1</string><key>KEY2</key><string>VALUE OF KEY2</string><key>KEY3</key><integer>42</integer><key>KEY3</key><array><integer>1</integer><integer>2</integer></array>"

func main() {
	result := map[string]interface{}{}
	dec := xml.NewDecoder(strings.NewReader(in))
	dec.Strict = false
	var workingKey string

	for {
		token, _ := dec.Token()
		if token == nil {
			break
		}
		switch start := token.(type) {
		case xml.StartElement:
			fmt.Printf("startElement = %+v\n", start)
			switch start.Name.Local {
			case "key":
				var k string
				err := dec.DecodeElement(&k, &start)
				if err != nil {
					fmt.Println(err.Error())
				}
				workingKey = k
			case "string":
				var s string
				err := dec.DecodeElement(&s, &start)
				if err != nil {
					fmt.Println(err.Error())
				}
				result[workingKey] = s
				workingKey = ""
			case "integer":
				var i int
				err := dec.DecodeElement(&i, &start)
				if err != nil {
					fmt.Println(err.Error())
				}
				result[workingKey] = i
				workingKey = ""
			case "array":
				var ai PlistArray
				err := dec.DecodeElement(&ai, &start)
				if err != nil {
					fmt.Println(err.Error())
				}
				result[workingKey] = ai
				workingKey = ""
			default:
				fmt.Errorf("Unrecognized token")
			}
		}
	}
	fmt.Printf("%+v", result)

}

