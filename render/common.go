package render

import (
	"github.com/gdamore/tcell/v2"
	c "github.com/lvhungdev/text/common"
)

type editor interface {
	Content() [][]rune
	Cursor() (row, col int)
	Selection() (begin, end c.Point)
}

type region struct {
	row    int
	col    int
	width  int
	height int
}

func renderChar(screen tcell.Screen, region region, row, col int, ch rune, style tcell.Style) {
	if row < 0 || row >= region.height || col < 0 || col >= region.width {
		return
	}

	x := col + region.col
	y := row + region.row

	screen.SetContent(x, y, ch, nil, style)
}

func renderChars(screen tcell.Screen, region region, row, col int, chs []rune, style tcell.Style) {
	for i, ch := range chs {
		renderChar(screen, region, row, col+i, ch, style)
	}
}
