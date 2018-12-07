package main

import (
	"github.com/bittelc/advent-of-code-2018/text"
	"log"
)

type depMap struct {
	Deps []string
	Node string
}

func main() {
	depSimple := text.defineData()
	log.Println("depSimple:\n", depSimple)
}
