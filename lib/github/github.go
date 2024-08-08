package github

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)
  
  func Clone(url string, directory string) error {
	// Clone the given repository to the given directory
	fmt.Printf("git clone %s %s --recursive\n", url, directory)
  
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
	  URL:               url,
	  RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
  
	return err
  }