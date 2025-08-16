package montecarlo

import (
	"math"
	"math/rand"
	"time"
)

// CalculatePi calculates an approximation of π using the Monte Carlo method
// numPoints: the number of random points to generate
// Returns the approximation of π and the ratio of points inside the circle
func CalculatePi(numPoints int) (float64, float64) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Counter for points inside the circle
	insideCircle := 0

	// Generate random points and count how many fall inside the circle
	for i := 0; i < numPoints; i++ {
		// Generate random coordinates between -1 and 1
		x := 2*rand.Float64() - 1
		y := 2*rand.Float64() - 1

		// Check if the point is inside the circle using the Pythagorean theorem
		// A point (x,y) is inside the circle of radius 1 if x² + y² ≤ 1
		if x*x+y*y <= 1 {
			insideCircle++
		}
	}

	// Calculate the ratio of points inside the circle to total points
	ratio := float64(insideCircle) / float64(numPoints)

	// π is approximately 4 times this ratio
	piApproximation := 4 * ratio

	return piApproximation, ratio
}

// CalculatePiWithError calculates π using the Monte Carlo method and returns the error compared to math.Pi
func CalculatePiWithError(numPoints int) (float64, float64) {
	piApproximation, _ := CalculatePi(numPoints)
	error := math.Abs(piApproximation - math.Pi)
	return piApproximation, error
}
