package summarize

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	util "reflexia/internal"
	"reflexia/pkg/project"
	"strings"

	helper "github.com/JackBekket/hellper/lib/langchain"
	"github.com/tmc/langchaingo/llms"
)

type SummarizerService struct {
	HelperURL  string
	Model      string
	ApiToken   string
	Network    string
	LlmOptions []llms.CallOption
}

func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
		s.LlmOptions...,
	)

	return response, err
}

func (s *SummarizerService) SummarizeRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"\n\n"+content,
		s.Model, s.ApiToken, s.Network,
		s.LlmOptions...,
	)

	return response, err
}

func (s *SummarizerService) SummarizeCode(
	projectConfig *project.ProjectConfig,
) (map[string]string, error) {

	fileMap := map[string]string{}

	if err := util.WalkDirIgnored(
		projectConfig.RootPath,
		filepath.Join(projectConfig.RootPath, ".gitignore"),
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
