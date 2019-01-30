package src

import (
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// UpdateSources updates the sources used by the index updater
func UpdateSources() (success bool) {
	// try and open the repository
	r, err := git.PlainOpen(harvestDir)
	if err != nil {
		return true
	}

	fmt.Printf("Trying to update git repository in %q\n", harvestDir)

	// get the working tree
	w, err := r.Worktree()
	if err != nil {
		fmt.Print("Failed to get worktree, aborting. \n")
		return false
	}

	// run git pull, and send the error to the user
	e := w.Pull(&git.PullOptions{Progress: os.Stdout})
	if e != nil {
		fmt.Println(e.Error())
	}

	// but continue anyways
	return true
}
