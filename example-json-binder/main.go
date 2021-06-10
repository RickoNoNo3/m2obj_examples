package main

import (
	"fmt"

	"github.com/rickonono3/m2obj/m2json"
)

// initial json string
var jsonStr = `{
 "a": 1,
 "b": "2",
 "c": true,
 "d": {
   "e": "3"
 }
}`

func main() {
	// new Formatter
	formatter := m2json.Formatter{}
	// convert jsonStr([]byte) to a new Object
	obj, err := formatter.Unmarshal([]byte(jsonStr))
	if err == nil {
		_ = obj.Set("d.f.g", 4)
		// convert Object to jsonStr2([]byte)
		jsonStr2, _ := formatter.Marshal(obj)
		fmt.Println(string(jsonStr2))
	}
}
