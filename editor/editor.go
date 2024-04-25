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
	oX   int
	oY   int
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

// TODO optimize this func
// we should only draw within screen size, not necessarily the whole data
func (e *Editor) draw() {
	for y := e.oY; y < len(e.data); y++ {
		for x := e.oX; x < len(e.data[y]); x++ {
			e.screen.SetContent(x-e.oX, y-e.oY, e.data[y][x], nil, tcell.StyleDefault)
		}
	}

	e.screen.ShowCursor(e.cX-e.oX, e.cY-e.oY)
	e.screen.Show()
}

func (e *Editor) handleCommand(cmd command.Command) {
	switch cmd := cmd.(type) {
	case command.InsertChar:
		e.insertChar(cmd.Char)
	case command.InsertNewLine:
		e.insertNewLine()
	case command.DelChar:
		e.deleteChar()
	case command.MovCurDown:
		e.MovCurDown()
	case command.MovCurUp:
		e.MovCurUp()
	case command.MovCurRight:
		e.MovCurRight()
	case command.MovCurLeft:
		e.MovCurLeft()
	}
}

func (e *Editor) clearLine(index int) {
	width, _ := e.screen.Size()
	for i := 0; i < width; i++ {
		e.screen.SetContent(i, index, ' ', nil, tcell.StyleDefault)
	}
}
