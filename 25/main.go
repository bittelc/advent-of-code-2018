package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var filename = "example_3.txt"

type Star struct {
	X, Y, Z, T int
}

type Stars []Star
type Constellations []Stars

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("err:", err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	stars := Stars{}
	for i := 0; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		stars = append(stars, structify(line))
	}
	log.Println("stars\n", stars)
	c := &Constellations{}
	for _, star := range stars {
		c.findConstellation(star)
	}
	log.Println("constellations =", *c)
	log.Println("constellations =", len(*c))

	c.condenseConstellations()
	log.Println("constellations =", *c)
	log.Println("constellations =", len(*c))

}

func structify(line string) Star {
	var s Star
	var err error
	chars := strings.Split(line, ",")
	chars[0] = strings.TrimSpace(chars[0])
	s.X, err = strconv.Atoi(chars[0])
	if err != nil {
		log.Fatalln("err:", err)
	}
	s.Y, err = strconv.Atoi(chars[1])
	if err != nil {
		log.Fatalln("err:", err)
	}
	s.Z, err = strconv.Atoi(chars[2])
	if err != nil {
		log.Fatalln("err:", err)
	}
	s.T, err = strconv.Atoi(chars[3])
	if err != nil {
		log.Fatalln("err:", err)
	}
	return s
}

func (c *Constellations) findConstellation(newStar Star) {
	log.Println("newStar:", newStar)
	found := false
CONSTFOUND:
	for i, stars := range *c {
		for _, star := range stars {
			xDistance := math.Abs(float64(newStar.X - star.X))
			yDistance := math.Abs(float64(newStar.Y - star.Y))
			zDistance := math.Abs(float64(newStar.Z - star.Z))
			tDistance := math.Abs(float64(newStar.T - star.T))
			total := xDistance + yDistance + zDistance + tDistance
			if int(total) <= 3 {
				(*c)[i] = append((*c)[i], newStar)
				found = true
				break CONSTFOUND
			}
		}
	}
	if !found {
		*c = append(*c, Stars{newStar})
	}
}

func (c *Constellations) condenseConstellations() {
	for a, aStars := range *c {
	BSTAR:
		for b, bStars := range *c {
			if a == b {
				continue
			}
			for _, aStar := range aStars {
				for _, bStar := range bStars {
					if aStar.X == bStar.X && aStar.Y == bStar.Y && aStar.Z == bStar.Z && aStar.T == bStar.T {
						*c = append((*c)[:b], (*c)[b+1:]...)
						break BSTAR
					}
				}
			}
		}
	}
}
