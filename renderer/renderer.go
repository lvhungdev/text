package renderer

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

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
	lineNumbPad := len(strconv.Itoa(len(content)))

	r.syncOffset(r.region.width-lineNumbPad-2, r.region.height)

	for row := r.offset.row; row < len(content) && row < r.region.height+r.offset.row; row++ {
		r.renderLine(row, lineNumbPad)
	}

	r.screen.ShowCursor(r.region.col-r.offset.col+lineNumbPad+2+cCol, r.region.row-r.offset.row+cRow)

	r.screen.Show()
}

func (r *Renderer) renderLine(contentRow int, lineNumbPad int) {
	line := r.editor.Content()[contentRow]
	row := contentRow - r.offset.row

	lineNumb := strconv.Itoa(contentRow + 1)
	lineNumbStyle := tcell.StyleDefault.Foreground(tcell.ColorGray)
	renderChars(r.screen, r.region, row, lineNumbPad-len(lineNumb), []rune(lineNumb), lineNumbStyle)

	line = line[min(r.offset.col, len(line)):]
	for i := 0; i < len(line) && i < r.region.width; i++ {
		style := newStyleBuilder().withSelection(r.isInSelection(contentRow, i+r.offset.col)).build()

		renderChar(r.screen, r.region, row, i+lineNumbPad+2, line[i], style)
	}
}

func (r *Renderer) syncOffset(regWidth, regHeight int) {
	cRow, cCol := r.editor.Cursor()

	if cCol < r.offset.col {
		r.offset.col = cCol
	} else if cCol >= r.offset.col+regWidth {
		r.offset.col = cCol - regWidth + 1
	}

	if cRow < r.offset.row {
		r.offset.row = cRow
	} else if cRow >= r.offset.row+regHeight {
		r.offset.row = cRow - regHeight + 1
	}
}

func (r *Renderer) isInSelection(row, col int) bool {
	bRow, bCol, eRow, eCol := r.editor.Selection()

	p := point{row: row, col: col}
	b := point{row: bRow, col: bCol}
	e := point{row: eRow, col: eCol}

	if b.Compare(e) == 0 {
		return false
	}

	return (p.Compare(b) >= 0 && p.Compare(e) <= 0) || (p.Compare(e) >= 0 && p.Compare(b) <= 0)
}
