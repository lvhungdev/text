package command

type Command interface {
	Command() string
}

type InsertChar struct {
	Char rune
}

func (c InsertChar) Command() string {
	return "Insert character"
}

type InsertNewLine struct{}

func (c InsertNewLine) Command() string {
	return "Insert new line"
}

type DelChar struct{}

func (c DelChar) Command() string {
	return "Delete character"
}

type MovCurDown struct{}

func (c MovCurDown) Command() string {
	return "Move cursor down"
}

type MovCurUp struct{}

func (c MovCurUp) Command() string {
	return "Move cursor left"
}

type MovCurRight struct{}

func (c MovCurRight) Command() string {
	return "Move cursor right"
}

type MovCurLeft struct{}

func (c MovCurLeft) Command() string {
	return "Move cursor left"
}
