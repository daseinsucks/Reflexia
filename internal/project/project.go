package project

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const DefaultCodePrompt = `
Compress the provided code logic to save the llm context space.
Do not lose the details, such as trivial code logic or edge case processing.
If you see any possibility to include the information about external configuration such as environment variables or config files - include them in a separate list like on example below.
Don't confuse the struct fields and os.Getenv calls or something similar when doing such lists.
Don't write the group section if it's empty, save the space!
Your output will be used for summarizing the whole project file by file.
It is mandatory to prepend the <end_of_summary> at the very end of your output.

Example input:
` + "```" + `
package main
import (
	"fmt"
	"os"
)
struct FooBar {
	Foo string
	Bar string
}
func main() {
	foobar := FooBar{
		Foo: "foo",
		Bar: "bar"
	}
	fmt.Println(foobar.Foo + foobar.Bar + os.Getenv("BAZ"))
}
` + "```" + `

Example output:
struct FooBar:
	- Fields: Foo, Bar
func main():
	- assigns var foobar with FooBar struct with fields Foo: "foo", Bar: "bar"
	- prints foobar.Foo + foobar.Bar + environment variable "BAZ"

external references:
	- fmt.Println
	- os.Getenv
environment variables:
	- "BAZ"
<end_of_summary>

Provided code:
`

const DefaultSummaryPrompt = `
Based on provided input from summary of project files create a markdown summary of what that project code does.
First write a short summary about provided project code summary.
Always specify which environment variables can be used for configuration.
Always include filenames. Try to guess the project name from them if possible.
Then write summary about every major code part, group it with markdown headers.
Try to explain relations between code entities, try to find unclear places, possibly dead code.		
If unclear places or dead code are not present - don't write anything about their absense.
Try to be clear, concise, and brief.
The main goal is to summarize the logic of the whole project.
It is mandatory to prepend the <end_of_summary> at the very end of your output.
`
const DefaultReadmePrompt = `
Based on provided input from technical summary of the project create a README.md contents for that project.
Try to guess the project name from filenames or other namings, if present.
Include short project description, possible configuration and/or run instructions.
Try to be clear, concise, and brief.
It is mandatory to prepend the <end_of_summary> at the very end of your output.
`

type ProjectConfig struct {
	FileFilter        []string
	ProjectRootFilter []string
	CodePrompt        string
	SummaryPrompt     string
	ReadmePrompt      string
}

func GetProjectConfig(currentDirectory string) *ProjectConfig {

	for _, config := range []ProjectConfig{
		ProjectConfig{
			FileFilter:        []string{".go"},
			ProjectRootFilter: []string{"go.mod"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
		},

		ProjectConfig{
			FileFilter:        []string{".py"},
			ProjectRootFilter: []string{"requirements.txt"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
		},

		ProjectConfig{
			FileFilter:        []string{".ts", ".d.ts"},
			ProjectRootFilter: []string{"package.json"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
		},
	} {
		if hasFilterFiles(currentDirectory, config.FileFilter) &&
			hasRootFilterFile(currentDirectory, config.ProjectRootFilter) {
			return &config
		}
	}

	log.Fatal("failed to detect project language, available languages: go, python, typescript")
	return &ProjectConfig{}
}

func hasFilterFiles(dirPath string, filters []string) bool {
	found := false
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, filter := range filters {
			if strings.HasSuffix(d.Name(), filter) {
				found = true
				return io.EOF
			}
		}
		return nil
	})
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return found
}

func hasRootFilterFile(dirPath string, filters []string) bool {
	for _, filter := range filters {
		if _, err := os.Stat(filepath.Join(dirPath, filter)); err != nil {
			return false
		}
	}

	return true
}
