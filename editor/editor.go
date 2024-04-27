package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lvhungdev/text/command"
)

type Editor struct {
	screen tcell.Screen
	path   string

	content      [][]rune
	cursorRow    int
	cursorCol    int
	offsetRow    int
	offsetCol    int
	padRow       int
	padCol       int
	screenWidth  int
	screenHeight int
}

func New() Editor {
	return Editor{
		content: [][]rune{{}},
	}
}

func (e *Editor) Start() error {
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := s.Init(); err != nil {
		return err
	}

	e.screen = s
	defer e.screen.Fini()

	for {
		ev := e.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			e.syncScreenSize()

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return nil
			}
			// TODO handle error
			_ = e.handleCommand(command.Get(ev))
		}

		// TODO handle error
		_ = e.draw()
	}
}

func (e *Editor) syncScreenSize() {
	e.screen.Sync()

	w, h := e.screen.Size()
	e.screenWidth = w - e.padCol
	e.screenHeight = h - e.padRow
}

func (e *Editor) syncOffset() {
	diffRow := e.cursorRow - e.offsetRow
	diffCol := e.cursorCol - e.offsetCol

	if diffCol < 0 {
		e.offsetCol = e.cursorCol
	} else if diffCol >= e.screenWidth {
		e.offsetCol = e.cursorCol - e.screenWidth + 1
	}

	if diffRow < 0 {
		e.offsetRow = e.cursorRow
	} else if diffRow >= e.screenHeight {
		e.offsetRow = e.cursorRow - e.screenHeight + 1
	}
}

func (e *Editor) handleCommand(cmd command.Command) error {
	switch cmd := cmd.(type) {
	case command.InsertChar:
		e.insertChar(cmd.Char)
	case command.InsertNewLine:
		e.insertNewLine()
	case command.DelChar:
		e.deleteChar()
	case command.MovCurDown:
		e.movCurDown()
	case command.MovCurUp:
		e.movCurUp()
	case command.MovCurRight:
		e.movCurRight()
	case command.MovCurLeft:
		e.movCurLeft()
	case command.Save:
		return e.save()
	}

	return nil
}
