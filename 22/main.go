package main

import "log"

type Point struct {
	X, Y int
}

type Qualities struct {
	GI, EL, Risk int
}

var graph = make(map[Point]Qualities)

const YTarg = 758
const XTarg = 9
const Depth = 8103

func main() {
	for x := 0; x <= XTarg; x++ {
		for y := 0; y <= YTarg; y++ {
			p := Point{x, y}
			gi := GetGI(p)
			el := GetEL(p)
			risk := p.GetRisk(el)
			graph[p] = Qualities{GI: gi, EL: el, Risk: risk}
		}
	}
	totalRisk := areaRisk()
	log.Println("totalRisk =", totalRisk)
}

func GetGI(p Point) int {
	if graph[p].GI != 0 {
		return graph[p].GI
	}
	if p.X == 0 && p.Y == 0 {
		return 0
	} else if p.X == XTarg && p.Y == YTarg {
		return 0
	} else if p.Y == 0 {
		return p.X * 16807
	} else if p.X == 0 {
		return p.Y * 48271
	}

	xless := GetEL(Point{(p.X - 1), p.Y})
	yless := GetEL(Point{p.X, (p.Y - 1)})
	return xless * yless
}

func GetEL(p Point) int {
	return (GetGI(p) + Depth) % 20183
}

func (p Point) GetRisk(el int) int {
	return el % 3
}

func areaRisk() int {
	var risk int
	for x := 0; x <= XTarg; x++ {
		for y := 0; y <= YTarg; y++ {
			p := Point{x, y}
			risk = risk + graph[p].Risk
		}
	}
	return risk
}
