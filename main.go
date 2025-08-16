package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command line flags
	versionFlag := flag.Bool("version", false, "Print version information")
	helpFlag := flag.Bool("help", false, "Print help information")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of pi:\n")
		fmt.Fprintf(os.Stderr, "  pi [options] [arguments]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
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

	// If no arguments provided, print welcome message
	if flag.NArg() == 0 {
		fmt.Println("Welcome to pi - a simple CLI program!")
		fmt.Println("Use -help flag to see available options.")
		return
	}

	// Handle positional arguments
	args := flag.Args()
	fmt.Println("Arguments provided:", args)
}
