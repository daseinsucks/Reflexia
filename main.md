# main

## Documentation

# Code Documentation

## Package main

This package is the main entry point of the application. It imports several packages and defines a custom `Package` struct and the `main` function.

## Imports

- `log`: This package provides functions for logging errors and other information.
- `os`: This package provides a way to use operating system dependent functionality.
- `reflexia`: This package provides a function for analyzing a repository.
- `github.com/go-git/go-git/v5`: This package provides a Git repository object that can be used to handle Git operations.
- `github.com/joho/godotenv`: This package provides a way to load environment variables from a `.env` file.

## Structs

### Package

This struct is used to represent a package. It has three fields:

- `Name`: A string representing the name of the package.
- `Files`: A slice of strings representing the files in the package.
- `Repo`: A pointer to a `git.Repository` object representing the Git repository.

## Functions

### main

This function is the main entry point of the application. It loads environment variables from a `.env` file and uses the `reflexia.Reflexate` function to analyze a Git repository.

If there is an error loading the `.env` file or if the `REPO_LINK` environment variable is not set, the function logs an error and exits.

```go
package main

import (
	"log"
	"os"

	reflexia "github.com/JackBekket/Reflexia/lib/reflexia"
	git "github.com/go-git/go-git/v5"
	"github.com/joho/godotenv"
)

type Package struct {
	Name  string
	Files []string
	Repo  *git.Repository
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repoLink := os.Getenv("REPO_LINK")
	if repoLink == "" {
		log.Fatal("REPO_LINK is not set in .env")
	}

	reflexia.Reflexate(repoLink)
}
```

This code is a documentation for the provided Go code. It explains the purpose of the package, the imports, the structs, and the functions.


