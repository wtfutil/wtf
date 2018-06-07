package wtf

type Position struct {
	top    int
	left   int
	width  int
	height int
}

func NewPosition(top, left, width, height int) Position {
	pos := Position{
		top:    top,
		left:   left,
		width:  width,
		height: height,
	}

	return pos
}

func (pos *Position) Top() int {
	return pos.top
}

func (pos *Position) Left() int {
	return pos.left
}

func (pos *Position) Width() int {
	return pos.width
}

func (pos *Position) Height() int {
	return pos.height
}
