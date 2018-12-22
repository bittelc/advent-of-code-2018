package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bigAssString, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Println("err =", err)
	}
	for iter := 0; iter < len(bigAssString)-2; iter++ {
		if bigAssString[iter] == bigAssString[iter+1]+32 || bigAssString[iter] == bigAssString[iter+1]-32 {
			bigAssString = append(bigAssString[:iter], bigAssString[iter+2:]...)
			iter = iter - 2
			if iter < -1 {
				iter = -1
			}
		}
	}
	log.Println(len(strings.TrimSpace(string(bigAssString))))
}
