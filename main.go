package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Largest Power - Ranking!!")

	dat, err := ioutil.ReadFile("./bench.out")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
}
