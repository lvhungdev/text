package editor

func (e *Editor) MovCurDown() {
	if e.cY == len(e.data)-1 {
		return
	}

	e.cY++

	if e.cX > len(e.data[e.cY]) {
		e.cX = len(e.data[e.cY])
	}
}

func (e *Editor) MovCurUp() {
	if e.cY == 0 {
		return
	}

	e.cY--

	if e.cX > len(e.data[e.cY]) {
		e.cX = len(e.data[e.cY])
	}
}

func (e *Editor) MovCurRight() {
	if e.cX == len(e.data[e.cY]) {
		return
	}
	e.cX++
}

func (e *Editor) MovCurLeft() {
	if e.cX == 0 {
		return
	}
	e.cX--
}
