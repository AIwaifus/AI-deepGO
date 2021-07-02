
package graph

import (
	"math/rand"
)

type Features struct {
	ClassWeights        []float64
	DisableClassWeights bool
	DisableShuffle      bool
	X                   [][]float64
	Y                   [][]float64
}