# Project: Code Summarizer

This project generates code summaries for a given project directory using a summarizer service. It utilizes environment variables for configuration, including HELPER_URL, MODEL, and API_TOKEN.

## Configuration

1. Set the HELPER_URL environment variable to the URL of the summarizer service.
2. Set the MODEL environment variable to the desired summarization model.
3. Set the API_TOKEN environment variable to the API token for the summarizer service.

## Run Instructions

1. Make sure the required environment variables are set.
2. Run the project using the command: `go run cmd/reflexia/reflexia.go`

This will process the code in the current directory, generate summaries for each file, and create a "SUMMARY.md" file containing a summary of all summaries. Additionally, it will generate a "README_GENERATED.md" or "README.md" file containing a readme based on the generated summaries.

