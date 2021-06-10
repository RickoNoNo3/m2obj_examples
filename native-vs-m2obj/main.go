package main

import (
	"fmt"
	"github.com/rickonono3/m2obj"
)

func main() {
	// ---------------------------------------------------------
	// Without M2Obj
	// ---------------------------------------------------------
	var M = map[string]interface{}{
		"info": map[string]interface{}{
			"name": "UKP",
			"schools": []interface{}{
				"Engineering|engineering@ukp.edu|135",
				map[string]interface{}{
					"name":         "Law",
					"email":        "law@ukp.edu",
					"studentCount": 300,
					"notice":       "",
				},
			},
		},
	}

	// set student count of law school to 295
	M["info"].(map[string]interface{})["schools"].([]interface{})[1].(map[string]interface{})["studentCount"] = 295
	// delete notice of law school
	delete(M["info"].(map[string]interface{})["schools"].([]interface{})[1].(map[string]interface{}), "notice")
	// println "UKP"
	fmt.Println(M["info"].(map[string]interface{})["name"].(string))
	// println data as map
	fmt.Println(M)

	// ---------------------------------------------------------
	// With M2Obj
	// ---------------------------------------------------------
	var M2 = m2obj.New(M) // m2obj can be constructed in a lot of styles, see examples below.

	// set student count of law school to 295
	M2.Set("info.schools.[1].studentCount", 295)
	// delete notice of law school
	M2.Remove("info.schools.[1].notice")
	// println "UKP"
	fmt.Println(M2.MustGet("info.name").ValStr())
	// println data as map
	fmt.Println(M2.Staticize())
}
