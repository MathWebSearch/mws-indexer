package main

import (
	"os"

	"github.com/MathWebSearch/mws-indexer/src"
)

func main() {
	// parse and validate arguments
	args := src.ParseArgs(os.Args)
	if !args.Validate() {
		os.Exit(1)
	}

	// update sources (if needed)
	if !src.UpdateSources(args) {
		os.Exit(1)
	}

	// generate and update the new index
	if !src.GenerateIndex(args) {
		os.Exit(1)
	}
}
