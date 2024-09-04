# util

This package provides utility functions for working with files and directories, including a function to walk through a directory and skip files that are ignored by a .gitignore file. It also includes a function to load environment variables.

## File Structure

```
internal/
  util.go
```

## Code Summary

### WalkDirIgnored

This function takes a working directory, a path to a .gitignore file, and a function as arguments. It compiles the .gitignore file using the `ignore.CompileIgnoreFile` function. Then, it walks through the directory using `filepath.WalkDir`, skipping the .git directory. For each file or directory found, it calculates the relative path using `filepath.Rel` and checks if the path matches the .gitignore using `ignoreFile.MatchesPath`. If the path is not ignored, it calls the provided function with the path and the corresponding `fs.DirEntry`. Finally, it returns any error encountered during the process.

### LoadEnv

This function takes a key as an argument and returns the value of the environment variable with that key. If the key is empty, it logs a fatal error.

## Edge Cases

The application can be launched with the following environment variables:

- key: The key of the environment variable to be loaded.

The application can also be launched with the following command-line arguments:

- -key: The key of the environment variable to be loaded.

The application can be configured using the following files:

- .gitignore: The file containing the list of files and directories to be ignored.

