# util

This package provides utility functions for working with files and directories, as well as loading environment variables from a .env file.

The package includes a function `WalkDirIgnored` that walks through a directory and calls a provided function for each file or directory, skipping directories named ".git" and files or directories matching a .gitignore file. It also includes a function `LoadEnv` that loads environment variables from a .env file using the godotenv library and returns the value for a given key.

The package can be configured using the environment variable "BAZ".

The package can be launched in the following ways:

1. By importing the package and calling its functions directly.

2. By using the package as a dependency in another project.

The package structure is as follows:

```
util/
├── internal/
│   └── util.go
```

## WalkDirIgnored

This function takes a directory path and a function as arguments. It compiles a .gitignore file from the directory path and walks through the directory, calling the provided function for each file or directory. It skips directories named ".git" and files or directories matching the .gitignore file.

## LoadEnv

This function takes a key as an argument and loads the .env file using the godotenv library. It returns the environment variable value for the given key. If the .env file loading fails or the key is empty, it logs a fatal error.

