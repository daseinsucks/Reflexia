package project

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DefaultFilePrompt = `
	Summarize every major declaration (functions, structs, viable constants) from provided code.
	Note that your output will be reused later for summarizing over all the project files.
	You can use less human readable language, since you producing intermediate result.
	The main goal is to summarize the logic of the whole project, and provide a brief output, since we trying to save space in your (llm) context.
	Try to be clear, concise, and brief.
`
var DefaultSummaryPrompt = ""

type ProjectConfig struct {
	FileFilter        []string
	ProjectRootFilter []string
	FilePrompt        string
	SummaryPrompt     string
}

func GetProjectConfig(currentDirectory string) *ProjectConfig {

	for _, config := range []ProjectConfig{
		ProjectConfig{
			FileFilter:        []string{".go"},
			ProjectRootFilter: []string{"go.mod"},
			FilePrompt:        DefaultFilePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
		},

		ProjectConfig{
			FileFilter:        []string{".py"},
			ProjectRootFilter: []string{"requirements.txt"},
			FilePrompt:        DefaultFilePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
		},

		ProjectConfig{
			FileFilter:        []string{".ts", ".d.ts"},
			ProjectRootFilter: []string{"package.json"},
			FilePrompt:        DefaultFilePrompt,
			SummaryPrompt:     DefaultSummaryPrompt,
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
