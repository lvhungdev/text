package editor

func (e *Editor) movCurDown() {
	if e.cY == len(e.data)-1 {
		return
	}

	e.cY++

	if e.cX > len(e.data[e.cY]) {
		e.cX = len(e.data[e.cY])
	}

	e.syncOffset()
}

func (e *Editor) movCurUp() {
	if e.cY == 0 {
		return
	}

	e.cY--

	if e.cX > len(e.data[e.cY]) {
		e.cX = len(e.data[e.cY])
	}

	e.syncOffset()
}

func (e *Editor) movCurRight() {
	if e.cX == len(e.data[e.cY]) {
		return
	}
	e.cX++
	e.syncOffset()
}

func (e *Editor) movCurLeft() {
	if e.cX == 0 {
		return
	}
	e.cX--
	e.syncOffset()
}
