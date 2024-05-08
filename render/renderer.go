package render

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	c "github.com/lvhungdev/text/common"
)

type Renderer struct {
	screen tcell.Screen
	editor editor
	region region
	offset c.Point
}

func NewRenderer(screen tcell.Screen, editor editor) Renderer {
	return Renderer{
		screen: screen,
		editor: editor,
	}
}

func (r *Renderer) SetRegion(row, col, width, height int) {
	r.region = region{row, col, width, height}
}

func (r *Renderer) Render() {
	// TODO calling clear() every render might not be performant
	r.screen.Clear()

	content := r.editor.Content()
	cRow, cCol := r.editor.Cursor()
	lineNumbPad := len(strconv.Itoa(len(content)))

	r.syncOffset(r.region.width-lineNumbPad-2, r.region.height)

	for row := r.offset.Row; row < len(content) && row < r.region.height+r.offset.Row; row++ {
		r.renderLine(row, lineNumbPad)
	}

	r.screen.ShowCursor(r.region.col-r.offset.Col+lineNumbPad+2+cCol, r.region.row-r.offset.Row+cRow)

	r.screen.Show()
}

func (r *Renderer) renderLine(contentRow int, lineNumbPad int) {
	line := r.editor.Content()[contentRow]
	row := contentRow - r.offset.Row

	lineNumb := strconv.Itoa(contentRow + 1)
	lineNumbStyle := tcell.StyleDefault.Foreground(tcell.ColorGray)
	renderChars(r.screen, r.region, row, lineNumbPad-len(lineNumb), []rune(lineNumb), lineNumbStyle)

	line = line[min(r.offset.Col, len(line)):]
	for i := 0; i < len(line) && i < r.region.width; i++ {
		style := newStyleBuilder().withSelection(r.isInSelection(contentRow, i+r.offset.Col)).build()

		renderChar(r.screen, r.region, row, i+lineNumbPad+2, line[i], style)
	}
}

func (r *Renderer) syncOffset(regWidth, regHeight int) {
	cRow, cCol := r.editor.Cursor()

	if cCol < r.offset.Col {
		r.offset.Col = cCol
	} else if cCol >= r.offset.Col+regWidth {
		r.offset.Col = cCol - regWidth + 1
	}

	if cRow < r.offset.Row {
		r.offset.Row = cRow
	} else if cRow >= r.offset.Row+regHeight {
		r.offset.Row = cRow - regHeight + 1
	}
}

func (r *Renderer) isInSelection(row, col int) bool {
	b, e := r.editor.Selection()
	p := c.Point{Row: row, Col: col}

	if b.Compare(e) == 0 {
		return false
	}

	return (p.Compare(b) >= 0 && p.Compare(e) <= 0) || (p.Compare(e) >= 0 && p.Compare(b) <= 0)
}
