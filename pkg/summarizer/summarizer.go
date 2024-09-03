package summarizer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	util "reflexia/internal"
	"reflexia/pkg/project"
	"strings"
)

func (s *SummarizerService) CreateReadme(
	projectConfig *project.ProjectConfig,
	summaryContent string,
) (string, error) {
	fmt.Println("Readme:")
	// We are doing print stuff since the llm library prints results to the console
	readmeContent, err := s.SummarizeRequest(projectConfig.ReadmePrompt, summaryContent)
	if err != nil {
		return "", err
	}
	return readmeContent, nil
}

func (s *SummarizerService) SummarizeProject(
	projectConfig *project.ProjectConfig,
	codeContent string,
) (string, error) {
	fmt.Println("Summary:")
	// We are doing print stuff since the llm library prints results to the console
	summary, err := s.SummarizeRequest(projectConfig.SummaryPrompt, codeContent)
	if err != nil {
		return "", err
	}

	return summary, nil
}

func (s *SummarizerService) SummarizeCode(
	projectConfig *project.ProjectConfig,
) (map[string]string, error) {

	fileMap := map[string]string{}

	if err := util.WalkDirIgnored(
		projectConfig.RootPath,
		func(path string, d fs.DirEntry) error {
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
