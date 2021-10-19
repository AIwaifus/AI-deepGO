package layers

type Dense struct {
	bias    Bias
	dense   UnbiasedDense
	Ne