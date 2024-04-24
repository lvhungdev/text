package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lvhungdev/text/command"
)

type Editor struct {
	screen tcell.Screen

	data [][]rune
	cX   int
	cY   int
}

func New() Editor {
	return Editor{
		data: [][]rune{
			{},
		},
	}
}

func (e *Editor) Init() error {
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := s.Init(); err != nil {
		return err
	}

	e.screen = s
	defer e.screen.Fini()

	e.screen.Clear()

	for {
		ev := e.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return nil
			}
			e.handleCommand(command.Get(ev))
		}

		e.draw()
	}
}

func (e *Editor) draw() {
	x, y := 0, 0

	for _, line := range e.data {
		for _, char := range line {
			e.screen.SetContent(x, y, char, nil, tcell.StyleDefault)
			x++
		}
		x = 0
		y++
	}

	e.screen.ShowCursor(e.cX, e.cY)
	e.screen.Show()
}

func (e *Editor) getLastLine() *[]rune {
	return &(e.data[len(e.data)-1])
}
