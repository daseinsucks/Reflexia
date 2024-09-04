pkg/project/project.go
package project
import: errors, go/parser, go/token, io, io/fs, log, os, path/filepath, reflexia/internal, strings

func GetProjectConfig(currentDirectory string, lightCheck bool) *ProjectConfig:
	- reads configuration from configuration file(s)
	- returns ProjectConfig struct with default values for each field
	- if lightCheck is true and hasFilterFiles function returns true, returns the config
	- if hasFilterFiles and hasRootFilterFile functions return true, returns the config
	- if no matching config is found, logs an error and returns an empty ProjectConfig struct

func (pc *ProjectConfig) BuildPackageFiles() (map[string][]string, error):
	- returns a map of package names to a list of file paths
	- based on the ModuleMatch field, it either uses directory or package_name mode
	- in directory mode, it walks through the directory and collects all files matching the FileFilter
	- in package_name mode, it parses the Go files and extracts the package name from the AST
	- returns the map of package names to file paths, or an error if any

func hasFilterFiles(dirPath string, filters []string) bool:
	- checks if any files matching the given filters exist in the directory
	- returns true if any matching files are found, false otherwise

func hasRootFilterFile(dirPath string, filters []string) bool:
	- checks if any files matching the given filters exist in the directory
	- returns true if any matching files are found, false otherwise



pkg/summarize/summarize.go
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
	- for each file, checks if it matches the file filter
	- if it matches, reads the file content, calculates the relative path, and calls s.CodeSummaryRequest to get the summary
	- stores the summary in the map with the relative path as the key
	- returns the map and any error encountered

environment variables: none
cli arguments: none
flags: none


cmd/reflexia/reflexia.go
package main
import: fmt, errors, flag, log, net/url, os, path/filepath, slices, strings, github.com/go-git/go-git/v5, github.com/go-git/go-git/v5/plumbing/transport/http, github.com/tmc/langchaingo/llms, reflexia/internal, reflexia/pkg/project, reflexia/pkg/summarize

type Config:
	- GithubLink: *string
	- GithubUsername: *string
	- GithubToken: *string
	- LightCheck: bool
	- NoSummary: bool
	- NoReadme: bool
	- NoPackageSummary: bool
	- NoBackupRootReadme: bool
	- WithFileSummary: bool
	- WithPackageReadmeBackup: bool

func main():
	- initializes config using initConfig()
	- calls processWorkingDirectory() to get the directory path
	- creates a summarizeService using summarize.SummarizerService
	- gets projectConfig using project.GetProjectConfig()
	- calls summarizeService.SummarizeCode() to get fileMap
	- if WithFileSummary is true, writes fileMap to FILES.md
	- if not NoSummary, calls summarizeService.SummarizeRequest() to get summaryContent and writes it to SUMMARY.md
	- if not NoReadme, calls summarizeService.SummarizeRequest() to get readmeContent and writes it to README.md

func processWorkingDirectory():
	- returns directory path based on provided github link or first command line argument

func initConfig():
	- initializes config using flag package
	- returns config

func fileMapToString():
	- converts fileMap to string

func getReadmePath():
	- returns readme filename based on directory path

func writeFile():
	- writes content to file

environment variables: "PWD", "GH_TOKEN"
cli arguments:
	- -g: valid link for github repository
	- -u: github username for ssh auth
	- -t: github token for ssh auth
	- -c: do not check project root folder specific files such as go.mod or package.json
	- -s: do not create SUMMARY.md and README.md, just print the file summaries
	- -r: do not create README.md
	- -p: do not create README.md for every package in the project
	- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation
	- -f: Save individual file summary intermediate result to the FILES.md
	- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting


internal/util.go
package util
import: io/fs, log, os, path/filepath, ignore, godotenv

type WalkDirIgnoredFunction:
	- Takes path string and fs.DirEntry as arguments
	- Returns error

func WalkDirIgnored(dirPath string, f WalkDirIgnoredFunction):
	- Compiles .gitignore file from dirPath
	- Walks through directory and calls f for each file/directory
	- Skips .git directory and files/directories matching .gitignore

func LoadEnv(key string):
	- Loads .env file using godotenv
	- Returns environment variable value for key
	- Logs fatal error if .env file loading fails or key is empty

environment variables:
	- .env file



