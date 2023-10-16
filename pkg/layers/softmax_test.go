package layers

import (
	"fmt"
	"testing"
)

func TestSoftmaxEstimate(t *testing.T) {
	var Max = func(a []float64) int {
		key :