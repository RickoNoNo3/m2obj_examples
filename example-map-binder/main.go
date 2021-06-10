package main

import (
	"fmt"

	"github.com/rickonono3/m2obj"
)

// initial map
var m = map[string]interface{}{
	"a": 1,
	"b": "2",
	"c": true,
	"d": map[string]interface{}{
		"e": "3",
	},
}

func main() {
	// new Object with param map[string]interface{}
	obj := m2obj.New(m)
	_ = obj.Set("d.f.g", 4)
	// staticize the object to map
	m2 := obj.Staticize()
	fmt.Println(m2)
}
