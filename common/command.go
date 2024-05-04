package common

import "github.com/gdamore/tcell/v2"

// TODO find out a common func for this interface
type Command interface{}

type InsertChar struct {
	Char rune
}

type InsertNewLine struct{}

type DelChar struct{}

type MovCurDown struct {
	Sel bool
}

type MovCurUp struct {
	Sel bool
}

type MovCurRight struct {
	Sel bool
}

type MovCurLeft struct {
	Sel bool
}

type Save struct{}

func GetCommand(e *tcell.EventKey) Command {
	mod, key, ch := e.Modifiers(), e.Key(), e.Rune()

	switch key {
	case tcell.KeyEnter:
		return InsertNewLine{}
	case tcell.KeyDelete, tcell.KeyBackspace2:
		return DelChar{}
	case tcell.KeyDown:
		return MovCurDown{
			Sel: mod == tcell.ModShift,
		}
	case tcell.KeyUp:
		return MovCurUp{
			Sel: mod == tcell.ModShift,
		}
	case tcell.KeyRight:
		return MovCurRight{
			Sel: mod == tcell.ModShift,
		}
	case tcell.KeyLeft:
		return MovCurLeft{
			Sel: mod == tcell.ModShift,
		}
	case tcell.KeyCtrlS:
		return Save{}

	default:
		return InsertChar{Char: ch}
	}
}
