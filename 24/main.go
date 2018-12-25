package main

import "fmt"

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

func main() {
	fmt.Println("vim-go")
}
