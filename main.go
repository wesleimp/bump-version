package main

import (
	"os"

	"github.com/wesleimp/bump-version/cmd"
)

var version = "v0.1.0"

func main() {
	err := cmd.Execute(version, os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
