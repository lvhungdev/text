package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lvhungdev/text/command"
)

func (e *Editor) handleCommand(cmd command.Command) {
	switch cmd := cmd.(type) {
	case command.InsertChar:
		e.handleCommandInsertChar(cmd.Char)
	case command.InsertNewLine:
		e.handleCommandInsertNewLine()
	case command.DeleteChar:
		e.handleCommandDeleteChar()
	}
}

func (e *Editor) handleCommandInsertChar(char rune) {
	line := e.getLastLine()
	*line = append(*line, char)
	e.cX++
}

func (e *Editor) handleCommandInsertNewLine() {
	e.data = append(e.data, []rune{})
	e.cY++
	e.cX = 0
}

func (e *Editor) handleCommandDeleteChar() {
	line := e.getLastLine()
	if len(*line) > 0 {
		*line = (*line)[:len(*line)-1]
		e.cX--

		e.screen.SetContent(e.cX, e.cY, ' ', nil, tcell.StyleDefault)
	} else {
		length := len(e.data)
		if length == 1 {
			return
		}

		e.data = e.data[:length-1]
		e.cY--
		e.cX = len(*e.getLastLine())
	}
}
