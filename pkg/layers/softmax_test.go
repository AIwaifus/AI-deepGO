package layers

import (
	"fmt"
	"testing"
)

func TestSoftmaxEstimate(t *testing.T) {
	var Max = func(a []float64) int {
		key := 0
		value := a[key]
		for k, v := range a {
			if v > value {
				key = k
				value = v
			}
		}
		return key
	}
	for k, v := range []struct {
		z    []float64
		want []float64
	}{
		{[]float64{0, 1, 2, 3, 4}, []float64{0, 0, 0, 0, 1}},
		{[]float64{4, 3, 2, 1, 0}, []float64{1, 0, 0, 0, 0}},
		{[]float64{5, 2, -1, 0, 3}, []float64{1, 0, 0, 0, 0