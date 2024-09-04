# Reflexia

Reflexia is a code summarizer that takes a directory path as input and generates summaries for each file in the directory. It uses a language model to generate the summaries and can be configured to include or exclude specific files based on their content or location.

## Configuration

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

## Run Instructions

1. Set the environment variables PWD and GH_TOKEN.
2. Run the command `go run cmd/reflexia/reflexia.go` with the desired command-line arguments.

For example, to summarize a project located at the GitHub repository https://github.com/example/project, you can run the following command:

```
go run cmd/reflexia/reflexia.go -g https://github.com/example/project
```

This will generate the SUMMARY.md, README.md, and FILES.md files in the current directory.

