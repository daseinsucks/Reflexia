The provided project code is a code summarizer that takes a directory of files as input and generates a summary of the code in each file. It uses an external summarizer service to perform the actual summarization. The project is named "reflexia" based on the filenames provided.

The project uses the following environment variables for configuration:
- PWD
- HELPER_URL
- MODEL
- API_TOKEN

## Project Structure

The project consists of three main parts:
1. cmd/reflexia/reflexia.go
2. internal/project/project.go
3. internal/summarizer/summarize.go

## cmd/reflexia/reflexia.go

This file contains the main function of the project. It loads environment variables using godotenv.Load(), sets up a SummarizerService with values from the environment variables, and gets the project configuration using project.GetProjectConfig().

The main function then iterates through files in the project directory using filepath.WalkDir. For each file, it checks if it matches any of the file filters in the project configuration. If a match is found, it reads the file content, calculates the relative path, and sends a code summary request to the summarizer service. The summary is stored in a map with the relative path as the key.

After processing all files, the main function generates a summary of all summaries using the summarizer service and writes it to a file named "SUMMARY.md". It also generates a readme using the summary and writes it to a file named "README_GENERATED.md" or "README.md" if a README.md file already exists.

## internal/project/project.go

This file contains the ProjectConfig struct and functions to get the project configuration based on the current directory. The ProjectConfig struct has fields for FileFilter, ProjectRootFilter, CodePrompt, SummaryPrompt, and ReadmePrompt.

The GetProjectConfig function iterates through predefined ProjectConfig structs and checks if the current directory contains files matching the FileFilter and ProjectRootFilter. If a match is found, it returns the matching ProjectConfig struct. If no match is found, it logs an error and returns an empty ProjectConfig struct.

The hasFilterFiles and hasRootFilterFile functions are used to check if any files matching the filters exist in the directory.

## internal/summarizer/summarize.go

This file contains the SummarizerService struct and functions to perform code summarization and content summarization. The SummarizerService struct has fields for HelperURL, Model, ApiToken, and Network.

The CodeSummaryRequest and SummarizeRequest functions call helper.GenerateContentInstruction with the appropriate prompt and content. They use the SummarizerService's HelperURL, Model, ApiToken, Network, and llmOptions to perform the summarization.

