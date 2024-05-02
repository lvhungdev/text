package editor

import (
	"github.com/lvhungdev/text/command"
)

type Editor struct {
	content [][]rune
	curRow  int
	curCol  int
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
	return e.curRow, e.curCol
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
		e.movCurDown()
	case command.MovCurUp:
		e.movCurUp()
	case command.MovCurRight:
		e.movCurRight()
	case command.MovCurLeft:
		e.movCurLeft()
	}

	return nil
}
