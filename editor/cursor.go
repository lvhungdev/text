package editor

func (e *Editor) movCurDown(sel bool) {
	if e.cur.row == len(e.content)-1 {
		return
	}

	e.cur.row++

	if e.cur.col > len(e.content[e.cur.row]) {
		e.cur.col = len(e.content[e.cur.row])
	}

	e.handleSel(sel)
}

func (e *Editor) movCurUp(sel bool) {
	if e.cur.row == 0 {
		return
	}

	e.cur.row--

	if e.cur.col > len(e.content[e.cur.row]) {
		e.cur.col = len(e.content[e.cur.row])
	}

	e.handleSel(sel)
}

func (e *Editor) movCurRight(sel bool) {
	if e.cur.col == len(e.content[e.cur.row]) {
		return
	}
	e.cur.col++

	e.handleSel(sel)
}

func (e *Editor) movCurLeft(sel bool) {
	if e.cur.col == 0 {
		return
	}
	e.cur.col--

	e.handleSel(sel)
}

func (e *Editor) handleSel(sel bool) {
	e.selEnd = e.cur
	if !sel {
		e.selBegin = e.cur
	}
}
