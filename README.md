# main

## Documentation

```
Package: main

Description: This package is responsible for cloning a repository and extracting information about its packages.

Functions:

main():
    1. Loads the .env file to retrieve the REPO_LINK environment variable.
    2. Checks if the REPO_LINK environment variable is set. If not, it logs a fatal error.
    3. Calls the Reflexate function from the Reflexia package to clone the repository and extract package information.

Reflexate(repoLink string):
    1. Clones the repository specified by the repoLink parameter.
    2. Iterates through the repository's files and extracts information about each package.
    3. Creates a Package struct for each package, containing its name, files, and repository.
    4. Returns a list of Package structs.

Package struct:
    1. Name: The name of the package.
    2. Files: A list of files belonging to the package.
    3. Repo: A pointer to the repository object.

Dependencies:

1. github.com/JackBekket/Reflexia/lib/reflexia: This package provides the Reflexate function, which clones the repository and extracts package information.
2. github.com/go-git/go-git/v5: This package is used for interacting with Git repositories.
3. github.com/joho/godotenv: This package is used for loading environment variables from a .env file.

Usage:

1. Set the REPO_LINK environment variable in a .env file.
2. Run the main function to clone the repository and extract package information.

Example:

1. Create a .env file with the following content:
    REPO_LINK=https://github.com/username/repository.git

2. Run the following command:
    go run main.go

3. The Reflexate function will clone the repository and extract package information.

Notes:

1. The code assumes that the .env file is located in the same directory as the main.go file.
2. The code requires the go-git and godotenv packages to be installed.



