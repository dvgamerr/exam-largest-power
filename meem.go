package main

import "math"

func meem(in int) (int, int) {
	loopbase := int(math.Floor(math.Sqrt(float64(in - 1))))
	sum := 0
	result := 0
	counter := 0

	if loopbase > 1 {
		sum = loopbase

		//loop for base
		for base := loopbase; base > 1; base-- {

			//loop for expo
			for expo := 2; true; expo++ {
				sum = int(math.Pow(float64(base), float64(expo)))

				if sum < in {
					if (in - sum) <= (in - result) {
						if result == sum {
							counter++
						} else {
							counter = 1
						}

						result = sum
					}
				} else {
					break
				}
			}

		}

	} else {
		result = loopbase
		counter = -1
	}

	return result, counter
}
