package command

import "github.com/gdamore/tcell/v2"

func Get(e *tcell.EventKey) Command {
	switch e.Key() {
	case tcell.KeyEnter:
		return InsertNewLine{}

	case tcell.KeyDelete, tcell.KeyBackspace2:
		return DeleteChar{}

	default:
		return InsertChar{Char: e.Rune()}
	}
}
