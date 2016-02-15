package main

import (
	m "math"

	"github.com/MaxHalford/gago"
)

// DropWave minimum is -1 reached in (0, 0)
// Recommended search domain is [-5.12, 5.12]
func dropWave(X []float64) float64 {
	numerator := 1 + m.Cos(12*m.Sqrt(m.Pow(X[0], 2)+m.Pow(X[1], 2)))
	denominator := 0.5*(m.Pow(X[0], 2)+m.Pow(X[1], 2)) + 2
	return -numerator / denominator
}

func main() {
	// Instantiate a population
	ga := gago.Float
	// Wrap the function
	ga.Ff = gago.FloatFunction{dropWave}
	// Initialize the genetic algorithm with two variables per individual
	ga.Initialize(2)
	// Enhancement
	for i := 0; i < 30; i++ {
		ga.Best.Display()
		ga.Enhance()
	}
}