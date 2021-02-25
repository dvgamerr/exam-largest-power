package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataSet struct {
	Input   int
	Largest int
	Count   int
}

func TestLargestPower(t *testing.T) {
	data := []DataSet{
		{Input: 1, Largest: 0, Count: 0},
		{Input: 3, Largest: 0, Count: 0},
		{Input: 4, Largest: 0, Count: 0},
		{Input: 5, Largest: 4, Count: 1},
		{Input: 6, Largest: 4, Count: 1},
		{Input: 7, Largest: 4, Count: 1},
		{Input: 9, Largest: 8, Count: 1},
		{Input: 90, Largest: 81, Count: 2},
		{Input: 81, Largest: 64, Count: 3},
		{Input: 65, Largest: 64, Count: 3},
		{Input: 1025, Largest: 1024, Count: 3},
	}
	t.Run("KEM", func(t *testing.T) {
		for _, r := range data {
			l, c := kem(r.Input)
			assert.Equal(t, l, r.Largest, fmt.Sprintf("%d Largest in wrong.", r.Input))
			assert.Equal(t, c, r.Count, fmt.Sprintf("%d Counting in wrong.", r.Input))
		}
	})
	t.Run("Copter", func(t *testing.T) {
		for _, r := range data {
			l, c := copter(r.Input)
			assert.Equal(t, l, r.Largest, fmt.Sprintf("%d Largest in wrong.", r.Input))
			assert.Equal(t, c, r.Count, fmt.Sprintf("%d Counting in wrong.", r.Input))
		}
	})
	t.Run("Pop - pphanpobe", func(t *testing.T) {
		for _, r := range data {
			l, c := pp(r.Input)
			assert.Equal(t, l, r.Largest, fmt.Sprintf("%d Largest in wrong.", r.Input))
			assert.Equal(t, c, r.Count, fmt.Sprintf("%d Counting in wrong.", r.Input))
		}
	})
	// t.Run("Boom", func(t *testing.T) {
	// 	for _, r := range data {
	// 		l, c := boom(float64(r.Input))
	// 		assert.Equal(t, l, r.Largest, fmt.Sprintf("%d Largest in wrong.", r.Input))
	// 		assert.Equal(t, c, r.Count, fmt.Sprintf("%d Counting in wrong.", r.Input))
	// 	}
	// })
}

func BenchmarkFunction(t *testing.B) {
	low := DataSet{Input: 65, Largest: 64, Count: 3}
	high := DataSet{Input: 1025, Largest: 1024, Count: 3}

	t.Run("KEM", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			kem(low.Input)
			kem(high.Input)
		}
	})

	t.Run("Copter", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			copter(low.Input)
			copter(high.Input)
		}
	})

	t.Run("Pop - pphanpobe", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			pp(low.Input)
			pp(high.Input)
		}
	})

	// t.Run("Boom", func(t *testing.B) {
	// 	for i := 0; i < t.N; i++ {
	// 		boom(float64(low.Input))
	// 		boom(float64(high.Input))
	// 	}
	// })
}
