package main

import "fmt"

func main() {
	fmt.Println("Power Larger!!")
	for i := 1; i <= 1025; i++ {
		l, c := kem(i)
		fmt.Printf("- %d\t%d\n", l, c)
	}
}
