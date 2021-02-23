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
		{Input: 90, Largest: 81, Count: 2},
		{Input: 81, Largest: 64, Count: 3},
		{Input: 65, Largest: 64, Count: 3},
		{Input: 6, Largest: 4, Count: 1},
		{Input: 3, Largest: 0, Count: 0},
		{Input: 1, Largest: 0, Count: 0},
		{Input: 4, Largest: 0, Count: 0},
		{Input: 5, Largest: 4, Count: 1},
		{Input: 7, Largest: 4, Count: 1},
		{Input: 9, Largest: 8, Count: 1},
	}
	t.Run("KEM", func(t *testing.T) {
		for _, r := range data {
			l, c := kem(r.Input)
			assert.Equal(t, l, r.Largest, fmt.Sprintf("%d Largest in wrong.", r.Input))
			assert.Equal(t, c, r.Count, fmt.Sprintf("%d Counting in wrong.", r.Input))
		}
	})
}
