package editor

import (
	c "github.com/lvhungdev/text/common"
)

func (e *Editor) InsertChar(char rune) {
	e.content[e.cur.Row] = c.SliceInsertAt(e.content[e.cur.Row], e.cur.Col, char)
	e.cur.Col++
}

func (e *Editor) InsertNewLine() {
	if e.cur.Col == 0 {
		e.content = c.SliceInsertAt(e.content, e.cur.Row, []rune{})
	} else if e.cur.Col == len(e.content[e.cur.Row]) {
		e.content = c.SliceInsertAt(e.content, e.cur.Row+1, []rune{})
	} else {
		current := e.content[e.cur.Row][:e.cur.Col]
		new := e.content[e.cur.Row][e.cur.Col:]

		e.content[e.cur.Row] = current
		e.content = c.SliceInsertAt(e.content, e.cur.Row+1, new)
	}

	e.cur.Row++
	e.cur.Col = 0
}

func (e *Editor) DeleteChar() {
	if e.cur.Col == 0 && e.cur.Row == 0 {
		return
	}

	if e.cur.Col == 0 {
		line := e.content[e.cur.Row]

		e.content = c.SliceRemoveAt(e.content, e.cur.Row)
		e.cur.Row--
		e.cur.Col = len(e.content[e.cur.Row])

		e.content[e.cur.Row] = append(e.content[e.cur.Row], line...)
	} else {
		e.content[e.cur.Row] = c.SliceRemoveAt(e.content[e.cur.Row], e.cur.Col-1)
		e.cur.Col--
	}
}
