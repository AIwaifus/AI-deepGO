package graph

import (
	"fmt"
	"strings"
)

type Metrics struct {
	Actual   []float64
	Estimate []float64
	Epoch    int
	Loss     []float64