package renderer

import (
	"github.com/gdamore/tcell/v2"
)

type point struct {
	row int
	col int
}

type editor interface {
	Content() [][]rune
	Cursor() (row, col int)
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

type Renderer struct {
	screen tcell.Screen
	editor editor
	region Region
	offset point
}

func NewRenderer(screen tcell.Screen, editor editor, region Region) Renderer {
	return Renderer{
		screen: screen,
		region: region,
		editor: editor,
	}
}

func (r *Renderer) Render() {
	// TODO calling clear() every render might not be performant
	r.screen.Clear()

	content := r.editor.Content()
	cRow, cCol := r.editor.Cursor()

	r.syncOffset()

	for i := 0; i < len(content) && i < r.region.height; i++ {
		rowToRender := min(i+r.offset.row, len(content)-1)
		colsToRender := min(r.offset.col, len(content[rowToRender]))

		r.renderChars(content[rowToRender][colsToRender:], i, 0)
	}

	r.screen.ShowCursor(r.region.col-r.offset.col+cCol, r.region.row-r.offset.row+cRow)

	r.screen.Show()
}

func (r *Renderer) renderChars(chars []rune, row, col int) {
	for i := 0; i < len(chars) && i < r.region.width; i++ {
		r.renderChar(chars[i], row, col+i)
	}
}

func (r *Renderer) renderChar(char rune, row, col int) {
	x := col + r.region.col
	y := row + r.region.row

	r.screen.SetContent(x, y, char, nil, tcell.StyleDefault)
}

func (r *Renderer) syncOffset() {
	cRow, cCol := r.editor.Cursor()

	if cCol < r.offset.col {
		r.offset.col = cCol
	} else if cCol >= r.offset.col+r.region.width {
		r.offset.col = cCol - r.region.width + 1
	}

	if cRow < r.offset.row {
		r.offset.row = cRow
	} else if cRow >= r.offset.row+r.region.height {
		r.offset.row = cRow - r.region.height + 1
	}
}
