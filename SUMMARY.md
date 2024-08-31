This project generates a README file and a SUMMARY file for a given directory based on the content of the files within it. It uses an external summarization service to generate summaries of the code and the overall content. The project also uses environment variables to configure the summarization service and the file filters.

## Project Configuration

The project uses environment variables to configure the summarization service and the file filters. The environment variables are:

- PWD: The current working directory.
- HELPER_URL: The URL of the helper service used for content generation.
- MODEL: The model used by the summarization service.
- API_TOKEN: The API token for the summarization service.

## Project Code Summary

### Project Configuration

The `project.go` file contains the logic for loading the project configuration. The `GetProjectConfig` function iterates through predefined project configurations and checks if the current directory contains files matching the specified file filters and root files matching the root file filters. If a match is found, the corresponding configuration is returned; otherwise, an error is logged.

### Summarizer Service

The `summarizer.go` file defines the `SummarizerService` struct, which contains the necessary fields for interacting with the summarization service. The `CodeSummaryRequest` and `SummarizeRequest` methods of the `SummarizerService` struct are responsible for sending requests to the summarization service and returning the responses.

### Main Logic

The `reflexia.go` file contains the main logic of the project. It first loads the environment variables using `godotenv.Load()`. Then, it sets the `dirPath` to the value of the "PWD" environment variable. Next, it creates a `SummarizerService` instance using the values from the environment variables.

The project then calls `project.GetProjectConfig(dirPath)` to get the project configuration for the current directory. It creates a `fileMap` to store the summaries of each file in the directory. The project then walks through the directory using `filepath.WalkDir`, and for each file, it checks if it matches any of the file filters in the project configuration. If a match is found, it reads the file content, calls `summarizerService.CodeSummaryRequest`, and stores the summary in the `fileMap`.

After processing all files, the project creates a content string by concatenating the file names and summaries from the `fileMap`. It then calls `summarizerService.SummarizeRequest` twice: first to get a summary of the content using the `projectConfig.SummaryPrompt`, and then again to get a summary of the previous summary using `projectConfig.ReadmePrompt`.

Finally, the project creates two files: `README_GENERATED.md` and `SUMMARY.md`. It writes the readme summary to `README_GENERATED.md` and the overall summary to `SUMMARY.md`.