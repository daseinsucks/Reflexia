This project code is designed to generate code summaries for a given project directory. It uses a summarizer service to process the code and create a summary of the project. The project code is structured into three main parts: cmd/reflexia/reflexia.go, internal/project/project.go, and internal/summarizer/summarize.go. The summarizer service is responsible for generating summaries of code and project-level summaries. The project configuration is loaded from environment variables and a project configuration file.

## cmd/reflexia/reflexia.go
This part of the code is responsible for loading the project configuration, iterating through the files in the project directory, and generating summaries for each file. It also generates a final summary of all the summaries and writes it to a file named "SUMMARY.md".

1. Loads environment variables using godotenv.Load()
2. Sets up SummarizerService with values from environment variables
3. Gets project configuration using project.GetProjectConfig()
4. Iterates through files in the project directory using filepath.WalkDir
5. For each file, checks if it matches any of the file filters in the project configuration
6. If a match is found, reads the file content, calculates the relative path, and sends a code summary request to the summarizer service
7. Stores the summary in a map with the relative path as the key
8. After processing all files, generates a summary of all summaries using the summarizer service
9. Writes the summary to a file named "SUMMARY.md"
10. Generates a readme using the summary and writes it to a file named "README_GENERATED.md" or "README.md" if a README.md file already exists

## internal/project/project.go
This part of the code defines the ProjectConfig struct and provides functions to load the project configuration based on the current directory.

1. Defines ProjectConfig struct with fields: FileFilter, ProjectRootFilter, CodePrompt, SummaryPrompt, ReadmePrompt
2. Provides a function GetProjectConfig(currentDirectory string) that iterates through predefined ProjectConfig structs and returns the matching one based on the current directory.
3. Provides helper functions hasFilterFiles and hasRootFilterFile to check if the current directory contains files matching the specified filters.

## internal/summarizer/summarize.go
This part of the code defines the SummarizerService struct and provides functions to generate code summaries and project-level summaries.

1. Defines SummarizerService struct with fields: HelperURL, Model, ApiToken, Network
2. Provides a function CodeSummaryRequest(prompt, content string) that takes a prompt and code content as input and returns a summary.
3. Provides a function SummarizeRequest(prompt, content string) that takes a prompt and content as input and returns a summary.

