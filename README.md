# Reflexia

Reflexia is a command-line tool that summarizes the contents of a given project directory. It takes a GitHub link or a directory path as input and uses an external API to generate summaries of the project's code, project, and README files. The tool creates two output files: README_GENERATED.md and SUMMARY.md, which contain the generated summaries.

## Configuration

- Set the environment variables: PWD, HELPER_URL, MODEL, API_TOKEN, GH_TOKEN
- Provide the GitHub link or directory path as a command-line argument:
    - `-g` (GitHub link)
    - `-u` (GitHub username)
    - `-t` (GitHub token)

## Run Instructions

1. Clone the repository: `git clone <repository_url>`
2. Navigate to the project directory: `cd reflexia`
3. Run the tool: `go run cmd/reflexia/reflexia.go -g <github_link> -u <github_username> -t <github_token>`

Alternatively, if you have the directory path, you can run:

4. Run the tool: `go run cmd/reflexia/reflexia.go <directory_path>`

This will generate the README_GENERATED.md and SUMMARY.md files containing the summaries of the project's code, project, and README files.
