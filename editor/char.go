package editor

import (
	"github.com/lvhungdev/text/utils"
)

func (e *Editor) insertChar(char rune) {
	e.data[e.cY] = utils.SliceInsertAt(e.data[e.cY], e.cX, char)
	e.cX++
	e.syncOffset()
}

func (e *Editor) insertNewLine() {
	var index int

	if e.cX == 0 {
		index = e.cY
		e.data = utils.SliceInsertAt(e.data, index, []rune{})
	} else if e.cX == len(e.data[e.cY]) {
		index = e.cY + 1
		e.data = utils.SliceInsertAt(e.data, index, []rune{})
	} else {
		index = e.cY
		current := e.data[e.cY][:e.cX]
		new := e.data[e.cY][e.cX:]

		e.data[e.cY] = current
		e.data = utils.SliceInsertAt(e.data, index+1, new)
	}

	w, _ := e.screen.Size()
	for i := index; i < len(e.data) && i < w; i++ {
		e.clearLine(i)
	}

	e.cY++
	e.cX = 0
	e.syncOffset()
}

func (e *Editor) deleteChar() {
	if e.cX == 0 && e.cY == 0 {
		return
	}

	if e.cX == 0 {
		line := e.data[e.cY]

		e.data = utils.SliceRemoveAt(e.data, e.cY)
		e.cY--
		e.cX = len(e.data[e.cY])

		e.data[e.cY] = append(e.data[e.cY], line...)

		w, _ := e.screen.Size()
		for i := e.cY; i <= len(e.data) && i < w; i++ {
			e.clearLine(i)
		}
	} else {
		e.data[e.cY] = utils.SliceRemoveAt(e.data[e.cY], e.cX-1)
		e.cX--
		e.clearLine(e.cY)
	}

	e.syncOffset()
}
