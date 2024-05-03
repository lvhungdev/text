package renderer

import "github.com/gdamore/tcell/v2"

type editor interface {
	Content() [][]rune
	Cursor() (row, col int)
	Selection() (beginRow, beginCol, endRow, endCol int)
}

type point struct {
	row int
	col int
}

func (p point) Compare(p1 point) int {
	if p.row > p1.row {
		return 1
	}
	if p.row < p1.row {
		return -1
	}
	if p.col > p1.col {
		return 1
	}
	if p.col < p1.col {
		return -1
	}
	return 0
}

type Region struct {
	row    int
	col    int
	width  int
	height int
}

func NewRegion(row, col, width, height int) Region {
	return Region{
		row:    row,
		col:    col,
		width:  width,
		height: height,
	}
}

func renderChar(screen tcell.Screen, region Region, row, col int, ch rune, style tcell.Style) {
	if row < 0 || row >= region.height ||
		col < 0 || col >= region.width {
		return
	}

	x := col + region.col
	y := row + region.row

	screen.SetContent(x, y, ch, nil, style)
}

func renderChars(screen tcell.Screen, region Region, row, col int, chs []rune, style tcell.Style) {
	for i, ch := range chs {
		renderChar(screen, region, row, col+i, ch, style)
	}
}
