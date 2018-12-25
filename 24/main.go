package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	NumUnits     int
	UnitHitPoint int
	Immunities   []string
	Weaknesses   []string
}

type Infection struct {
	Groups []Group
}

type ImmuneSystem struct {
	Groups []Group
}

var is ImmuneSystem
var inf Infection

func main() {
	err := structify("example_text.txt")
	if err != nil {
		log.Fatalln("err:", err)
	}
}

func structify(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("line:", line)
		err := makeArmy(line)
		if err != nil {
			return err
		}
	}
	log.Println("is:", is)
	return nil
}

func makeArmy(sentence string) error {
	unitIndex := strings.Index(sentence, " ")
	if unitIndex <= 0 {
		return nil
	}
	is.Groups[0].NumUnits = unitIndex
	log.Println("is:", is)
	return nil
}
