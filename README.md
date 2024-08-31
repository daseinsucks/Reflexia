## README.md

This project generates a README file and a SUMMARY file for a given directory based on the content of the files within it. It uses an external summarization service to generate summaries of the code and the overall content.

### Configuration

The project uses environment variables to configure the summarization service and the file filters. Set the following environment variables:

- `PWD`: The current working directory.
- `HELPER_URL`: The URL of the helper service used for content generation.
- `MODEL`: The model used by the summarization service.
- `API_TOKEN`: The API token for the summarization service.

### Run Instructions

1. Make sure the environment variables are set.
2. Run the project using the following command: `go run reflexia.go`

This will generate two files: `README_GENERATED.md` and `SUMMARY.md` in the current directory.