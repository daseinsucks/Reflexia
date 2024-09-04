# util

This package provides a utility function for walking through a directory tree and ignoring files based on a .gitignore file. It takes a working directory, a path to the .gitignore file, and a function to be called for each directory entry. The function will skip directories that match the .gitignore file and call the provided function for the remaining entries.

## Environment variables

None

## Flags

None

## Cmdline arguments

None

## Files and their paths

- .gitignore file: path to the .gitignore file

## Edge cases

- Launching the application with an empty .gitignore file will result in all files and directories being processed.
- Launching the application with a non-existent .gitignore file will result in an error.

## Project package structure

- internal/util.go

## Summary of major code parts

### WalkDirIgnored function

This function takes a working directory, a path to the .gitignore file, and a function to be called for each directory entry. It first compiles the .gitignore file using the ignore.CompileIgnoreFile function. Then, it iterates through the directory tree using filepath.WalkDir. For each entry, it checks if the entry is a directory and if its name is ".git". If it is, it skips the directory using filepath.SkipDir. It also calculates the relative path from the working directory to the current path using filepath.Rel. If the ignoreFile is not nil and it matches the relative path, it skips the directory if the entry is a directory, otherwise it returns nil. Finally, it calls the provided function f(path, d) with the current path and directory entry.

