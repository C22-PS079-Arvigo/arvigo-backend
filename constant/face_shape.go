package constant

var GetFaceShapeTags = map[uint64][]uint64{
	1: {3},
	2: {1, 5},
	3: {4},
	4: {1, 2, 3, 4, 5, 6},
	5: {2, 6},
	6: {1, 5},
}

var GetIDByShape = map[string]uint64{
	"circle":   1,
	"heart":    2,
	"oblong":   3,
	"oval":     4,
	"square":   5,
	"triangle": 6,
}
