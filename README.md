# Reflexia

Reflexia is a code summarizer that takes a directory path as input and generates a summary of the code in that directory. It uses a language model to generate the summaries and can be configured to ignore specific files or directories. The project also includes a helper function to clone a GitHub repository if a valid link is provided.

## Configuration and Run Instructions

1. Clone the repository: `git clone <repository_url>`
2. Set up environment variables:
    - `PWD`: current working directory
    - `HELPER_URL`: helper URL
    - `MODEL`: model
    - `API_TOKEN`: API token
    - `STOP_WORD`: stop word
3. Run the project: `go run cmd/reflexia/reflexia.go -g <github_link> -u <username> -t <token> -c <directory_path>`

## CLI Arguments

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

## Environment Variables

- `PWD`: current working directory
- `HELPER_URL`: helper URL
- `MODEL`: model
- `API_TOKEN`: API token
- `STOP_WORD`: stop word

