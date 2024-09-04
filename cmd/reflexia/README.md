# reflexia

This package provides a tool for summarizing code and generating README files for projects. It takes a GitHub repository link, username, and token as input, and then processes the project's files to create summaries and README files. The package also supports various command-line arguments and environment variables for customization.

## Environment variables
- PWD: current working directory
- HELPER_URL: helper URL
- MODEL: model
- API_TOKEN: API token

## Command-line arguments
- -g: valid link for github repository
- -u: github username for ssh auth
- -t: github token for ssh auth
- -l: config filename in project_config to use
- -c: do not check project root folder specific files such as go.mod or package.json
- -s: do not create SUMMARY.md and README.md, just print the file summaries
- -r: do not create README.md
- -p: do not create README.md for every package in the project
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- -f: Save individual file summary intermediate result to the FILES.md
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

## Files and their paths
- cmd/reflexia/reflexia.go

## Edge cases
- Launching the application with no arguments or invalid arguments.
- Launching the application with invalid environment variables.
- Launching the application with an invalid GitHub repository link.

## Project package structure
- cmd/reflexia/reflexia.go

## Code summary
The package starts by initializing the config struct, parsing command-line arguments, and loading environment variables. It then processes the working directory, gets the project config, and creates a summarizer service. The summarizer service is used to summarize the code and generate README files for the project. The results are then written to files, or an error is logged if there is an issue.

The package also includes functions for loading environment variables, getting the project file structure, converting file summaries to strings, and writing files. These functions are used to support the main logic of the package.

The package does not appear to have any dead code or unclear places.