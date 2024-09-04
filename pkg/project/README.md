# project

This package provides a way to configure and manage project files based on a set of filters and prompts. It reads configuration from .toml files and builds a map of package names to file paths. The package also includes functions to check if files match specific filters and to determine if a directory contains a root filter file.

## ProjectConfig

The `ProjectConfig` struct holds the configuration for the project, including file filters, project root filters, module matches, stop words, code prompts, summary prompts, package prompts, readme prompts, and the root path.

## GetProjectConfig

The `GetProjectConfig` function reads all .toml files in the current directory and its subdirectories and returns the first matching `ProjectConfig` found. If no matching `ProjectConfig` is found, it returns an error.

## BuildPackageFiles

The `BuildPackageFiles` function builds a map of package names to a list of file paths based on the provided `ProjectConfig`.

## hasFilterFiles

The `hasFilterFiles` function checks if any of the files in the given directory match any of the filters provided. It returns true if a match is found, false otherwise.

## hasRootFilterFile

The `hasRootFilterFile` function checks if any of the files in the given directory match any of the filters provided. It returns true if a match is found, false otherwise.

## File Structure

```
pkg/
  project/
    project.go
```

