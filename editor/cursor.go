package editor

func (e *Editor) movCurDown() {
	if e.cursorRow == len(e.content)-1 {
		return
	}

	e.cursorRow++

	if e.cursorCol > len(e.content[e.cursorRow]) {
		e.cursorCol = len(e.content[e.cursorRow])
	}

	e.syncOffset()
}

func (e *Editor) movCurUp() {
	if e.cursorRow == 0 {
		return
	}

	e.cursorRow--

	if e.cursorCol > len(e.content[e.cursorRow]) {
		e.cursorCol = len(e.content[e.cursorRow])
	}

	e.syncOffset()
}

func (e *Editor) movCurRight() {
	if e.cursorCol == len(e.content[e.cursorRow]) {
		return
	}
	e.cursorCol++
	e.syncOffset()
}

func (e *Editor) movCurLeft() {
	if e.cursorCol == 0 {
		return
	}
	e.cursorCol--
	e.syncOffset()
}
