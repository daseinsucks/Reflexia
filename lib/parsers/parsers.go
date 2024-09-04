package parsers

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	util "reflexia/lib/internal"
	prompt "reflexia/lib/prompts"
	"strings"

	helper "github.com/JackBekket/hellper/lib/langchain"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/tmc/langchaingo/llms"

	"github.com/go-git/go-git"
)

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
type SummarizerService struct {
	HelperURL  string
	Model      string
	ApiToken   string
	Network    string
	LlmOptions []llms.CallOption
}

func processWorkingDirectory(ghLink, ghUsername, ghToken string) (string, error) {
	dirPath := util.LoadEnv("PWD")

	if ghLink != "" {
		u, err := url.ParseRequestURI(ghLink)
		if err != nil {
			return "", err
		}

		sPath := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
		if len(sPath) != 2 {
			return "", errors.New("github repository url does not have two path elements")
		}

		tempDirEl := []string{dirPath, "temp"}
		tempDirEl = append(tempDirEl, sPath...)
		tempDir := filepath.Join(tempDirEl...)

		dirPath = tempDir

		if _, err := os.Stat(dirPath); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(dirPath, os.FileMode(0755)); err != nil {
					return "", err
				}

				cloneOptions := git.CloneOptions{
					URL:               ghLink,
					RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
					Depth:             1,
				}
				if ghUsername != "" && ghToken != "" {
					cloneOptions.Auth = &http.BasicAuth{
						Username: ghUsername,
						Password: ghToken,
					}
				}

				if _, err := git.PlainClone(dirPath, false, &cloneOptions); err != nil {
					if err := os.RemoveAll(dirPath); err != nil {
						return "", err
					}
					return "", err
				}

			} else {
				return "", err
			}
		}
	} else if len(flag.Args()) > 0 {
		dirPath = flag.Arg(0)
		if _, err := os.Stat(dirPath); err != nil {
			return "", err
		}
	}

	return dirPath, nil
}

// pass dirPath here
func GetProjectConfig(currentDirectory string, lightCheck bool) *ProjectConfig {
	// TODO: read that from configuration file(s)
	for _, config := range []ProjectConfig{
		ProjectConfig{
			FileFilter:        []string{".go"},
			ProjectRootFilter: []string{"go.mod"},
			CodePrompt:        prompt.DefaultCodePrompt,
			SummaryPrompt:     prompt.DefaultSummaryPrompt,
			ReadmePrompt:      prompt.DefaultReadmePrompt,
			PackagePrompt:     prompt.DefaultPackageSummaryPrompt,
			RootPath:          currentDirectory,
			ModuleMatch:       "package_name",
		},

		ProjectConfig{
			FileFilter:        []string{".py"},
			ProjectRootFilter: []string{"requirements.txt"},
			CodePrompt:        prompt.DefaultCodePrompt,
			SummaryPrompt:     prompt.DefaultSummaryPrompt,
			ReadmePrompt:      prompt.DefaultReadmePrompt,
			PackagePrompt:     prompt.DefaultPackageSummaryPrompt,
			RootPath:          currentDirectory,
			ModuleMatch:       "directory",
		},

		ProjectConfig{
			FileFilter:        []string{".ts", ".d.ts", ".tsx"},
			ProjectRootFilter: []string{"package.json"},
			CodePrompt:        prompt.DefaultCodePrompt,
			SummaryPrompt:     prompt.DefaultSummaryPrompt,
			ReadmePrompt:      prompt.DefaultReadmePrompt,
			PackagePrompt:     prompt.DefaultPackageSummaryPrompt,
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

// then put projectConfig here
func (s *SummarizerService) SummarizeCode(
	projectConfig *ProjectConfig,
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

func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
		s.LlmOptions...,
	)

	return response, err
}
