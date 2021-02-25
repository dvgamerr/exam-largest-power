package main

import (
	"math"
)


func pp(n int) (a int, b int) {
	var fixn int = n
	var num, currentNear, currentNearTemp float64
	var maxNo int
	var count int
	var flag bool = false
	for i := 2; i <= int(math.Floor(math.Log2(float64(n)))); i++ {
		num = math.Floor(math.Pow(float64(n), float64(1.0/float64(i))))
		currentNearTemp = float64(fixn) - math.Pow(num, float64(i))
		k := fixn / 2
		j := math.Pow(num, float64(i))

		if j <= float64(k) {
			m := math.Pow(float64(i), float64(i+1))
			if m < float64(fixn) {
				currentNearTemp = float64(fixn) - m
			}
		}

		if num == 1 {
			i = 20
		} else {
			if currentNearTemp != 0 {

				if flag == false {
					flag = true
					currentNear = currentNearTemp
				}
				if currentNearTemp < currentNear {
					currentNear = currentNearTemp
					maxNo = int(fixn) - int(currentNear)
					count = 1
				} else if currentNearTemp == currentNear {
					currentNear = currentNearTemp
					maxNo = int(fixn) - int(currentNear)
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
