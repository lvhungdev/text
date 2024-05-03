package command

// TODO find out a common func for this interface
type Command interface{}

type InsertChar struct {
	Char rune
}

type InsertNewLine struct{}

type DelChar struct{}

type MovCurDown struct {
	Sel bool
}

type MovCurUp struct {
	Sel bool
}

type MovCurRight struct {
	Sel bool
}

type MovCurLeft struct {
	Sel bool
}

type Save struct{}
