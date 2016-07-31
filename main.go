package main

import (
    	"encoding/xml"
	"strings"
	"fmt"
)

type PlistArray struct {
	val []int  `xml:"integer"`
}

// const in = "<key>KEY1</key><string>VALUE OF KEY1</string><key>KEY2</key><string>VALUE OF KEY2</string><key>KEY3</key><integer>42</integer><key>KEY3</key><array><integer>1</integer><integer>2</integer></array>"
const in = "<key>KEY1</key><string>VALUE OF KEY1</string><key>KEY2</key><string>VALUE OF KEY2</string><key>KEY3</key><integer>42</integer>"

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
				dec.DecodeElement(&k, &start)
				workingKey = k
			case "string":
				var s string
				err := dec.DecodeElement(&s, &start)
				if err != nil {
					fmt.Println(err)
				}
				result[workingKey] = s
				fmt.Printf("adding %s:%s to result\n", workingKey, s)
				workingKey = ""
			case "integer":
				var i int
				dec.DecodeElement(&i, &start)
				fmt.Printf("adding %s:%d to result\n", workingKey, i)
				result[workingKey] = i
				workingKey = ""
			case "array":
				var ai PlistArray
				dec.DecodeElement(&ai, &start)
				fmt.Printf("adding %s:%v to result\n", workingKey, ai)
				result[workingKey] = ai
				workingKey = ""
			default:
				fmt.Errorf("Unrecognized token")
			}
		}
	}
	fmt.Printf("%+v", result)

}

