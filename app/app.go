package app

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
	c "github.com/lvhungdev/text/common"
	"github.com/lvhungdev/text/editor"
	"github.com/lvhungdev/text/render"
)

type App struct {
	screen   tcell.Screen
	renderer *render.Renderer
	editor   *editor.Editor
	path     string
}

func New(path string) (*App, error) {
	scr, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, errors.New("path is empty")
	}

	data, err := readFile(path)
	if err != nil {
		return nil, err
	}

	e := editor.New(data)

	if err := scr.Init(); err != nil {
		return nil, err
	}

	w, h := scr.Size()
	renderer := render.NewRenderer(scr, &e, render.NewRegion(0, 0, w, h))

	return &App{
		screen:   scr,
		renderer: &renderer,
		editor:   &e,
		path:     path,
	}, nil
}

func (a *App) Start() error {
	for {
		ev := a.screen.PollEvent()

		switch ev := ev.(type) {
		// TODO handle event resize
		// case *tcell.EventResize:
		// 	e.syncScreenSize()

		case *tcell.EventKey:
			mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()

			if key == tcell.KeyCtrlQ {
				return nil
			}

			err := a.HandleCommand(c.GetCommand(mod, key, ch))
			if err != nil {
				return err
			}
		}

		a.renderer.Render()
	}
}

func (a *App) Exit() {
	a.screen.Fini()
}

func (a *App) HandleCommand(cmd c.Command) error {
	switch cmd := cmd.(type) {
	case c.InsertChar:
		a.editor.InsertChar(cmd.Char)
	case c.InsertNewLine:
		a.editor.InsertNewLine()
	case c.DelChar:
		a.editor.DeleteChar()
	case c.MovCurDown:
		a.editor.MovCurDown(cmd.Sel)
	case c.MovCurUp:
		a.editor.MovCurUp(cmd.Sel)
	case c.MovCurRight:
		a.editor.MovCurRight(cmd.Sel)
	case c.MovCurLeft:
		a.editor.MovCurLeft(cmd.Sel)
	}

	return nil
}

func readFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
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
				return nil, err
			}
		}

		data = append(data, b)
	}

	content := [][]rune{}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		content = append(content, []rune(string(line)))
	}

	return content, nil
}
