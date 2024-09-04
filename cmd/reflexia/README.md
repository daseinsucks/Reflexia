# reflexia

This package provides a tool for summarizing code and generating documentation for projects. It uses a summarizer service to generate summaries of code files, packages, and the entire project. The tool can be configured using command-line arguments and environment variables.

## File Structure

- cmd/reflexia/reflexia.go

## Code Summary

### Main Function

The main function initializes a config struct with default values, parses command-line arguments, and calls various functions to process the working directory, summarize code, and write summaries to files.

### processWorkingDirectory Function

This function determines the directory path to process. If a GitHub link is provided, it clones the repository to a temporary directory. Otherwise, it uses the first command-line argument as the directory path.

### initConfig Function

This function initializes a config struct with default values and parses command-line arguments. It returns the config struct.

### Summarization Process

The tool uses a summarizer service to generate summaries of code files, packages, and the entire project. The summaries are written to files named FILES.md, README.md, SUMMARY.md, and README_GENERATED.md.

### Command-Line Arguments

- -g, --github-link: valid link for GitHub repository
- -u, --github-username: GitHub username for SSH authentication
- -t, --github-token: GitHub token for SSH authentication
- -c: do not check project root folder specific files such as go.mod or package.json
- -s: do not create SUMMARY.md and README.md, just print the file summaries
- -r: do not create README.md
- -p: do not create README.md for every package in the project
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- -f: Save individual file summary intermediate result to the FILES.md
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

### Environment Variables

- PWD: current working directory
- HELPER_URL: helper URL for summarizer service
- MODEL: model for summarizer service
- API_TOKEN: API token for summarizer service
- STOP_WORD: stop words for summarizer service

