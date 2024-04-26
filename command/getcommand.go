package command

import "github.com/gdamore/tcell/v2"

func Get(e *tcell.EventKey) Command {
	_, key, ch := e.Modifiers(), e.Key(), e.Rune()

	switch key {
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
	case tcell.KeyCtrlS:
		return Save{}

	default:
		return InsertChar{Char: ch}
	}
}
