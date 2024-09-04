# util

This package provides utility functions for working with files and directories, as well as loading environment variables from a .env file.

The package includes a function `WalkDirIgnored` that walks through a directory and calls a provided function for each file or directory, skipping files and directories matching a .gitignore file. It also includes a function `LoadEnv` that loads environment variables from a .env file and returns the value for a given key.

The package can be configured using the following environment variables:

- .env file: Contains environment variables to be loaded by the `LoadEnv` function.

Edge cases for launching the application:

- The .env file may not exist or be empty, causing the `LoadEnv` function to return an error.

Project package structure:

- internal/util.go

## WalkDirIgnored

This function takes a directory path and a function as arguments. It compiles a .gitignore file from the directory path and walks through the directory, calling the provided function for each file or directory. It skips files and directories matching the .gitignore file, as well as the .git directory.

## LoadEnv

This function takes a key as an argument and loads the .env file using godotenv. It returns the environment variable value for the given key. If the .env file loading fails or the key is empty, it logs a fatal error.

