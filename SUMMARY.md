This project code is designed to generate code summaries for a given project directory. It utilizes a summarizer service to process the code and create a summary of the project's code. The project uses environment variables for configuration, including HELPER_URL, MODEL, and API_TOKEN.

## SummarizerService
The SummarizerService struct is responsible for handling the summarization process. It has four fields: HelperURL, Model, ApiToken, and Network. The main functions of this struct are CodeSummaryRequest and SummarizeRequest.

### CodeSummaryRequest
This function takes a prompt and code content as input and returns a string containing the summarized code. It calls the helper.GenerateContentInstruction function with the prompt and code content, using the HelperURL, Model, ApiToken, Network, and llmOptions.

### SummarizeRequest
This function takes a prompt and content as input and returns a string containing the summarized content. It calls the helper.GenerateContentInstruction function with the prompt and content, using the HelperURL, Model, ApiToken, Network, and llmOptions.

## ProjectConfig
The ProjectConfig struct is used to define the configuration for the project. It has five fields: FileFilter, ProjectRootFilter, CodePrompt, SummaryPrompt, and ReadmePrompt. The GetProjectConfig function takes the current directory as input and returns a matching ProjectConfig struct based on the file filters and root filters.

## File Processing
The main function in cmd/reflexia/reflexia.go loads environment variables using godotenv.Load and sets up the SummarizerService with values from these variables. It then gets the project configuration using project.GetProjectConfig and iterates through files in the project directory using filepath.WalkDir. For each file, it checks if it matches any of the file filters in the project configuration. If a match is found, it reads the file content, calculates the relative path, and sends a code summary request to the summarizer service. The summary is stored in a map with the relative path as the key.

After processing all files, the main function generates a summary of all summaries using the summarizer service and writes it to a file named "SUMMARY.md". It also generates a readme using the summary and writes it to a file named "README_GENERATED.md" or "README.md" if a README.md file already exists.

