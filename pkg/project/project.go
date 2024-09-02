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
Your main task is to compress and simplify the provided code logic and save output space.
Reduct your output as much as possible, yet don't lose original namings.
Do not lose the details, such as trivial code logic or edge case processing.
Always specify environment variables, cli arguments, flags, anything that can configure application behaviour.
Always omit empty sections! Write about section only if they present!
Always omit already written external references!
It is mandatory to prepend the <end_of_summary> at the very end of your output.

Example input:
` + "```" + `
package foobar
import (
	"fmt"
	"os"
)

var zab = "zab"
const zoob = "zoob"

struct FooBar {
	Foo string
	Bar string
}
func main() {
	foobar := FooBar{
		Foo: "foo",
		Bar: "bar"
	}
	// Print current directory
	fmt.Println(os.Getenv("PWD"))
	fmt.Println(foobar.Foo + foobar.Bar + os.Getenv("BAZ") + zab + zoob)
	// Reverse of above output
	fmt.Println("boozbazZABraboff")
}
` + "```" + `

Example output:
package foobar

struct FooBar:
	- Fields: Foo, Bar
var:
	- zab: "zab"
const:
	- zoob: "zoob"
func main():
	- assigns var foobar with FooBar struct with fields Foo: "foo", Bar: "bar"
	- prints current working directory
	- prints foobar.Foo + foobar.Bar + environment variable "BAZ" + var zab + const zoob
	- prints "boozbazZABraboff" which is reverse of above output

environment variables: "PWD", "BAZ"
external references: fmt.Println, os.Getenv
<end_of_summary>

Example input:
` + "```" + `
package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
}
` + "```" + `

Example output:
package main

func main():
	- prints "Hello world!"

external references: fmt.Println
<end_of_summary>

Provided code:
`

const DefaultSummaryPrompt = `
Based on provided input from summary of project files create a markdown summary of what that project code does.
First write a short summary about provided project code summary.
Always specify which environment variables, flags, cmdline arguments can be used for configuration.
Always specify the edgecases of how application can be launched.
Try to guess the project name from the filenames if possible.
Write out all file names as a project structure.
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
	RootPath          string
}

func GetProjectConfig(currentDirectory string) *ProjectConfig {
	// TODO: read that from configuration file(s)
	for _, config := range []ProjectConfig{
		ProjectConfig{
			FileFilter:        []string{".go"},
			ProjectRootFilter: []string{"go.mod"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			RootPath:          currentDirectory,
		},

		ProjectConfig{
			FileFilter:        []string{".py"},
			ProjectRootFilter: []string{"requirements.txt"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			RootPath:          currentDirectory,
		},

		ProjectConfig{
			FileFilter:        []string{".ts", ".d.ts"},
			ProjectRootFilter: []string{"package.json"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			RootPath:          currentDirectory,
		},
	} {
		if hasFilterFiles(currentDirectory, config.FileFilter) ||
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
			if !os.IsNotExist(err) {
				panic(err)
			}
			return false
		}
	}

	return true
}
