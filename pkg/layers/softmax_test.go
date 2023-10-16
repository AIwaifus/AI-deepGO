package layers

import (
	"fmt"
	"testing"
)

func TestSoftmaxEstimate(t *testing.T) {
	var Max = func