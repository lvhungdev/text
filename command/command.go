package command

// TODO find out a common func for this interface
type Command interface{}

type InsertChar struct {
	Char rune
}

type InsertNewLine struct{}

type DelChar struct{}

type MovCurDown struct{}

type MovCurUp struct{}

type MovCurRight struct{}

type MovCurLeft struct{}

type Save struct{}
