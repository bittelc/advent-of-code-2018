package main

import (
	"github.com/bittelc/advent-of-code-2018/7/text"
	"log"
)

type DepTree map[string][]string

func main() {
	depsSimple := text.DefineData()
	tree := make(map[string][]string)
	//for _, singleDep := range (*depsSimple)[:9] {
	for _, singleDep := range *depsSimple {
		tree[singleDep.Node] = append(tree[singleDep.Node], singleDep.Dep)
		if len(tree[singleDep.Dep]) == 0 {
			tree[singleDep.Dep] = make([]string, 0)
		}
	}

	for v, k := range tree {
		log.Printf("%v:%v\n", v, k)
	}

	for n := 1; n < 27; n++ {
		log.Println("iteration = ", n)
		for node, deps := range tree {
			if len(deps) == 0 {
				log.Println("no deps:", node)
			}
		}
	}
}
