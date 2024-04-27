package main

import (
	"os"

	"github.com/lvhungdev/text/editor"
)

func main() {
	var e editor.Editor

	args := os.Args[1:]
	if len(args) >= 1 {
		v, err := editor.FromFile(args[0])
		if err != nil {
			panic(err)
		}

		e = v
	} else {
		e = editor.New()
	}

	err := e.Start()
	if err != nil {
		panic(err)
	}
}
