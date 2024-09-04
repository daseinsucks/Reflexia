# reflexia

This package provides a tool for summarizing code and generating documentation for Go projects. It can be used to create a summary of the code in a given directory, as well as a README file for the project. The tool can also be used to generate summaries for individual files within the project.

Environment variables: "PWD", "GH_TOKEN"
CLI arguments:
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

Project package structure:
- cmd/reflexia/reflexia.go
- reflexia/internal
- reflexia/pkg/project
- reflexia/pkg/summarize

## Main logic

The main function initializes the configuration, gets the directory path, creates a summarize service, gets the project configuration, and calls the summarize service to get the file map. If the WithFileSummary flag is true, it writes the file map to a file called FILES.md. If the NoSummary flag is false, it calls the summarize service to get the summary content and writes it to a file called SUMMARY.md. If the NoReadme flag is false, it calls the summarize service to get the readme content and writes it to a file called README.md.

The processWorkingDirectory function returns the directory path based on the provided GitHub link or the first command-line argument. The initConfig function initializes the configuration using the flag package and returns the configuration. The fileMapToString function converts the file map to a string. The getReadmePath function returns the readme filename based on the directory path. The writeFile function writes the content to a file.

## Edge cases

The application can be launched with various command-line arguments to configure its behavior. For example, the -s flag can be used to skip the creation of SUMMARY.md and README.md files, and only print the file summaries. The -r flag can be used to skip the creation of the README.md file, and the -p flag can be used to skip the creation of README.md files for each package in the project. The -br flag can be used to overwrite the existing README.md file in the root project directory instead of creating a new file called README_GENERATED.md. The -f flag can be used to save the individual file summary intermediate result to a file called FILES.md, and the -bp flag can be used to create a README_GENERATED.md file if a README.md file already exists in the package directory instead of overwriting it.

