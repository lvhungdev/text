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
