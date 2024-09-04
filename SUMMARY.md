# Reflexia

This project is a code summarizer that takes a directory path as input and generates summaries for each file in the directory. It uses a language model to generate the summaries and can be configured to include or exclude specific files based on their content or location.

The project consists of the following files:

- cmd/reflexia/reflexia.go
- internal/util.go
- pkg/project/project.go
- pkg/summarize/summarize.go

The main function in reflexia.go initializes the configuration, processes the working directory, and calls the summarizeService to generate summaries for each file. The summarizeService uses a language model to generate the summaries and can be configured to include or exclude specific files based on their content or location.

## Code Summary

1. Configuration:
    - The project uses environment variables and command-line arguments to configure the summarization process.
    - Environment variables: PWD, GH_TOKEN
    - Command-line arguments:
        - -g: valid link for github repository
        - -u: github username for ssh auth
        - -t: github token for ssh auth
        - -c: do not check project root folder specific files such as go.mod or package.json
        - -s: do not create SUMMARY.md and README.md, just print the file summaries
        - -r: do not create README.md
        - -p: do not create README.md for every package in the project
        - -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
        - -f: Save individual file summary intermediate result to the FILES.md
        - -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

2. Working Directory:
    - The `processWorkingDirectory()` function determines the directory path based on the provided GitHub link or the first command-line argument.

3. Summarization Service:
    - The `SummarizerService` is responsible for generating summaries for each file in the project.
    - It uses a language model to generate the summaries and can be configured to include or exclude specific files based on their content or location.
    - The `SummarizeCode()` function takes a project configuration as input and walks through the project directory, generating summaries for each file that matches the file filter.

4. File Summaries:
    - The `SummarizeRequest()` function takes a prompt and content as input and calls the helper function `helper.GenerateContentInstruction` to generate the summary.
    - The `CodeSummaryRequest()` function is similar to `SummarizeRequest()` but is specifically designed for code summaries.

5. Output:
    - The project generates three output files:
        - SUMMARY.md: Contains the overall summary of the project.
        - README.md: Contains the README file for the project.
        - FILES.md: Contains individual file summaries.

6. Dead Code:
    - There is no dead code in the project.

7. Unclear Places:
    - There are no unclear places in the project.

