package editor

func (e *Editor) MovCurDown(sel bool) {
	if e.cur.Row == len(e.content)-1 {
		return
	}

	e.cur.Row++

	if e.cur.Col > len(e.content[e.cur.Row]) {
		e.cur.Col = len(e.content[e.cur.Row])
	}

	e.handleSel(sel)
}

func (e *Editor) MovCurUp(sel bool) {
	if e.cur.Row == 0 {
		return
	}

	e.cur.Row--

	if e.cur.Col > len(e.content[e.cur.Row]) {
		e.cur.Col = len(e.content[e.cur.Row])
	}

	e.handleSel(sel)
}

func (e *Editor) MovCurRight(sel bool) {
	if e.cur.Col == len(e.content[e.cur.Row]) {
		return
	}
	e.cur.Col++

	e.handleSel(sel)
}

func (e *Editor) MovCurLeft(sel bool) {
	if e.cur.Col == 0 {
		return
	}
	e.cur.Col--

	e.handleSel(sel)
}

func (e *Editor) handleSel(sel bool) {
	e.selEnd = e.cur
	if !sel {
		e.selBegin = e.cur
	}
}
