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

var GetTagNameByDetailTag = map[uint64][]string{
	1: {"Heart", "Oval", "Triangle"},
	2: {"Oval", "Square"},
	3: {"Circle", "Oval"},
	4: {"Oblong", "Oval"},
	5: {"Heart", "Oval", "Triangle"},
	6: {"Oval", "Square"},

	7:  {"Extraversion"},
	8:  {"Extraversion", "Opennes"},
	9:  {"Extraversion"},
	10: {"Neurotic"},
	11: {"Neurotic"},
	12: {"Neurotic"},
	13: {"Agreeable"},
	14: {"Agreeable"},
	15: {"Agreeable"},
	16: {"Conscientious"},
	17: {"Conscientious"},
	18: {"Conscientious"},
	19: {"Opennes"},
	20: {"Opennes"},
}
