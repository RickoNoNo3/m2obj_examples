package main

import (
	"fmt"
)

// Print str if IsDebugging && level >= DebugLevel
func debugPrint(str string, level int) {
	debug := Config.MustGet("Debug")
	if debug.MustGet("IsDebugging").ValBool() {
		if level >= debug.MustGet("Level").ValInt() {
			fmt.Println(">> " + str)
		}
	}
}

func main() {
	debugPrint("This is Info1", LevelInfo)   // filtered
	debugPrint("This is Warn1", LevelWarn)   // printed
	debugPrint("This is Error1", LevelError) // printed

	fmt.Println("----------")
	_ = Config.Set("Debug.Level", LevelError)

	debugPrint("This is Info2", LevelInfo)   // filtered
	debugPrint("This is Warn2", LevelWarn)   // filtered
	debugPrint("This is Error2", LevelError) // printed
}
