package layers

import (
	"fmt"
	"math"
	"strings"
)

type Polynomial struct {
	layer
	learner
	Degree int
	input  []float64
	terms  [][]float64
}

func (l *Polynomial) Estimate(input []float64) []float64 {
	copy(l