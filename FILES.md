# pkg/summarize/summarize.go  
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
	- walks through the project directory, ignoring files specified in .gitignore  
	- for each file, it checks if the file extension matches any of the file filters specified in projectConfig  
	- if a match is found, it reads the file content, calculates the relative path, and calls s.CodeSummaryRequest to generate a summary  
	- stores the summary in the fileMap along with the relative path  
	- returns the fileMap and any error encountered during the process  
  
environment variables: none  
config files: .gitignore  
cli arguments: none  
flags: none  
  
  
# cmd/reflexia/reflexia.go  
package main  
import: fmt, errors, flag, os, path/filepath, slices, strings, github.com/go-git/go-git/v5, github.com/go-git/go-git/v5/plumbing/transport/http, github.com/joho/godotenv, github.com/tmc/langchaingo/llms, reflexia/internal, reflexia/pkg/project, reflexia/pkg/summarize  
  
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
	- calls processWorkingDirectory() to get the workdir  
	- calls summarizeService.SummarizeCode() to get fileMap  
	- calls writeFile() to write fileMap to FILES.md if WithFileSummary is true  
	- calls summarizeService.SummarizeRequest() to get pkgSummaryContent for each package  
	- calls writeFile() to write pkgSummaryContent to the appropriate README.md file  
	- calls summarizeService.SummarizeRequest() to get summaryContent  
	- calls writeFile() to write summaryContent to SUMMARY.md  
	- calls summarizeService.SummarizeRequest() to get readmeContent  
	- calls writeFile() to write readmeContent to README.md  
  
func initConfig():  
	- loads environment variables using godotenv.Load()  
	- initializes config struct  
	- sets default values for config fields  
	- parses command-line arguments using flag.Parse()  
  
func processWorkingDirectory():  
	- gets workdir from environment variable PWD  
	- if GithubLink is provided, clones the repository to the workdir  
	- if no GithubLink is provided, uses the first command-line argument as the workdir  
  
func fileMapToString():  
	- converts fileMap to a string  
  
func fileMapToMd():  
	- converts fileMap to a markdown string  
  
func getReadmePath():  
	- returns the path to the README.md file  
  
func writeFile():  
	- writes content to the specified path  
  
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
	- HELPER_URL: helper URL  
	- MODEL: model  
	- API_TOKEN: API token  
  
config files:  
	- .env: contains environment variables  
  
  
  
# internal/util.go  
package util  
import: io/fs, log, os, path/filepath, ignore  
  
type WalkDirIgnoredFunction:  
	- takes path string and fs.DirEntry as arguments  
	- returns error  
  
func WalkDirIgnored:  
	- takes workdir, gitignorePath, and f as arguments  
	- compiles .gitignore file using ignore.CompileIgnoreFile  
	- walks through directory using filepath.WalkDir  
	- skips .git directory  
	- calculates relative path using filepath.Rel  
	- checks if path matches .gitignore using ignoreFile.MatchesPath  
	- calls f with path and fs.DirEntry if not ignored  
	- returns error  
  
func LoadEnv:  
	- takes key as argument  
	- returns value of environment variable with key  
	- logs fatal error if key is empty  
  
environment variables: key  
  
  
# pkg/project/project.go  
package project  
import: errors, go/parser, go/token, io, io/fs, log, os, path/filepath, util, strings, github.com/pelletier/go-toml/v2  
  
type ProjectConfig:  
	- Fields: FileFilter, ProjectRootFilter, ModuleMatch, StopWords, CodePrompt, SummaryPrompt, PackagePrompt, ReadmePrompt, RootPath  
  
func GetProjectConfig(currentDirectory string, lightCheck bool) (*ProjectConfig, error):  
	- walks through all .toml files in the current directory and its subdirectories  
	- returns the first ProjectConfig that matches the given criteria  
	- if no matching ProjectConfig is found, returns an error  
  
func (pc *ProjectConfig) BuildPackageFiles() (map[string][]string, error):  
	- builds a map of package names to a list of file paths  
	- the map is built based on the ModuleMatch field of the ProjectConfig  
	- if ModuleMatch is "directory", it walks through the directory and collects all files that match the FileFilter  
	- if ModuleMatch is "package_name", it parses the Go code and extracts the package name from the AST  
	- returns the map of package names to file paths  
  
func hasFilterFiles(workdir string, filters []string) bool:  
	- checks if any files in the given directory match the given filters  
	- returns true if a matching file is found, false otherwise  
  
func hasRootFilterFile(workdir string, filters []string) bool:  
	- checks if any files in the given directory match the given filters  
	- returns true if a matching file is found, false otherwise  
  
environment variables: none  
config files: project_config/*.toml  
  
  
