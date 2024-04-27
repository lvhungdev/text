package editor

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func (e *Editor) draw() error {
	// TODO screen.Clear() should not be called in everytime
	e.screen.Clear()

	// TODO optimize this, we should not loop the whole content
	// only need to loop content with screen size
	for r := 0; r < len(e.content); r++ {
		for c := 0; c < len(e.content[r]); c++ {
			err := e.drawChar(r, c)
			if err != nil {
				return err
			}
		}
	}

	e.drawCursor()

	e.screen.Show()
	return nil
}

func (e *Editor) drawChar(row int, col int) error {
	if row < 0 || row >= len(e.content) {
		return fmt.Errorf("invalid row %v", row)
	}
	if col < 0 || col >= len(e.content[row]) {
		return fmt.Errorf("invalid col %v", col)
	}

	posX := col - e.offsetCol
	posY := row - e.offsetRow

	if posX >= 0 && posX < e.screenWidth-1 && posY >= 0 && posY < e.screenHeight-1 {
		e.screen.SetContent(posX+e.padCol, posY+e.padRow, e.content[row][col], nil, tcell.StyleDefault)
	}

	return nil
}

func (e *Editor) drawCursor() {
	posX := e.cursorCol - e.offsetCol
	posY := e.cursorRow - e.offsetRow

	if posX >= 0 && posX < e.screenWidth-1 && posY >= 0 || posY < e.screenHeight-1 {
		e.screen.ShowCursor(posX+e.padCol, posY+e.padRow)
	}

}
