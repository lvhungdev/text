package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
	c "github.com/lvhungdev/text/common"
	"github.com/lvhungdev/text/editor"
	"github.com/lvhungdev/text/render"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	}

	defer screen.Fini()

	e, err := fromFile("./cmd/main.go")
	if err != nil {
		panic(err)
	}

	// w, h := screen.Size()
	region := render.NewRegion(2, 4, 40, 20)

	r := render.NewRenderer(screen, &e, region)

	for {
		ev := screen.PollEvent()

		switch ev := ev.(type) {
		// TODO handle event resize
		// case *tcell.EventResize:
		// 	e.syncScreenSize()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlQ {
				return
			} else if ev.Modifiers() == tcell.ModCtrl {
				continue
			} else if ev.Modifiers() == tcell.ModAlt && ev.Key() == tcell.KeyLeft {
				fmt.Println("asd")
				continue
			}

			// TODO handle error
			_ = e.HandleCommand(c.GetCommand(ev))
		}

		r.Render()
	}
}

func fromFile(path string) (editor.Editor, error) {
	file, err := os.Open(path)
	if err != nil {
		return editor.New(nil), nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var data []byte
	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return editor.New(nil), err
			}
		}

		data = append(data, b)
	}

	content := [][]rune{}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		content = append(content, []rune(string(line)))
	}

	return editor.New(content), nil
}
