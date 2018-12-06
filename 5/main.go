package main

import (
	"io/ioutil"
	"log"
)

func main() {
	bigassstring, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Println("err =", err)
	}

	for iter, _ := range bigassstring {
		if bigassstring[iter] == bigassstring[iter+1]+32 || bigassstring[iter] == bigassstring[iter+1]+32 {
			log.Println("iter, lastLetter, letter =", iter, string(bigassstring[iter+1]), string(bigassstring[iter]))
		}
	}
}
