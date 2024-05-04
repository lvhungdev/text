package common

type Point struct {
	Row int
	Col int
}

func (p Point) Compare(p1 Point) int {
	if p.Row > p1.Row {
		return 1
	}
	if p.Row < p1.Row {
		return -1
	}
	if p.Col > p1.Col {
		return 1
	}
	if p.Col < p1.Col {
		return -1
	}
	return 0
}
