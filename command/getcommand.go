package command

import "github.com/gdamore/tcell/v2"

func Get(e *tcell.EventKey) Command {
	switch e.Key() {
	case tcell.KeyEnter:
		return InsertNewLine{}
	case tcell.KeyDelete, tcell.KeyBackspace2:
		return DelChar{}
	case tcell.KeyDown:
		return MovCurDown{}
	case tcell.KeyUp:
		return MovCurUp{}
	case tcell.KeyRight:
		return MovCurRight{}
	case tcell.KeyLeft:
		return MovCurLeft{}

	default:
		return InsertChar{Char: e.Rune()}
	}
}
