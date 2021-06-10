package main

import (
	"github.com/rickonono3/m2obj"
	"github.com/rickonono3/m2obj/m2json"
)

const (
	LevelInfo = iota
	LevelWarn
	LevelError
)

var Config *m2obj.Object
var FileSyncer *m2obj.FileSyncer

func init() {
	Config = m2obj.New(m2obj.Group{
		"Debug": m2obj.Group{
			"IsDebugging": true,
			"Level":       LevelWarn,
		},
	})
	// FileSyncer
	FileSyncer = m2obj.NewFileSyncer("./config.json", m2json.Formatter{})
	FileSyncer.BindObject(Config)
	// FileSyncer.Load()

	// DEFAULT FILE_SYNCER OPTIONS:
	//   Auto Saving  : On bound object changes
	//   Auto Loading : Never
	//   Hard Load    : False
}
