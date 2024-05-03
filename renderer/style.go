package renderer

import "github.com/gdamore/tcell/v2"

type styleBuilder struct {
	inSelection bool
}

func newStyleBuilder() *styleBuilder {
	return &styleBuilder{}
}

func (b *styleBuilder) withSelection(sel bool) *styleBuilder {
	b.inSelection = sel
	return b
}

func (b *styleBuilder) build() tcell.Style {
	style := tcell.StyleDefault

	if b.inSelection {
		style = style.Background(tcell.ColorBlue)
	}

	return style
}
