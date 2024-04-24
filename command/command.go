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

type DeleteChar struct{}

func (c DeleteChar) Command() string {
	return "Delete character"
}
