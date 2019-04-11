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

	mwsResult := generateMWSIndex(args, tmpDir)
	if !mwsResult {
		fmt.Printf("mws-index failed, exiting.")
		return false
	}

	fmt.Printf("Update %q with new index from %q.\n", args.indexDir, tmpDir)

	content := updateContents(tmpDir, args.indexDir)
	if content != nil {
		fmt.Println("Update Content failed (is there enough space?)")
		return false
	}

	if args.temaSearchMode {
		return generateIndexTema(args)
	}

	return true
}

func generateIndexTema(args *Args) bool {
	// create a new temporary directory
	tmpDir, err := ioutil.TempDir("", "mws-tema-index")
	if err != nil {
		return false
	}
	fmt.Printf("Created temporary directory %q\n", tmpDir)
	defer os.RemoveAll(tmpDir)

	temaResult := generateTemaIndex(args, tmpDir)
	if !temaResult {
		fmt.Printf("harvests2json failed, exiting.")
		return false
	}

	temaMoveResult := moveTemaIndex(args, tmpDir)
	if !temaMoveResult {
		fmt.Printf("failed to move generated tema-search index, exiting.")
		return false
	}

	content := updateContents(tmpDir, args.temaIndexDir)
	if content != nil {
		fmt.Println("Update Content failed (is there enough space?)")
		return false
	}

	return true
}

func generateMWSIndex(args *Args, tmpDir string) bool {
	fmt.Println("Running mws-index")
	return callWithInheritIO(args.mwsIndexExec, "--recursive", "--include-harvest-path", args.harvestDir, "--output-directory", tmpDir)
}

func generateTemaIndex(args *Args, tmpDir string) bool {
	fmt.Println("Running harvests2json")
	return callWithInheritIO(args.harvests2jsonExec, "--recursive", "--harvest-path", args.harvestDir, "--index-path", tmpDir)
}

func moveTemaIndex(args *Args, tmpDir string) bool {
	fmt.Println("Moving tema-search index")

	// iterate over json files
	err := filepath.Walk(args.harvestDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".json" {

			// find the destination path to move it to
			relPath, err := filepath.Rel(args.harvestDir, path)
			if err != nil {
				return err
			}
			destPath := filepath.Join(tmpDir, relPath)

			// create it
			destParent := filepath.Dir(destPath)
			if err := os.MkdirAll(destParent, 0700); err != nil {
				return err
			}

			// and move the file
			return moveFile(path, destPath)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return err == nil
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
