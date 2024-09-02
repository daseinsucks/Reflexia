# Reflexia

Reflexia is a command-line tool that analyzes and summarizes code files within a given directory. It uses an external summarization service to generate summaries for each file and then combines these summaries into a project-level summary. The tool also generates a README file based on the project summary and file summaries.

## Configuration

1. Set environment variables:
    - `GH_LINK`: Valid link for the GitHub repository
    - `GH_USERNAME`: GitHub username for SSH authentication
    - `GH_TOKEN`: GitHub token for SSH authentication

2. Create a `.env` file in the project directory and add the environment variables.

## Run Instructions

1. Clone the repository to a local directory.
2. Run the following command:

```bash
go run cmd/reflexia/reflexia.go -g <GH_LINK> -u <GH_USERNAME> -t <GH_TOKEN>
```

3. The tool will analyze the code files in the specified directory and generate a project summary, file summaries, and a README file.

## Optional Flags

- `-c`: Do not check project root folder specific files such as go.mod or package.json
- `-s`: Do not create SUMMARY.md and README.md, just print the file summaries
- `-r`: Do not create README.md

## Output

The tool will generate the following files:

- `SUMMARY.md`: Project summary
- `README_GENERATED.md`: README file based on the project summary and file summaries

