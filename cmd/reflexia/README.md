# reflexia

This package provides a command-line tool for summarizing code in a given directory. It can be used to generate a README.md file for a project, as well as individual file summaries. The tool supports various configuration options, including command-line arguments and environment variables.

## Project Package Structure

- cmd/reflexia/reflexia.go

## Code Summary

### Main Function

The main function initializes a config struct with default values, parses command-line arguments, and calls the processWorkingDirectory function to get the directory path. It then calls the summarizeService.SummarizeCode function to get the file map and the summarizeService.SummarizeRequest function to get the summary content. Finally, it writes the file map and summary content to files.

### processWorkingDirectory Function

This function gets the current working directory and, if a GitHub link is provided, clones the repository to a temporary directory. If there are command-line arguments, it sets the directory path to the first argument. It returns the directory path.

### initConfig Function

This function initializes a config struct with default values and parses command-line arguments. It returns the config struct.

### Command-Line Arguments

- -g, --g: valid link for GitHub repository
- -u, --u: GitHub username for SSH authentication
- -t, --t: GitHub token for SSH authentication
- -c: do not check project root folder specific files such as go.mod or package.json
- -s: do not create SUMMARY.md and README.md, just print the file summaries
- -r: do not create README.md
- -p: do not create README.md for every package in the project
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- -f: Save individual file summary intermediate result to the FILES.md
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

### Environment Variables

- PWD: current working directory
- HELPER_URL: helper URL
- MODEL: model
- API_TOKEN: API token
- STOP_WORD: stop word

### Edge Cases

- Launching the application without providing a GitHub link or directory path will result in an error.
- If the provided GitHub link is invalid or the repository cannot be cloned, the application will fail.
- If the directory path is not valid or does not contain any files, the application will not generate any output.

