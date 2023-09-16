package layers

import (
	"math"
	"strconv"
	"testing"
)

func TestSigmoidActivate(t *testing.T) {
	var sigmoid = func(x float64) float64 {
		return math.Exp(x) / (1.0 + m