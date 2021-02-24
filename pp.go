package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(display(10000))
}

func display(n int) (a int, b int) {
	var num, currentNear, currentNearTemp float64
	var maxNo int
	var count int
	var flag bool = false
	for i := 2; i <= int(math.Round(math.Log2(float64(n)))); i++ {
		num = math.Floor(math.Pow(float64(n), float64(1.0/float64(i))))
		if num == 1 {
			i = 20
		} else {
			currentNearTemp = float64(n) - math.Pow(num, float64(i))

			if currentNearTemp != 0 {
				if flag == false {
					flag = true
					currentNear = currentNearTemp
				}
				if currentNearTemp < currentNear {
					currentNear = currentNearTemp
					maxNo = int(n) - int(currentNear)
					count = 1
				} else if currentNearTemp == currentNear {
					currentNear = currentNearTemp
					maxNo = int(n) - int(currentNear)
					count++
				}
			} else {
				n--
				i--
			}

		}
	}
	return maxNo, count
}
