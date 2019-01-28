package src

import (
	"flag"
	"fmt"
	"os"
)

// ValidateArguments validates the command-line arguments or panics
func ValidateArguments() {
	// harvest directory: ensure that this is a directory
	ensureDirectoryOrPanic(harvestDir, fmt.Sprintf("%q is not a directory", harvestDir))
	fmt.Printf("harvest-dir: %q\n", harvestDir)

	// index directory: ensure that this is a directory
	ensureDirectoryOrPanic(indexDir, fmt.Sprintf("%q is not a directory", indexDir))
	fmt.Printf("index-dir: %q\n", indexDir)

	// index: executable to call
	fmt.Printf("mws-index: %q\n", mwsIndexExec)

	// docker-label: name of docker container to restart
	fmt.Printf("docker-label: %q\n", dockerLabel)
}

// ensureDirectoryOrPanic ensures that caniddate is a directory or otherwise panics with message
func ensureDirectoryOrPanic(candidate string, message string) {
	fi, err := os.Stat(candidate)
	if err != nil {
		panic(message)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		panic(message)
	}
}

var harvestDir string
var indexDir string
var mwsIndexExec string
var dockerLabel string

func init() {
	flag.StringVar(&harvestDir, "harvest-dir", "/data/", "Path to harvest directory")
	flag.StringVar(&indexDir, "index-dir", "/index/", "Path to index directory")
	flag.StringVar(&mwsIndexExec, "mws-index", "/mws/bin/mws-index", "mws-index executable")
	flag.StringVar(&dockerLabel, "docker-label", "", "label of docker container to restart")
	flag.Parse()
}
