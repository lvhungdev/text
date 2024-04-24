package main

import "github.com/lvhungdev/text/editor"

func main() {
	e := editor.New()
	err := e.Init()
	if err != nil {
		panic(err)
	}
}
