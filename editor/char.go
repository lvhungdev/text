package editor

import (
	"github.com/lvhungdev/text/utils"
)

func (e *Editor) insertChar(char rune) {
	e.content[e.cursorRow] = utils.SliceInsertAt(e.content[e.cursorRow], e.cursorCol, char)
	e.cursorCol++
	e.syncOffset()
}

func (e *Editor) insertNewLine() {
	if e.cursorCol == 0 {
		e.content = utils.SliceInsertAt(e.content, e.cursorRow, []rune{})
	} else if e.cursorCol == len(e.content[e.cursorRow]) {
		e.content = utils.SliceInsertAt(e.content, e.cursorRow+1, []rune{})
	} else {
		current := e.content[e.cursorRow][:e.cursorCol]
		new := e.content[e.cursorRow][e.cursorCol:]

		e.content[e.cursorRow] = current
		e.content = utils.SliceInsertAt(e.content, e.cursorRow+1, new)
	}

	e.cursorRow++
	e.cursorCol = 0
	e.syncOffset()
}

func (e *Editor) deleteChar() {
	if e.cursorCol == 0 && e.cursorRow == 0 {
		return
	}

	if e.cursorCol == 0 {
		line := e.content[e.cursorRow]

		e.content = utils.SliceRemoveAt(e.content, e.cursorRow)
		e.cursorRow--
		e.cursorCol = len(e.content[e.cursorRow])

		e.content[e.cursorRow] = append(e.content[e.cursorRow], line...)
	} else {
		e.content[e.cursorRow] = utils.SliceRemoveAt(e.content[e.cursorRow], e.cursorCol-1)
		e.cursorCol--
	}

	e.syncOffset()
}
