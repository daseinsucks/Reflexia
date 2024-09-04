# Reflexia

This project is a code summarizer that takes a directory path as input and generates a summary of the code in that directory. It uses a language model to generate the summaries and can be configured to ignore specific files or directories. The project also includes a helper function to clone a GitHub repository if a valid link is provided.

Project structure:
- cmd/reflexia/reflexia.go
- internal/util.go
- pkg/project/project.go
- pkg/summarize/summarize.go

Summary of the code:

## Main logic

The main function in `cmd/reflexia/reflexia.go` initializes a config struct with default values, parses command-line arguments, and calls the `processWorkingDirectory` function to get the directory path. It then calls the `summarizeService.SummarizeCode` function to get the file map and writes it to a file. Finally, it calls the `summarizeService.SummarizeRequest` function to get the summary content and writes it to a file.

## Process working directory

The `processWorkingDirectory` function takes the GitHub link, username, and token as input. It gets the current working directory and, if a GitHub link is provided, clones the repository to a temporary directory. If there are command-line arguments, it sets the directory path to the first argument. Finally, it returns the directory path.

## Initialize config

The `initConfig` function initializes a config struct with default values and parses command-line arguments. It returns the config struct.

## CLI arguments

- `-g, --g`: valid link for GitHub repository
- `-u, --u`: GitHub username for SSH auth
- `-t, --t`: GitHub token for SSH auth
- `-c`: do not check project root folder specific files such as go.mod or package.json
- `-s`: do not create SUMMARY.md and README.md, just print the file summaries
- `-r`: do not create README.md
- `-p`: do not create README.md for every package in the project
- `-br`: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- `-f`: Save individual file summary intermediate result to the FILES.md
- `-bp`: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

## Environment variables

- `PWD`: current working directory
- `HELPER_URL`: helper URL
- `MODEL`: model
- `API_TOKEN`: API token
- `STOP_WORD`: stop word

## Internal util

The `util.WalkDirIgnored` function compiles a .gitignore file from the given directory path and walks through the directory, calling the provided function for each file/directory. It skips directories named ".git" and files/directories matching the .gitignore file.

The `LoadEnv` function loads the .env file using godotenv and returns the environment variable value for the given key. It logs a fatal error if the .env file loading fails or the key is empty.

## Project

The `GetProjectConfig` function reads the configuration from files and returns a `ProjectConfig` struct.

The `BuildPackageFiles` function returns a map of package files based on the `ModuleMatch` field in the `ProjectConfig` struct.

The `hasFilterFiles` and `hasRootFilterFile` functions check if any of the provided filters are present in the given directory.

## Summarize

The `SummarizerService` struct has fields for the helper URL, model, API token, network, and LLM options.

The `CodeSummaryRequest` and `SummarizeRequest` functions take a prompt and content as input, call the helper function to generate the content instruction, and return the response and any error encountered.

The `SummarizeCode` function takes a `ProjectConfig` struct as input, creates a map to store file paths and summaries, and walks through the project directory using `util.WalkDirIgnored`. For each file, it checks if it matches the file filter in the `ProjectConfig` struct. If it matches, it reads the file content, calculates the relative path, and calls `s.CodeSummaryRequest` to get the summary. Finally, it stores the summary in the map with the relative path as the key and returns the map and any error encountered.

