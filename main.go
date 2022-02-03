package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/wesleimp/bump-version/cmd"
)

var version = "v0.1.0"

func main() {
	err := cmd.Execute(version, os.Args)
	if err != nil {
		color.New(color.FgRed).Println(err.Error())
		os.Exit(1)
	}
}
