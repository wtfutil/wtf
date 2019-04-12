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

func (pos *Position) IsValid() bool {
	if pos.height < 1 || pos.left < 0 || pos.top < 0 || pos.width < 1 {
		return false
	}

	return true
}

func (pos *Position) Height() int {
	return pos.height
}

func (pos *Position) Left() int {
	return pos.left
}

func (pos *Position) Top() int {
	return pos.top
}

func (pos *Position) Width() int {
	return pos.width
}
