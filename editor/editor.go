package editor

import (
	"github.com/lvhungdev/text/command"
)

type point struct {
	row int
	col int
}

type Editor struct {
	content  [][]rune
	cur      point
	selBegin point
	selEnd   point
}

func New(content [][]rune) Editor {
	if content == nil || len(content) == 0 {
		content = [][]rune{{}}
	}

	return Editor{
		content: content,
	}
}

func (e *Editor) Content() [][]rune {
	return e.content
}

func (e *Editor) Cursor() (int, int) {
	return e.cur.row, e.cur.col
}

func (e *Editor) Selection() (int, int, int, int) {
	return e.selBegin.row, e.selBegin.col, e.selEnd.row, e.selEnd.col
}

func (e *Editor) HandleCommand(cmd command.Command) error {
	switch cmd := cmd.(type) {
	case command.InsertChar:
		e.insertChar(cmd.Char)
	case command.InsertNewLine:
		e.insertNewLine()
	case command.DelChar:
		e.deleteChar()
	case command.MovCurDown:
		e.movCurDown(cmd.Sel)
	case command.MovCurUp:
		e.movCurUp(cmd.Sel)
	case command.MovCurRight:
		e.movCurRight(cmd.Sel)
	case command.MovCurLeft:
		e.movCurLeft(cmd.Sel)
	}

	return nil
}
