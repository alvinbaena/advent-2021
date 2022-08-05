package day5

type vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func newVent(x1 int, y1 int, x2 int, y2 int) *vent {
	return &vent{x1: x1, y1: y1, x2: x2, y2: y2}
}
