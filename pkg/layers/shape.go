package layers

type Shape []uint64

func (s Shape) Shape() []uint64 {
	shape := make([]uint64, len(s))
	copy(shape, s)
	return shape
}

func (s Shape) Size() int {
	var size uint6