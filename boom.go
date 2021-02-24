package main

import (
	"math"
	"reflect"
)

// Reduce si
func Reduce(slice, pairFunction, zero interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("reduce: not slice")
	}
	n := in.Len()
	switch n {
	case 0:
		return zero
	case 1:
		return in.Index(0)
	}
	fn := reflect.ValueOf(pairFunction)

	var ins [2]reflect.Value
	ins[0] = in.Index(0)
	ins[1] = in.Index(1)
	out := fn.Call(ins[:])[0]
	// Run from index 2 to the end.
	for i := 2; i < n; i++ {
		ins[0] = out
		ins[1] = in.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

func boom(item float64) (int, int) {
	var listUp = []float64{}
	for i := 0.0; i < item; i++ {
		c := i + 1
		for j := 0.0; j < item; j++ {
			b := j + 1
			if b > 1 {
				a := math.Pow(c, b)
				listUp = append(listUp, a)
			}
		}
	}

	sum := Reduce(listUp, func(a int, b int) int {
		if math.Abs(float64(b)-item) < math.Abs(float64(b)-item) {
			return b
		}
		return a
	}, 1).(int)

	// sum := listUp.reduce((prev, curr) => );

	count := 0
	for t := 0; t < sum; t++ {
		t1 := t + 1
		for k := 0; k < sum; k++ {
			b1 := k + 1
			if b1 > 1 {
				a1 := math.Pow(float64(t1), float64(b1))
				if int(a1) == sum {
					count++
				}
			}
		}
	}
	return sum, count
}
