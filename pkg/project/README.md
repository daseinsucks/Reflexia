# project

This package provides a way to configure and manage project files based on a set of filters and prompts. It reads configuration from files and builds package files based on the provided ModuleMatch. The package also includes functions to check if specific filters are present in the directory.

Environment variables, flags, or cmdline arguments can be used for configuration.

Edge cases for launching the application include:

- No configuration files found
- Invalid configuration file format
- Missing or invalid ModuleMatch

Project package structure:

- pkg/project/project.go

Summary of major code parts:

## ProjectConfig

The ProjectConfig struct holds the configuration for the project, including:

- FileFilter: A list of file extensions to include in the project.
- ProjectRootFilter: A list of file extensions to exclude from the project.
- CodePrompt: A prompt to be used when generating code.
- SummaryPrompt: A prompt to be used when generating a project summary.
- ReadmePrompt: A prompt to be used when generating a README file.
- PackagePrompt: A prompt to be used when generating a package file.
- RootPath: The root directory of the project.
- ModuleMatch: A regular expression to match the module name.

## GetProjectConfig

The GetProjectConfig function reads the configuration from files and returns a ProjectConfig struct.

## BuildPackageFiles

The BuildPackageFiles function takes a ProjectConfig struct and returns a map of package files to their corresponding source files. It builds the package files based on the ModuleMatch provided in the ProjectConfig struct.

## hasFilterFiles

The hasFilterFiles function takes a directory path and a list of file extensions as input and returns true if any of the filters are present in the directory.

## hasRootFilterFile

The hasRootFilterFile function takes a directory path and a list of file extensions as input and returns true if any of the filters are present in the directory.

