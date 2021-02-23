package main

import "fmt"

func kem(n int) (int, int) {

	for l := n; l >= 2; l-- {
		for i := n; i >= 2; i-- {
			fmt.Printf("l=%d  i=%d", l, i)
		}
	}

	return 0, 0
}
