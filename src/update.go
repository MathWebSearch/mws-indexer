package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

// GenerateIndex generates a new index for use with MathWebSearch
func GenerateIndex(args *Args) bool {
	// create a new temporary directory
	tmpDir, err := ioutil.TempDir("", "mws-index")
	if err != nil {
		return false
	}
	fmt.Printf("Created temporary directory %q\n", tmpDir)
	defer os.RemoveAll(tmpDir)

	result := callWithInheritIO(args.mwsIndexExec, "--recursive", "--include-harvest-path", args.harvestDir, "--output-directory", tmpDir)
	if !result {
		fmt.Printf("mws-index failed, exiting.")
		return false
	}

	fmt.Printf("Update %q with new index from %q.\n", args.indexDir, tmpDir)

	content := updateContents(tmpDir, args.indexDir)
	if content != nil {
		fmt.Println("Update Content failed (is there enough space?)")
		return false
	}

	return true
}

func callWithInheritIO(command string, args ...string) bool {
	fmt.Printf("Calling %q %s\n", command, strings.Join(args, " "))

	// create the command
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run command and check that no error occurs
	return cmd.Run() == nil
}

func updateContents(source string, dest string) (err error) {
	err = removeContents(dest)
	if err != nil {
		return
	}

	return copy.Copy(source, dest)
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
