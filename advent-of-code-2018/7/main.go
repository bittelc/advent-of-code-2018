package main

import (
	"fmt"
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

	solutionString := ""
	var candidates []string
FINDER:
	for {
		candidates = []string{}
		for node, deps := range tree {
			if len(deps) == 0 {
				candidates = append(candidates, node)
				log.Println("node, candidates:", node, candidates)
			}
		}
		solutionString = solutionString + fmt.Sprintf("%v", candidates) + "|"
			
		for node, deps := range tree {
				for _, candidate := range candidates {
					if contains(deps, candidate) {
						log.Printf("%v has dependency on %v from its list: %v", , node, depsWithNode)
						i := indexOf(node, depsWithNode)
						log.Println("i =", i)
						depsWithNode := append(depsWithNode[:i], depsWithNode[i+1:]...)
						tree[nodeWithNode] = depsWithNode
						log.Printf("%v doesn't have dep on %v list:        %v", nodeWithNode, node, depsWithNode)
					}
				}
			}
		}

		if len(solutionString) > 60 {
			for v, k := range tree {
				log.Printf("%v:%v\n", v, k)
			}
			break FINDER
		}
	}
	log.Println(solutionString)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
