package main

import (
	"log"
	"math"
)

type location struct {
	rowNum int
	xCord  int
	yCord  int
	area   int
}

var locs = []location{
	{1, 77, 279, 0},
	{2, 216, 187, 0},
	{3, 72, 301, 0},
	{4, 183, 82, 0},
	{5, 57, 170, 0},
	{6, 46, 335, 0},
	{7, 55, 89, 0},
	{8, 71, 114, 0},
	{9, 313, 358, 0},
	{10, 82, 88, 0},
	{11, 78, 136, 0},
	{12, 339, 314, 0},
	{13, 156, 281, 0},
	{14, 260, 288, 0},
	{15, 125, 249, 0},
	{16, 150, 130, 0},
	{17, 210, 271, 0},
	{18, 190, 258, 0},
	{19, 73, 287, 0},
	{20, 187, 332, 0},
	{21, 283, 353, 0},
	{22, 66, 158, 0},
	{23, 108, 97, 0},
	{24, 237, 278, 0},
	{25, 243, 160, 0},
	{26, 61, 52, 0},
	{27, 353, 107, 0},
	{28, 260, 184, 0},
	{29, 234, 321, 0},
	{30, 181, 270, 0},
	{31, 104, 84, 0},
	{32, 290, 109, 0},
	{33, 193, 342, 0},
	{34, 43, 294, 0},
	{35, 134, 211, 0},
	{36, 50, 129, 0},
	{37, 92, 112, 0},
	{38, 309, 130, 0},
	{39, 291, 170, 0},
	{40, 89, 204, 0},
	{41, 186, 177, 0},
	{42, 286, 302, 0},
	{43, 188, 145, 0},
	{44, 40, 52, 0},
	{45, 254, 292, 0},
	{46, 270, 287, 0},
	{47, 238, 216, 0},
	{48, 299, 184, 0},
	{49, 141, 264, 0},
	{50, 117, 129, 0},
	{51, 0, 0, 0},
}

func printArea() {
	max := 0
	for _, val := range locs {
		if val.area > max {
			max = val.area
		}
	}
	log.Println("max =", max)
}

func main() {
	var distance, minDistance, minRow int

	for x := 40; x <= 353; x++ {
		for y := 52; y <= 358; y++ {
			minDistance = 98908
			for rowNum, l := range locs {
				distance = int(math.Abs(float64(l.xCord-x)) + math.Abs(float64(l.yCord-y)))
				if distance == minDistance {
					minRow = 50
				}
				if distance < minDistance {
					minDistance = distance
					minRow = rowNum
				}
				if x == 40 || y == 52 || x == 353 || y == 358 {
					locs[minRow].area = -100000
				}
			}
			locs[minRow].area++
		}
	}
	printArea()
}
