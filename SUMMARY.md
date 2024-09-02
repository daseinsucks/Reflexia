This project is a command-line tool that analyzes and summarizes code files within a given directory. It uses an external summarization service to generate summaries for each file and then combines these summaries into a project-level summary. The tool also generates a README file based on the project summary and file summaries.

Project structure:
- cmd/reflexia/reflexia.go
- pkg/project/project.go
- pkg/summarizer/service.go
- pkg/summarizer/summarizer.go

Summary of code parts:

### Main function in cmd/reflexia/reflexia.go
1. Loads environment variables from .env file.
2. Defines flags:
    - "g": valid link for github repository
    - "u": github username for ssh auth
    - "t": github token for ssh auth
3. Defines boolean flags:
    - "c": do not check project root folder specific files such as go.mod or package.json
    - "s": do not create SUMMARY.md and README.md, just print the file summaries
    - "r": do not create README.md
4. Parses flags.
5. Calls processWorkingDirectory function to get directory path.
6. Creates summarizerService instance.
7. Gets project configuration.
8. Calls SummarizeProjectFiles function to get file summaries.
9. If noSummary flag is not set:
    - Calls SummarizeProject function to get project summary.
    - Creates SUMMARY.md file and writes summary content.
10. If noReadme flag is not set:
    - Calls SummarizeReadme function to get readme content.
    - Creates README_GENERATED.md file and writes readme content.

### processWorkingDirectory function in cmd/reflexia/reflexia.go
1. Gets current working directory from environment variable PWD.
2. If ghLink is not empty:
    - Parses github repository url.
    - Creates temporary directory.
    - Clones repository to temporary directory.
3. If flag.Args() is not empty:
    - Sets directory path to first argument.
4. Returns directory path.

### ProjectConfig struct in pkg/project/project.go
1. Fields:
    - FileFilter: []string
    - ProjectRootFilter: []string
    - CodePrompt: string
    - SummaryPrompt: string
    - ReadmePrompt: string
    - RootPath: string

### GetProjectConfig function in pkg/project/project.go
1. Reads configuration from files.
2. Returns ProjectConfig struct.

### hasFilterFiles and hasRootFilterFile functions in pkg/project/project.go
1. Checks if files with given filters exist in directory.
2. Returns true if found, false otherwise.

### SummarizerService struct in pkg/summarizer/service.go
1. Fields:
    - HelperURL
    - Model
    - ApiToken
    - Network
    - LlmOptions

### CodeSummaryRequest and SummarizeRequest functions in pkg/summarizer/service.go
1. Takes prompt and content as input.
2. Calls helper.GenerateContentInstruction with prompt+"```"+content+"```" or prompt+"\n\n"+content.
3. Returns response and error.

### SummarizeReadme, SummarizeProject, and SummarizeProjectFiles functions in pkg/summarizer/summarizer.go
1. SummarizeReadme:
    - Prints "Readme:".
    - Calls SummarizeRequest with projectConfig.ReadmePrompt and summaryContent.
    - Returns the result of SummarizeRequest and nil.
2. SummarizeProject:
    - Initializes summaryContent as an empty string.
    - Iterates through fileMap and appends each file and its summary to summaryContent.
    - Prints "Summary:".
    - Calls SummarizeRequest with projectConfig.SummaryPrompt and summaryContent.
    - Returns the result of SummarizeRequest and nil.
3. SummarizeProjectFiles:
    - Initializes fileMap as an empty map.
    - Calls filepath.WalkDir with projectConfig.RootPath and a callback function.
    - Callback function iterates through each file in the directory.
    - For each file, it checks if the file name matches any of the filters in projectConfig.FileFilter.
    - If a match is found, it reads the file content, calculates the relative path, and calls CodeSummaryRequest with projectConfig.CodePrompt and the file content.
    - The result of CodeSummaryRequest is stored in fileMap.
    - Returns fileMap and nil.

