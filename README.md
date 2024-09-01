# Reflexia

Reflexia is a code summarizer that takes a directory of files as input and generates a summary of the code in each file. It uses an external summarizer service to perform the actual summarization.

## Configuration

The project uses the following environment variables for configuration:

- PWD: The current working directory.
- HELPER_URL: The URL of the summarizer service.
- MODEL: The model to use for summarization.
- API_TOKEN: The API token for the summarizer service.

## Run Instructions

1. Set the environment variables as described above.
2. Run the project using the following command: `go run cmd/reflexia/reflexia.go`

This will process all files in the current directory and generate a summary of the code in each file. The summary will be stored in a file named "SUMMARY.md". A readme file named "README_GENERATED.md" or "README.md" will also be generated, containing the summary of all summaries.

