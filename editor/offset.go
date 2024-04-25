package editor

func (e *Editor) syncOffset() {
	w, h := e.screen.Size()
	diffX := e.cX - e.oX
	diffY := e.cY - e.oY

	if diffX < 0 {
		e.oX = e.cX
		e.screen.Clear()
	} else if diffX >= w {
		e.oX = e.cX - w + 1
		e.screen.Clear()
	}

	if diffY < 0 {
		e.oY = e.cY
		e.screen.Clear()
	} else if diffY >= h {
		e.oY = e.cY - h + 1
		e.screen.Clear()
	}
}
