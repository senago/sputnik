package domain

type Position struct {
	X float32
	Y float32
}

func NewPosition(x, y float32) Position {
	return Position{
		X: x,
		Y: y,
	}
}
