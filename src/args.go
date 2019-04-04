package src

import (
	"flag"
	"fmt"
	"os"
)

// Args represents command-line arguments
type Args struct {
	harvestDir string
	indexDir   string

	mwsIndexExec      string
	harvests2jsonExec string

	temaSearchMode bool
}

// ParseArgs parses arguments from a list of strings
func ParseArgs(args []string) *Args {
	var flags Args

	// create a new flagset
	// that prints it's usage on --help
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(flagSet.Output(), "Usage of %s:\n", args[0])
		flagSet.PrintDefaults()
	}

	// all our arguments
	flagSet.BoolVar(&flags.temaSearchMode, "tema", false, "Generate indexes for tema-search")
	flagSet.StringVar(&flags.harvestDir, "harvest-dir", "/data/", "Path to harvest directory")
	flagSet.StringVar(&flags.indexDir, "index-dir", "/index/", "Path to index directory")
	flagSet.StringVar(&flags.mwsIndexExec, "mws-index", "/mws/bin/mws-index", "mws-index executable")
	flagSet.StringVar(&flags.harvests2jsonExec, "harvests2json", "/mws/bin/harvests2json", "harvests2json executable")

	// parse and exit
	flagSet.Parse(args[1:])
	return &flags
}

// Validate validates the command-line arguments or panics
func (args *Args) Validate() bool {

	// tema: Tema-search mode
	fmt.Printf("tema: %t\n", args.temaSearchMode)

	// harvest directory: ensure that this is a directory
	if !ensureDirectory(args.harvestDir) {
		fmt.Printf("harvest-dir: %q is not a directory\n", args.harvestDir)
		return false
	}
	fmt.Printf("harvest-dir: %q\n", args.harvestDir)

	// index directory: ensure that this is a directory
	if !ensureDirectory(args.indexDir) {
		fmt.Printf("index-dir: %q is not a directory\n", args.indexDir)
		return false
	}
	fmt.Printf("index-dir: %q\n", args.indexDir)

	// mws-index: executable to call
	fmt.Printf("mws-index: %q\n", args.mwsIndexExec)

	// harvests-2-json: json executable
	fmt.Printf("harvests2json: %q\n", args.harvests2jsonExec)

	fmt.Println("------------------------------------------")
	return true
}

// ensureDirectoryOrPanic ensures that caniddate is a directory or otherwise panics with message
func ensureDirectory(candidate string) bool {
	fi, err := os.Stat(candidate)
	if err != nil {
		return false
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return false
	}

	return true
}
