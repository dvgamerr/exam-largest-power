package main

import "math"

func copter(n int) (int, int) {
	// powerCloser := []string{};
	max := 0
	no := 0
	for i := 2; i < n/2; i++ {
		for j := 2; j < n/2; j++ {

			powerNo := int(math.Pow(float64(i), float64(j)))

			if powerNo > n {
				break
			}

			if powerNo == max {
				// powerCloser = append(powerCloser,"Number Base: "+ fmt.Sprint(i) +" Power No: "+ fmt.Sprint(powerNo))
				no++
				continue
			}

			if powerNo < n && powerNo > max {
				// powerCloser = []string{}
				// powerCloser = append(powerCloser,"Number Base: "+ fmt.Sprint(i) +" Power No: "+ fmt.Sprint(powerNo))
				no = 1
				max = powerNo
			}
		}
	}
	// fmt.Println("Number Input " + fmt.Sprint(n))
	// for no := 0 ; no < len(powerCloser); no++ {
	// 	fmt.Println(powerCloser[no])
	// }
	return max, no
}
