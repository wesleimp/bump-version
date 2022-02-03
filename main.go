package main

import (
	"log"
	"os"

	"github.com/wesleimp/bump-version/cmd"
)

var version = "v0.1.0"

func main() {
	err := cmd.Execute(version, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
