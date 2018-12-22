package main

import (
	//"fmt"
	"github.com/bittelc/advent-of-code-2018/7/text"
	"log"
	"sort"
)

type DepTree map[string][]string

func main() {
	depsSimple := text.DefineData()
	tree := make(map[string][]string)
	for _, singleDep := range *depsSimple {
		tree[singleDep.Node] = append(tree[singleDep.Node], singleDep.Dep)
		if len(tree[singleDep.Dep]) == 0 {
			tree[singleDep.Dep] = make([]string, 0)
		}
	}

	solutionString := ""
	var candidates []string
	for {
		candidates = []string{}
		for nodeA, _ := range tree {
			if len(tree[nodeA]) == 0 {
				candidates = append(candidates, nodeA)
			}
		}
		sort.Strings(candidates)
		if len(candidates) == 0 {
			break
		}
		delete(tree, candidates[0])
		solutionString = solutionString + candidates[0]

		for nodeB, _ := range tree {
			if contains(tree[nodeB], candidates[0]) {
				i := indexOf(candidates[0], tree[nodeB])
				tree[nodeB] = remove(tree[nodeB], i)
			}
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

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
