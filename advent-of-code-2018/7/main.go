package main

import (
	"github.com/bittelc/advent-of-code-2018/7/text"
	"log"
)

type depMap struct {
	Deps []string
	Node string
}

func main() {
	depSimple := text.DefineData()
	log.Println("depSimple:\n", depSimple)
}
