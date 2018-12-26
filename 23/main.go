package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Chars struct {
	X         int
	Y         int
	Z         int
	R         int
	CloseBots int
}

var filename = "input.txt"
var bots = make(map[int]*Chars)

var botWithBigRadius = 0

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("err:", err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	for i := 0; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		bots[i] = structify(line)
		if bots[i].R > bots[botWithBigRadius].R {
			botWithBigRadius = i
		}
	}
	updateCloseBots(botWithBigRadius)
	log.Println("Bots within biggest radius:", bots[botWithBigRadius].CloseBots)
}

func structify(line string) *Chars {
	chars := &Chars{CloseBots: 0}
	var err error
	comp := strings.Split(line, ",")
	comp = strings.TrimSpace(comp)
	x := strings.TrimPrefix(comp[0], "pos=<")
	chars.X, err = strconv.Atoi(x)
	if err != nil {
		log.Fatalln("err:", err)
	}
	chars.Y, err = strconv.Atoi(comp[1])
	if err != nil {
		log.Fatalln("err:", err)
	}
	z := strings.Trim(comp[2], ">")
	chars.Z, err = strconv.Atoi(z)
	if err != nil {
		log.Fatalln("err:", err)
	}
	r := strings.TrimPrefix(comp[3], " r=")
	chars.R, err = strconv.Atoi(r)
	if err != nil {
		log.Fatalln("err:", err)
	}
	return chars
}

func updateCloseBots(a int) error {
	for b := 0; b < len(bots); b++ {
		xDistance := math.Abs(float64(bots[a].X - bots[b].X))
		yDistance := math.Abs(float64(bots[a].Y - bots[b].Y))
		zDistance := math.Abs(float64(bots[a].Z - bots[b].Z))
		total := xDistance + yDistance + zDistance
		if int(total) <= bots[a].R {
			bots[a].CloseBots++
		}
	}
	return nil
}
