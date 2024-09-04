#cmd/reflexia/reflexia.go
package main
  import: fmt, os, strings, slices, filepath, net/url, github.com/go-git/go-git/v5, github.com/go-git/go-git/v5/plumbing/transport/http, github.com/tmc/langchaingo/llms, reflexia/internal, reflexia/pkg/project, reflexia/pkg/summarize
  
  func main():
  	- initializes config struct with default values
  	- parses command line arguments
  	- calls processWorkingDirectory function to get the directory path
  	- calls summarizeService.SummarizeCode function to get file summaries
  	- calls writeFile function to write file summaries to FILES.md
  	- calls summarizeService.SummarizeRequest function to get package summaries
  	- calls writeFile function to write package summaries to README.md
  	- calls summarizeService.SummarizeRequest function to get project summary
  	- calls writeFile function to write project summary to SUMMARY.md
  	- calls summarizeService.SummarizeRequest function to get readme summary
  	- calls writeFile function to write readme summary to README.md
  
  func processWorkingDirectory(ghLink, ghUsername, ghToken string) (string, error):
  	- gets current working directory
  	- if ghLink is provided, clones the repository to a temporary directory
  	- if no ghLink is provided, uses the first command line argument as the directory path
  
  func initConfig():
  	- initializes config struct with default values
  	- parses command line arguments
  	- returns config struct
  
  cli arguments:
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
  
  environment variables:
  	- PWD: current working directory
  	- HELPER_URL: helper URL for summarizer service
  	- MODEL: model for summarizer service
  	- API_TOKEN: API token for summarizer service
  	- STOP_WORD: stop words for summarizer service
  
  

#internal/util.go
package util
  import: io/fs, log, os, path/filepath, ignore, godotenv
  
  type WalkDirIgnoredFunction:
  	- takes path string and fs.DirEntry as arguments
  	- returns error
  
  func WalkDirIgnored(dirPath string, f WalkDirIgnoredFunction):
  	- compiles .gitignore file from dirPath
  	- walks through directory and calls f for each file/directory
  	- skips directories named ".git"
  	- skips files/directories matching .gitignore
  
  func LoadEnv(key string):
  	- loads .env file using godotenv
  	- returns environment variable value for key
  	- logs fatal error if .env file loading fails or key is empty
  
  environment variables: "BAZ"
  

#pkg/project/project.go
package project
  import: errors, go/parser, go/token, io, io/fs, log, os, path/filepath, reflexia/internal, strings
  
  type ProjectConfig:
  	- FileFilter: []string
  	- ProjectRootFilter: []string
  	- CodePrompt: string
  	- SummaryPrompt: string
  	- ReadmePrompt: string
  	- PackagePrompt: string
  	- RootPath: string
  	- ModuleMatch: string
  
  func GetProjectConfig(currentDirectory string, lightCheck bool):
  	- reads configuration from files
  	- returns ProjectConfig struct
  
  func (pc *ProjectConfig) BuildPackageFiles():
  	- returns map[string][]string
  	- builds package files based on ModuleMatch
  
  func hasFilterFiles(dirPath string, filters []string):
  	- returns bool
  	- checks if any of the filters are present in the directory
  
  func hasRootFilterFile(dirPath string, filters []string):
  	- returns bool
  	- checks if any of the filters are present in the directory
  
  

#pkg/summarize/summarize.go
package summarize
  import: fmt, io/fs, os, path/filepath, reflexia/internal, reflexia/pkg/project, strings, github.com/JackBekket/hellper/lib/langchain, github.com/tmc/langchaingo/llms
  
  type SummarizerService:
  	- Fields: HelperURL, Model, ApiToken, Network, LlmOptions
  
  func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error):
  	- takes prompt and content as input
  	- calls helper.GenerateContentInstruction with provided parameters
  	- returns response and error
  
  func (s *SummarizerService) SummarizeRequest(prompt, content string) (string, error):
  	- takes prompt and content as input
  	- calls helper.GenerateContentInstruction with provided parameters
  	- returns response and error
  
  func (s *SummarizerService) SummarizeCode(projectConfig *project.ProjectConfig): (map[string]string, error):
  	- takes projectConfig as input
  	- creates a map to store file paths and summaries
  	- walks through the project directory using util.WalkDirIgnored
  	- for each file, checks if it matches the file filter in projectConfig
  	- if it matches, reads the file content, calculates the relative path, and calls s.CodeSummaryRequest to get the summary
  	- stores the summary in the map with the relative path as the key
  	- returns the map and any error encountered
  
  cli arguments: none
  flags: none
  env vars: none
  

