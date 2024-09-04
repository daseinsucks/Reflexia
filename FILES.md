# cmd/reflexia/reflexia.go  
package: reflexia/internal  
imports: errors, flag, fmt, io/fs, log, net/url, os, path/filepath, slices, strings, github.com/go-git/go-git/v5, github.com/go-git/go-git/v5/plumbing/transport/http, github.com/joho/godotenv, github.com/tmc/langchaingo/llms, reflexia/pkg/project, reflexia/pkg/summarize  
  
cmdline arguments:  
- -g: valid link for github repository  
- -u: github username for ssh auth  
- -t: github token for ssh auth  
- -l: config filename in project_config to use  
- -c: do not check project root folder specific files such as go.mod or package.json  
- -s: do not create SUMMARY.md and README.md, just print the file summaries  
- -r: do not create README.md  
- -p: do not create README.md for every package in the project  
- -br: overwrite README.md for the root project directory instead of README_GENERATED.md creation  
- -f: Save individual file summary intermediate result to the FILES.md  
- -bp: create README_GENERATED.md if README.md exists in the package directory instead of overwriting  
  
env:  
- PWD: current working directory  
- HELPER_URL: helper URL  
- MODEL: model  
- API_TOKEN: API token  
  
struct Config:  
	- GithubLink: *string  
	- GithubUsername: *string  
	- GithubToken: *string  
	- WithConfigFile: *string  
	- LightCheck: bool  
	- NoSummary: bool  
	- NoReadme: bool  
	- NoPackageSummary: bool  
	- NoBackupRootReadme: bool  
	- WithFileSummary: bool  
	- WithPackageReadmeBackup: bool  
  
func loadEnv(key string) string:  
	- returns the value of the environment variable with the given key, or logs an error and exits if the key is empty  
  
func getProjectStructure(workdir string) (string, error):  
	- returns the project file structure as a string, or an error if there is an issue  
  
func fileMapToString(fileMap map[string]string) string:  
	- converts a map of file names to summaries into a string  
  
func fileMapToMd(fileMap map[string]string) string:  
	- converts a map of file names to summaries into a markdown string  
  
func getReadmePath(workdir string) (string, error):  
	- returns the path to the README.md file, or an error if there is an issue  
  
func writeFile(path, content string) error:  
	- writes the given content to the file at the given path, or returns an error if there is an issue  
  
func processWorkingDirectory(githubLink, githubUsername, githubToken string) (string, error):  
	- returns the working directory, or an error if there is an issue  
  
func initConfig() (*Config, error):  
	- initializes the config struct, parses command-line arguments, and loads environment variables, or returns an error if there is an issue  
  
func main():  
	- initializes the config, processes the working directory, gets the project config, creates a summarizer service, summarizes the code, and writes the results to files, or logs an error if there is an issue  
  
  
  
# internal/util.go  
package: util  
imports: io/fs, log, path/filepath, github.com/crackcomm/go-gitignore  
  
type WalkDirIgnoredFunction:  
	- Takes two arguments: path string, d fs.DirEntry  
	- Returns an error  
  
func WalkDirIgnored(workdir, gitignorePath string, f WalkDirIgnoredFunction) error:  
	- Takes three arguments: workdir string, gitignorePath string, f WalkDirIgnoredFunction  
	- Returns an error  
	- Compiles the .gitignore file using ignore.CompileIgnoreFile(gitignorePath)  
	- Calls filepath.WalkDir(workdir, func(path string, d fs.DirEntry, err error) error) to iterate through the directory tree  
	- Inside the filepath.WalkDir function:  
		- If an error occurs, it returns the error  
		- If the current entry is a directory and its name is ".git", it returns filepath.SkipDir to skip the directory  
		- Calculates the relative path from the workdir to the current path using filepath.Rel(workdir, path)  
		- If the ignoreFile is not nil and it matches the relative path, it returns filepath.SkipDir if the entry is a directory, otherwise it returns nil  
		- Calls the provided function f(path, d) with the current path and directory entry  
	- Returns the error from the filepath.WalkDir function  
  
  
  
# pkg/project/project.go  
package: project  
imports: errors, fmt, go/parser, go/token, io, io/fs, log, os, path/filepath, reflexia/internal, strconv, strings, github.com/pelletier/go-toml/v2  
  
struct ProjectConfig:  
	- Fields: FileFilter, ProjectRootFilter, ModuleMatch, StopWords, CodePrompt, SummaryPrompt, PackagePrompt, ReadmePrompt, RootPath  
func GetProjectConfig(currentDirectory, withConfigFile string, lightCheck bool) (*ProjectConfig, error):  
	- reads all .toml files in the current directory and its subdirectories  
	- returns the first matching ProjectConfig found  
	- if no matching ProjectConfig is found, returns an error  
func (pc *ProjectConfig) BuildPackageFiles() (map[string][]string, error):  
	- builds a map of package names to a list of file paths  
	- returns the map of package names to file paths  
func hasFilterFiles(workdir string, filters []string) bool:  
	- checks if any of the files in the given directory match any of the filters  
	- returns true if a match is found, false otherwise  
func hasRootFilterFile(workdir string, filters []string) bool:  
	- checks if any of the files in the given directory match any of the filters  
	- returns true if a match is found, false otherwise  
  
  
  
# pkg/summarize/summarize.go  
package: summarize  
imports: fmt, io/fs, os, path/filepath, reflexia/internal, reflexia/pkg/project, strings, github.com/JackBekket/hellper/lib/langchain, github.com/tmc/langchaingo/llms  
  
struct SummarizerService:  
	- Fields: HelperURL, Model, ApiToken, Network, LlmOptions  
  
func CodeSummaryRequest(prompt, content string) (string, error):  
	- Takes prompt and content as input  
	- Calls helper.GenerateContentInstruction with provided parameters  
	- Returns response and error  
  
func SummarizeRequest(prompt, content string) (string, error):  
	- Takes prompt and content as input  
	- Calls helper.GenerateContentInstruction with provided parameters  
	- Returns response and error  
  
func SummarizeCode(projectConfig *project.ProjectConfig) (map[string]string, error):  
	- Takes projectConfig as input  
	- Creates a map to store file paths and summaries  
	- Walks through the project directory, ignoring files specified in .gitignore  
	- For each file, it checks if it matches any of the file filters specified in projectConfig  
	- If a match is found, it reads the file content, calculates the relative path, and calls CodeSummaryRequest to generate a summary  
	- Stores the summary in the fileMap along with the relative path  
	- Returns the fileMap and any error encountered during the process  
  
  
  
