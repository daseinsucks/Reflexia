package project

import (
	"errors"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	util "reflexia/internal"
	"strings"
)

var DefaultCodePrompt = `
Describe each code symbol in short form as in the examples below.
Reduct your output as much as possible.
Don't lose original namings.
Always specify environment variables, cli arguments, flags, anything that can configure application behaviour.
Always omit empty sections! Write about section only if they present!
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.

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

func SomeFunc() {
	// Unimplemented
	// TODO: implement SomeFunc
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
	//fmt.Println("TEST")
}
` + "```" + `

Example output:
package foobar
import: fmt, os

struct FooBar:
	- Fields: Foo, Bar
var:
	- zab: "zab"
const:
	- zoob: "zoob"

func SomeFunc(): unimplemented
func main():
	- assigns var foobar with FooBar struct with fields Foo: "foo", Bar: "bar"
	- prints current working directory
	- prints foobar.Foo + foobar.Bar + environment variable "BAZ" + var zab + const zoob
	- prints "boozbazZABraboff" which is reverse of above output

environment variables: "PWD", "BAZ"
` + util.LoadEnv("STOP_WORD") + `

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
import: fmt

func main():
	- prints "Hello world!"
` + util.LoadEnv("STOP_WORD") + `

Provided code:
`

var DefaultPackageSummaryPrompt = `
Based on provided input from the summary of project package files create a markdown summary of what that package code does.
First write a short summary about provided project code summary.
Always specify which environment variables, flags, cmdline arguments can be used for configuration.
Always specify the edgecases of how application can be launched.
Try to guess package name from the file contents or file paths and add it as the markdown header of the summary.
Write out all file names as a project package structure.
Then write summary about every major code part, group it with the markdown subheaders.
Try to explain relations between code entities, try to find unclear places, possibly dead code.		
If unclear places or dead code are not present - don't write anything about their absense.
Try to be clear, concise, and brief.
The main goal is to summarize the logic of the whole package.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`

var DefaultSummaryPrompt = `
Based on provided input from summary of the project files create a markdown summary of what that project code does.
First write a short summary about provided project code summary.
Always specify which environment variables, flags, cmdline arguments can be used for configuration.
Always specify the edgecases of how application can be launched.
Try to guess the project name from the filenames if possible and add it as the markdown header of the summary.
Write out all file names as a project structure.
Then write summary about every major code part, group it with the markdown subheaders.
Try to explain relations between code entities, try to find unclear places, possibly dead code.		
If unclear places or dead code are not present - don't write anything about their absense.
Try to be clear, concise, and brief.
The main goal is to summarize the logic of the whole project.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`
var DefaultReadmePrompt = `
Based on provided input from technical summary of the project create a README.md contents for that project.
Try to guess the project name from filenames or other namings, if present.
Include short project description, possible configuration and/or run instructions.
Try to be clear, concise, and brief.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`

type ProjectConfig struct {
	FileFilter        []string
	ProjectRootFilter []string
	CodePrompt        string
	SummaryPrompt     string
	ReadmePrompt      string
	PackagePrompt     string
	RootPath          string
	ModuleMatch       string
}

func GetProjectConfig(currentDirectory string, lightCheck bool) *ProjectConfig {
	// TODO: read that from configuration file(s)
	for _, config := range []ProjectConfig{
		ProjectConfig{
			FileFilter:        []string{".go"},
			ProjectRootFilter: []string{"go.mod"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			PackagePrompt:     DefaultPackageSummaryPrompt,
			RootPath:          currentDirectory,
			ModuleMatch:       "package_name",
		},

		ProjectConfig{
			FileFilter:        []string{".py"},
			ProjectRootFilter: []string{"requirements.txt"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			PackagePrompt:     DefaultPackageSummaryPrompt,
			RootPath:          currentDirectory,
			ModuleMatch:       "directory",
		},

		ProjectConfig{
			FileFilter:        []string{".ts", ".d.ts", ".tsx"},
			ProjectRootFilter: []string{"package.json"},
			CodePrompt:        DefaultCodePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
			ReadmePrompt:      DefaultReadmePrompt,
			PackagePrompt:     DefaultPackageSummaryPrompt,
			RootPath:          currentDirectory,
			ModuleMatch:       "directory",
		},
	} {
		if lightCheck && hasFilterFiles(currentDirectory, config.FileFilter) {
			return &config
		}
		if hasFilterFiles(currentDirectory, config.FileFilter) &&
			hasRootFilterFile(currentDirectory, config.ProjectRootFilter) {
			return &config
		}
	}

	log.Fatal("failed to detect project language, available languages: go, python, typescript")
	return &ProjectConfig{}
}

func (pc *ProjectConfig) BuildPackageFiles() (map[string][]string, error) {
	packageFileMap := map[string][]string{}
	switch pc.ModuleMatch {
	case "directory":
		if err := util.WalkDirIgnored(
			pc.RootPath,
			func(path string, d fs.DirEntry) error {
				for _, filter := range pc.FileFilter {
					if strings.HasSuffix(d.Name(), filter) {
						key := filepath.Base(filepath.Dir(path))
						relPath, err := filepath.Rel(pc.RootPath, path)
						if err != nil {
							return err
						}

						if _, exists := packageFileMap[key]; !exists {
							packageFileMap[key] = []string{}
						}
						packageFileMap[key] = append(packageFileMap[key], relPath)
					}
				}

				return nil
			}); err != nil {
			return nil, err
		}

	case "package_name":
		if err := util.WalkDirIgnored(
			pc.RootPath,
			func(path string, d fs.DirEntry) error {
				for _, filter := range pc.FileFilter {
					if strings.HasSuffix(d.Name(), filter) {
						fset := token.NewFileSet()
						ast, err := parser.ParseFile(fset, path, nil, 0)
						if err != nil {
							return err
						}
						key := ast.Name.Name
						relPath, err := filepath.Rel(pc.RootPath, path)
						if err != nil {
							return err
						}

						if _, exists := packageFileMap[key]; !exists {
							packageFileMap[key] = []string{}
						}
						packageFileMap[key] = append(packageFileMap[key], relPath)
					}
				}

				return nil
			}); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(pc.ModuleMatch + " module match mode unimplemented")
	}

	return packageFileMap, nil
}

func hasFilterFiles(dirPath string, filters []string) bool {
	found := false

	err := util.WalkDirIgnored(dirPath, func(path string, d fs.DirEntry) error {
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
