package main

import (
	"os"

	"github.com/MathWebSearch/mws-updater/src"
)

func main() {
	// validate the arguments
	src.ValidateArguments()

	// generate and update the new index
	if !src.GenerateIndex() {
		os.Exit(1)
	}

	// run post-update hooks
	src.PostUpdateHooks()
}
