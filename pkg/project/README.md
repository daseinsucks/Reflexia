# project

This package provides a way to read and process project configuration files, extract package names from Go files, and collect file paths based on the configuration. It also includes functions to check if specific files exist in a directory.

Environment variables: None
Flags: None
Cmdline arguments: None

Edge cases:
- Launching the application with an invalid configuration file.
- Launching the application with an empty directory.

Project package structure:
- pkg/project/project.go

## ProjectConfig

The `ProjectConfig` struct holds the configuration for the project. It has fields for the directory, package name, and file filters. The `GetProjectConfig` function reads the configuration from the configuration file(s) and returns a `ProjectConfig` struct with default values for each field.

## BuildPackageFiles

The `BuildPackageFiles` function takes a `ProjectConfig` struct as input and returns a map of package names to a list of file paths. It uses the `ModuleMatch` field to determine whether to use directory or package_name mode. In directory mode, it walks through the directory and collects all files matching the `FileFilter`. In package_name mode, it parses the Go files and extracts the package name from the AST.

## hasFilterFiles

The `hasFilterFiles` function takes a directory path and a list of file filters as input. It checks if any files matching the given filters exist in the directory. It returns true if any matching files are found, false otherwise.

## hasRootFilterFile

The `hasRootFilterFile` function is similar to `hasFilterFiles`, but it checks if any files matching the given filters exist in the root directory of the project. It returns true if any matching files are found, false otherwise.

