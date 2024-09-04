# project

This package provides a way to configure and manage project files based on a set of criteria. It reads configuration files in the project_config directory and uses them to build a map of package names to file paths. The package also includes functions to check if files in a given directory match specific filters.

## ProjectConfig

The ProjectConfig struct holds the configuration settings for the package. It has the following fields:

- FileFilter: A string that specifies the pattern to match files.
- ProjectRootFilter: A string that specifies the pattern to match project root directories.
- ModuleMatch: A string that specifies whether to match files based on directory or package name.
- StopWords: A list of strings that should be ignored when matching files.
- CodePrompt, SummaryPrompt, PackagePrompt, ReadmePrompt: Strings that specify prompts for code, summary, package, and readme files, respectively.
- RootPath: A string that specifies the root path of the project.

## GetProjectConfig

The GetProjectConfig function takes the current directory and a boolean flag as input. It walks through all .toml files in the current directory and its subdirectories and returns the first ProjectConfig that matches the given criteria. If no matching ProjectConfig is found, it returns an error.

## BuildPackageFiles

The BuildPackageFiles function takes a ProjectConfig as input and builds a map of package names to a list of file paths. The map is built based on the ModuleMatch field of the ProjectConfig. If ModuleMatch is "directory", it walks through the directory and collects all files that match the FileFilter. If ModuleMatch is "package_name", it parses the Go code and extracts the package name from the AST.

## hasFilterFiles and hasRootFilterFile

The hasFilterFiles and hasRootFilterFile functions check if any files in the given directory match the given filters. They return true if a matching file is found, false otherwise.

## Edge Cases

- If no matching ProjectConfig is found, the package will return an error.
- If the ModuleMatch field is not "directory" or "package_name", the package will return an error.
- If the FileFilter or ProjectRootFilter fields are not provided, the package will return an error.

## Project Package Structure

- project_config/*.toml

