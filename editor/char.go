package editor

import (
	"github.com/lvhungdev/text/utils"
)

func (e *Editor) insertChar(char rune) {
	e.content[e.curRow] = utils.SliceInsertAt(e.content[e.curRow], e.curCol, char)
	e.curCol++
}

func (e *Editor) insertNewLine() {
	if e.curCol == 0 {
		e.content = utils.SliceInsertAt(e.content, e.curRow, []rune{})
	} else if e.curCol == len(e.content[e.curRow]) {
		e.content = utils.SliceInsertAt(e.content, e.curRow+1, []rune{})
	} else {
		current := e.content[e.curRow][:e.curCol]
		new := e.content[e.curRow][e.curCol:]

		e.content[e.curRow] = current
		e.content = utils.SliceInsertAt(e.content, e.curRow+1, new)
	}

	e.curRow++
	e.curCol = 0
}

func (e *Editor) deleteChar() {
	if e.curCol == 0 && e.curRow == 0 {
		return
	}

	if e.curCol == 0 {
		line := e.content[e.curRow]

		e.content = utils.SliceRemoveAt(e.content, e.curRow)
		e.curRow--
		e.curCol = len(e.content[e.curRow])

		e.content[e.curRow] = append(e.content[e.curRow], line...)
	} else {
		e.content[e.curRow] = utils.SliceRemoveAt(e.content[e.curRow], e.curCol-1)
		e.curCol--
	}
}
