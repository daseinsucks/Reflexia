# Reflexia

This project is a code summarizer that takes a Go project as input and generates summaries for each file, package, and the entire project. It uses a combination of environment variables, command-line arguments, and configuration files to determine the behavior and output of the summarizer.

Project Structure:
- pkg/project/project.go
- pkg/summarize/summarize.go
- cmd/reflexia/reflexia.go
- internal/util.go

Summary of Code Parts:

1. Project Configuration (pkg/project/project.go):
   - The `ProjectConfig` struct defines the configuration for the summarizer, including file filters, project root filters, module match, stop words, and prompts for code, summary, package, readme, and root path.
   - The `GetProjectConfig` function walks through all .toml files in the current directory and its subdirectories to find the first matching ProjectConfig.
   - The `BuildPackageFiles` function builds a map of package names to a list of file paths based on the ModuleMatch field of the ProjectConfig.

2. Summarization Service (pkg/summarize/summarize.go):
   - The `SummarizerService` struct contains fields for helper URL, model, API token, network, and LLM options.
   - The `CodeSummaryRequest` and `SummarizeRequest` functions take prompt and content as input and call helper.GenerateContentInstruction to generate a response.
   - The `SummarizeCode` function takes a projectConfig as input, walks through the project directory, and calls `CodeSummaryRequest` to generate summaries for each file.

3. Main Application (cmd/reflexia/reflexia.go):
   - The `Config` struct defines the configuration for the main application, including GitHub link, username, token, light check, no summary, no readme, no package summary, no backup root readme, with file summary, and with package readme backup.
   - The `main` function initializes the config, gets the workdir, calls `summarizeService.SummarizeCode` to get fileMap, and writes the fileMap to FILES.md if WithFileSummary is true.
   - It also calls `summarizeService.SummarizeRequest` to get pkgSummaryContent for each package and write it to the appropriate README.md file.

4. Utility Functions (internal/util.go):
   - The `WalkDirIgnored` function takes a workdir, gitignorePath, and f as arguments, compiles the .gitignore file, walks through the directory, and calls f with the path and fs.DirEntry if not ignored.
   - The `LoadEnv` function takes a key as argument and returns the value of the environment variable with the key.

Edge Cases:

- If no matching ProjectConfig is found, the summarizer will return an error.
- If no GitHub link is provided, the summarizer will use the first command-line argument as the workdir.
- If the .gitignore file is not found, the summarizer will not ignore any files.

Environment Variables:

- PWD: current working directory
- HELPER_URL: helper URL
- MODEL: model
- API_TOKEN: API token

Config Files:

- project_config/*.toml
- .env
- .gitignore

CLI Arguments:

- -g, --github-link: valid link for github repository
- -u, --github-username: github username for ssh auth
- -t, --github-token: github token for ssh auth
- -c: do not check project root folder specific files such as go.mod or package.json
- -s: do not create SUMMARY.md and README.md, just print the file summaries
- -r: do not create README.md
- -p: do not create README.md for every package in the project
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
- -f: Save individual file summary intermediate result to the FILES.md
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

