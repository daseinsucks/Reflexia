The provided project code is a command-line tool that summarizes the contents of a given project directory. It takes a GitHub link or a directory path as input and uses an external API to generate summaries of the project's code, project, and README files. The tool also creates two output files: README_GENERATED.md and SUMMARY.md, which contain the generated summaries.

Project structure:
- cmd/reflexia/reflexia.go
- pkg/project/project.go
- pkg/summarizer/service.go
- pkg/summarizer/summarizer.go

Environment variables: PWD, HELPER_URL, MODEL, API_TOKEN, GH_TOKEN
CLI arguments: -g (github link), -u (github username), -t (github token)

## SummarizerService
The SummarizerService struct contains the necessary information for interacting with the external API, including the helper URL, model, API token, network, and LLM options.

## processWorkingDirectory
This function takes a GitHub link, GitHub username, and GitHub token as input and returns the directory path of the project. If a GitHub link is provided, it parses the link, extracts the repository path, creates a temporary directory, and clones the repository using the provided credentials. If no GitHub link is provided, it takes the first command-line argument as the directory path.

## loadEnv
This function loads the value of a given environment variable.

## main
The main function loads environment variables, parses command-line arguments, calls processWorkingDirectory to get the directory path, creates a SummarizerService instance, gets the project configuration, summarizes the project files, project, and README, and writes the summaries to the output files.

## GetProjectConfig
This function iterates through predefined ProjectConfig structs and checks if the current directory contains files matching the FileFilter and ProjectRootFilter. If a match is found, it returns the corresponding ProjectConfig struct. Otherwise, it logs an error and returns an empty ProjectConfig struct.

## SummarizeReadme
This function takes a ProjectConfig struct and a summary content string as input and calls the SummarizeRequest method with the project's ReadmePrompt and the summary content. It returns the result and any error encountered.

## SummarizeProject
This function takes a ProjectConfig struct and a file map as input. It initializes a summary content string, iterates through the file map, appends each file and its summary to the summary content, calls the SummarizeRequest method with the project's SummaryPrompt and the summary content, and returns the result and any error encountered.

## SummarizeProjectFiles
This function takes a ProjectConfig struct as input and initializes a file map. It calls filepath.WalkDir with the project's RootPath, iterates through the files, filters them based on the project's FileFilter, reads the file content, calculates the relative path, calls the CodeSummaryRequest method with the project's CodePrompt and the file content, and stores the response in the file map. Finally, it returns the file map and any error encountered.

