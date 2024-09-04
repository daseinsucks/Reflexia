# summarize

This package provides a service for summarizing code within a project. It utilizes a language model to generate summaries for code files based on a given prompt and content. The package also includes functions for handling requests and summarizing code based on a project configuration.

## Environment variables

- HELPER_URL: URL of the helper service
- API_TOKEN: API token for the helper service
- NETWORK: Network for the language model (e.g., "gpt-3.5-turbo")

## Flags

- -helper-url: URL of the helper service
- -api-token: API token for the helper service
- -network: Network for the language model (e.g., "gpt-3.5-turbo")

## Cmdline arguments

- -project-config: Path to the project configuration file

## Files and their paths

- pkg/summarize/summarize.go: Main package file containing the SummarizerService struct and related functions.

## Edge cases

- Launching the application without providing the required environment variables or flags.
- Providing an invalid project configuration file.

## Project package structure

```
pkg/
  summarize/
    summarize.go
```

## Code summary

### SummarizerService

The SummarizerService struct holds the necessary information for interacting with the helper service and language model. It includes fields for the helper URL, model, API token, network, and LLM options.

### CodeSummaryRequest and SummarizeRequest

These functions take a prompt and content as input and call the helper service to generate a response. The response is returned along with any error encountered during the process.

### SummarizeCode

This function takes a project configuration as input and generates summaries for code files within the project. It creates a map to store file paths and summaries, walks through the project directory, and ignores files specified in .gitignore. For each file, it checks if it matches any of the file filters specified in the project configuration. If a match is found, it reads the file content, calculates the relative path, and calls CodeSummaryRequest to generate a summary. The summary is then stored in the file map along with the relative path. Finally, the function returns the file map and any error encountered during the process.

