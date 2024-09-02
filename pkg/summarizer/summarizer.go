package summarizer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflexia/pkg/project"
	"strings"
)

func (s *SummarizerService) SummarizeReadme(
	projectConfig *project.ProjectConfig,
	summaryContent string,
) (string, error) {
	fmt.Println("Readme:")
	readmeContent, err := s.SummarizeRequest(projectConfig.ReadmePrompt, summaryContent)
	if err != nil {
		return "", err
	}
	return readmeContent, nil
}

func (s *SummarizerService) SummarizeProject(
	projectConfig *project.ProjectConfig,
	fileMap map[string]string,
) (string, error) {
	summaryContent := ""
	for file, summary := range fileMap {
		summaryContent += file + "\n" + summary + "\n\n"
	}
	fmt.Println("Summary:")
	// We are doing print stuff since the llm library prints results to the console
	summary, err := s.SummarizeRequest(projectConfig.SummaryPrompt, summaryContent)
	if err != nil {
		return "", err
	}
	fmt.Printf("\n")

	return summary, nil
}

func (s *SummarizerService) SummarizeProjectFiles(
	projectConfig *project.ProjectConfig,
) (map[string]string, error) {

	fileMap := map[string]string{}

	// TODO: add .gitignore support
	if err := filepath.WalkDir(
		projectConfig.RootPath,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			for _, filter := range projectConfig.FileFilter {
				if strings.HasSuffix(d.Name(), filter) {
					content, err := os.ReadFile(path)
					if err != nil {
						return err

					}
					relPath, err := filepath.Rel(projectConfig.RootPath, path)
					if err != nil {
						return err
					}
					fmt.Println(relPath)
					// We are doing print stuff since the llm library prints results to the console
					response, err := s.CodeSummaryRequest(
						projectConfig.CodePrompt, string(content))
					fmt.Printf("\n")
					fileMap[relPath] = response
					if err != nil {
						return err
					}
					break
				}
			}
			return nil
		}); err != nil {
		return nil, err
	}

	return fileMap, nil
}
