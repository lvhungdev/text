package editor

func (e *Editor) movCurDown() {
	if e.curRow == len(e.content)-1 {
		return
	}

	e.curRow++

	if e.curCol > len(e.content[e.curRow]) {
		e.curCol = len(e.content[e.curRow])
	}
}

func (e *Editor) movCurUp() {
	if e.curRow == 0 {
		return
	}

	e.curRow--

	if e.curCol > len(e.content[e.curRow]) {
		e.curCol = len(e.content[e.curRow])
	}
}

func (e *Editor) movCurRight() {
	if e.curCol == len(e.content[e.curRow]) {
		return
	}
	e.curCol++
}

func (e *Editor) movCurLeft() {
	if e.curCol == 0 {
		return
	}
	e.curCol--
}
