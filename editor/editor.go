package editor

import (
	c "github.com/lvhungdev/text/common"
)

type Editor struct {
	content  [][]rune
	cur      c.Point
	selBegin c.Point
	selEnd   c.Point
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
	return e.cur.Row, e.cur.Col
}

func (e *Editor) Selection() (c.Point, c.Point) {
	return e.selBegin, e.selEnd
}

func (e *Editor) HandleCommand(cmd c.Command) error {
	switch cmd := cmd.(type) {
	case c.InsertChar:
		e.insertChar(cmd.Char)
	case c.InsertNewLine:
		e.insertNewLine()
	case c.DelChar:
		e.deleteChar()
	case c.MovCurDown:
		e.movCurDown(cmd.Sel)
	case c.MovCurUp:
		e.movCurUp(cmd.Sel)
	case c.MovCurRight:
		e.movCurRight(cmd.Sel)
	case c.MovCurLeft:
		e.movCurLeft(cmd.Sel)
	}

	return nil
}
