package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/aragorn/montecarlo/montecarlo"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	// Define command line flags
	versionFlag := flag.Bool("version", false, "Print version information")
	helpFlag := flag.Bool("help", false, "Print help information")
	pointsFlag := flag.Int("points", 1000000, "Number of points to use in Monte Carlo simulation")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of pi:\n")
		fmt.Fprintf(os.Stderr, "  pi [options] [arguments]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  pi -points 10000000\n")
	}

	// Parse command line arguments
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Println("pi version 1.0.0")
		return
	}

	// Handle help flag
	if *helpFlag {
		flag.Usage()
		return
	}

	// If arguments provided, try to parse as number of points
	if flag.NArg() > 0 {
		if numPoints, err := strconv.Atoi(flag.Arg(0)); err == nil {
			*pointsFlag = numPoints
		} else {
			fmt.Println("Arguments provided:", flag.Args())
			return
		}
	}

	// Calculate Pi using Monte Carlo method
	pi, _error := montecarlo.CalculatePiWithError(*pointsFlag)

	fmt.Printf("π approximation using Monte Carlo method with %d points: %.10f\n", *pointsFlag, pi)
	fmt.Printf("Error compared to math.Pi: %.10f\n", _error)
	fmt.Printf("Actual value of π: %.10f\n", 3.14159265359)
}
