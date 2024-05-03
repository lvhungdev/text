package editor

import (
	"github.com/lvhungdev/text/utils"
)

func (e *Editor) insertChar(char rune) {
	e.content[e.cur.row] = utils.SliceInsertAt(e.content[e.cur.row], e.cur.col, char)
	e.cur.col++
}

func (e *Editor) insertNewLine() {
	if e.cur.col == 0 {
		e.content = utils.SliceInsertAt(e.content, e.cur.row, []rune{})
	} else if e.cur.col == len(e.content[e.cur.row]) {
		e.content = utils.SliceInsertAt(e.content, e.cur.row+1, []rune{})
	} else {
		current := e.content[e.cur.row][:e.cur.col]
		new := e.content[e.cur.row][e.cur.col:]

		e.content[e.cur.row] = current
		e.content = utils.SliceInsertAt(e.content, e.cur.row+1, new)
	}

	e.cur.row++
	e.cur.col = 0
}

func (e *Editor) deleteChar() {
	if e.cur.col == 0 && e.cur.row == 0 {
		return
	}

	if e.cur.col == 0 {
		line := e.content[e.cur.row]

		e.content = utils.SliceRemoveAt(e.content, e.cur.row)
		e.cur.row--
		e.cur.col = len(e.content[e.cur.row])

		e.content[e.cur.row] = append(e.content[e.cur.row], line...)
	} else {
		e.content[e.cur.row] = utils.SliceRemoveAt(e.content[e.cur.row], e.cur.col-1)
		e.cur.col--
	}
}
