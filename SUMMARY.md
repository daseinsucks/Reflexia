# Reflexia

This project is a code summarizer that can be used to generate summaries of code files, packages, and the entire project. It uses a summarizer service to generate the summaries, which can be configured using environment variables, flags, and command-line arguments.

The project consists of the following files:

- pkg/summarize/summarize.go
- cmd/reflexia/reflexia.go
- internal/util.go
- pkg/project/project.go

The main logic of the project is as follows:

## Summarizer Service

The SummarizerService is responsible for generating summaries of code files, packages, and the entire project. It has the following methods:

- CodeSummaryRequest: Takes a prompt and content as input and returns a summary.
- SummarizeRequest: Takes a prompt and content as input and returns a summary.
- SummarizeCode: Takes a project configuration as input and returns a map of file paths and summaries.

## Main Function

The main function in cmd/reflexia/reflexia.go initializes the config struct, parses command-line arguments, and calls the SummarizeCode method to get file summaries. It then writes the file summaries to a file called FILES.md. The main function also calls the SummarizeRequest method to get package summaries and writes them to a file called README.md. It also calls the SummarizeRequest method to get project and readme summaries and writes them to SUMMARY.md and README.md, respectively.

## Process Working Directory

The processWorkingDirectory function gets the current working directory and clones the repository if a GitHub link is provided. If no GitHub link is provided, it uses the first command-line argument as the directory path.

## Init Config

The initConfig function initializes the config struct with default values and parses command-line arguments.

## Cli Arguments

The following command-line arguments can be used to configure the project:

- -g, --github-link: valid link for GitHub repository
- -u, --github-username: GitHub username for SSH auth
- -t, --github-token: GitHub token for SSH auth
- -c: do not check project root folder specific files such as go.mod or package.json
- -s: do not create SUMMARY.md and README.md, just print the file summaries
- -r: do not create README.md
- -p: do not create README.md for every package in the project
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- -f: Save individual file summary intermediate result to the FILES.md
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

## Environment Variables

The following environment variables can be used to configure the project:

- PWD: current working directory
- HELPER_URL: helper URL for summarizer service
- MODEL: model for summarizer service
- API_TOKEN: API token for summarizer service
- STOP_WORD: stop words for summarizer service

## Edge Cases

The project can be launched with the following edge cases:

- No GitHub link provided, so the first command-line argument is used as the directory path.
- No summaries are created, and only file summaries are printed to the console.
- README.md is not created for every package in the project.
- README.md for the root project directory is overwritten instead of creating a new file.
- Individual file summary intermediate result is saved to the FILES.md file.
- README_GENERATED.md is created if README.md exists in the package directory instead of overwriting it.

## Dead Code

There is no dead code in the project.

