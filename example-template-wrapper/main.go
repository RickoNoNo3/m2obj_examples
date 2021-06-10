package main

import (
	"github.com/rickonono3/m2obj"
	"html/template"
	"os"
)

// Config is the global configuration
var Config = m2obj.New(m2obj.Group{
	"cdn": "https://example.com",
})

func main() {
	// new template
	t, err := template.ParseFiles("index.gohtml")
	if err == nil {
		// define data of the executed template
		data := m2obj.New(m2obj.Group{
			"title": "M2Obj Examples",
			"body": m2obj.Group{
				"h1":   "M2Obj Examples for Go Template Data Wrapper",
				"text": "Enjoy!",
			},
		})
		// add the config object to data of the executed template.
		data.Set("config", Config)
		// staticize the data
		_ = t.Execute(os.Stdout, data.Staticize())
	}
}
