package main

import (
	"github.com/lvhungdev/text/app"
)

func main() {
	a, err := app.New("./app/app.go")
	if err != nil {
		panic(err)
	}

	defer a.Exit()

	if err := a.Start(); err != nil {
		panic(err)
	}
}
