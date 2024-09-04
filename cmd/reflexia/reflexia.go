package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/tmc/langchaingo/llms"

	util "reflexia/internal"
	"reflexia/pkg/project"
	"reflexia/pkg/summarize"
)

type Config struct {
	GithubLink              *string
	GithubUsername          *string
	GithubToken             *string
	LightCheck              bool
	NoSummary               bool
	NoReadme                bool
	NoPackageSummary        bool
	NoBackupRootReadme      bool
	WithFileSummary         bool
	WithPackageReadmeBackup bool
}

func main() {
	config, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	dirPath, err := processWorkingDirectory(
		*config.GithubLink, *config.GithubUsername, *config.GithubToken)
	if err != nil {
		log.Fatal(err)
	}

	summarizeService := &summarize.SummarizerService{
		HelperURL: util.LoadEnv("HELPER_URL"),
		Model:     util.LoadEnv("MODEL"),
		ApiToken:  util.LoadEnv("API_TOKEN"),
		Network:   "local",
		LlmOptions: []llms.CallOption{
			llms.WithStopWords(
				[]string{
					util.LoadEnv("STOP_WORD"),
				},
			),
			llms.WithRepetitionPenalty(0.7),
		},
	}
	projectConfig := project.GetProjectConfig(dirPath, config.LightCheck)

	fileMap, err := summarizeService.SummarizeCode(projectConfig)
	if err != nil {
		log.Fatal(err)
	}
	if config.WithFileSummary {
		if err := writeFile(
			filepath.Join(dirPath, "FILES.md"),
			fileMapToMd(fileMap),
		); err != nil {
			log.Fatal(err)
		}
	}

	if !config.NoSummary {
		if !config.NoPackageSummary {
			pkgFiles, err := projectConfig.BuildPackageFiles()
			if err != nil {
				log.Fatal(err)
			}

			for pkg, files := range pkgFiles {
				fmt.Printf("\n%s:\n", pkg)

				pkgFileMap := map[string]string{}
				pkgDir := ""
				for file, content := range fileMap {
					if slices.Contains(files, file) {
						pkgFileMap[file] = content
						if pkgDir == "" {
							pkgDir = filepath.Join(dirPath, filepath.Dir(file))
						}
					}
				}

				pkgSummaryContent, err := summarizeService.SummarizeRequest(
					projectConfig.PackagePrompt, fileMapToString(pkgFileMap),
				)
				if err != nil {
					log.Fatal(err)
				}

				readmeFilename := "README.md"
				if config.WithPackageReadmeBackup {
					readmeFilename, err = getReadmePath(pkgDir)
					if err != nil {
						log.Fatal(err)
					}
				}
				if err := writeFile(
					filepath.Join(pkgDir, readmeFilename),
					pkgSummaryContent,
				); err != nil {
					log.Fatal(err)
				}

			}
		}

		fmt.Println("\nSummary: ")

		summaryContent, err := summarizeService.SummarizeRequest(
			projectConfig.SummaryPrompt, fileMapToString(fileMap),
		)
		if err != nil {
			log.Fatal(err)
		}
		if err := writeFile(
			filepath.Join(dirPath, "SUMMARY.md"), summaryContent); err != nil {
			log.Fatal(err)
		}

		if !config.NoReadme {
			fmt.Println("\nReadme: ")

			readmeContent, err := summarizeService.SummarizeRequest(
				projectConfig.ReadmePrompt, summaryContent,
			)
			if err != nil {
				log.Fatal(err)
			}

			readmeFilename := "README.md"
			if !config.NoBackupRootReadme {
				readmeFilename, err = getReadmePath(dirPath)
				if err != nil {
					log.Fatal(err)
				}
			}
			if err := writeFile(
				filepath.Join(dirPath, readmeFilename), readmeContent); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func fileMapToString(fileMap map[string]string) string {
	content := ""
	for file, summary := range fileMap {
		content += file + "\n" + summary + "\n\n"
	}
	return content
}

func fileMapToMd(fileMap map[string]string) string {
	content := ""
	for file, summary := range fileMap {
		content += "#" + file + "\n"
		summaryNLEscaped := strings.ReplaceAll(summary, "\n", "\n  ")
		content += summaryNLEscaped + "\n\n"
	}
	return content
}
func getReadmePath(dirPath string) (string, error) {
	readmeFilename := "README_GENERATED.md"
	if _, err := os.Stat(filepath.Join(dirPath, "README.md")); err != nil {
		if os.IsNotExist(err) {
			readmeFilename = "README.md"
		} else {
			return "", err
		}
	}
	return readmeFilename, nil
}

func writeFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
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

func initConfig() (*Config, error) {
	// Temporary moved to util.LoadEnv, until we move stop words to the toml files
	// if err := godotenv.Load(); err != nil {
	// 	return nil, err
	// }

	config := Config{}

	config.GithubLink = flag.String("g", "", "valid link for github repository")
	config.GithubUsername = flag.String("u", "", "github username for ssh auth")

	githubToken := os.Getenv("GH_TOKEN")
	config.GithubToken = &githubToken
	flag.StringVar(config.GithubToken, "t", *config.GithubToken, "github token for ssh auth")

	config.LightCheck = false
	config.NoSummary = false
	config.NoReadme = false
	config.NoPackageSummary = false
	config.NoBackupRootReadme = false
	config.WithFileSummary = false
	config.WithPackageReadmeBackup = false

	flag.BoolFunc("c",
		"do not check project root folder specific files such as go.mod or package.json",
		func(_ string) error {
			config.LightCheck = true
			return nil
		})
	flag.BoolFunc("s",
		"do not create SUMMARY.md and README.md, just print the file summaries",
		func(_ string) error {
			config.NoSummary = true
			return nil
		})
	flag.BoolFunc("r",
		"do not create README.md",
		func(_ string) error {
			config.NoReadme = true
			return nil
		})
	flag.BoolFunc("p",
		"do not create README.md for every package in the project",
		func(_ string) error {
			config.NoPackageSummary = true
			return nil
		})
	flag.BoolFunc("br",
		"overwrite README.md for the root project directory instead of README_GENERATED.md creation",
		func(_ string) error {
			config.NoBackupRootReadme = true
			return nil
		})
	flag.BoolFunc("f",
		"Save individual file summary intermediate result to the FILES.md",
		func(_ string) error {
			config.WithFileSummary = true
			return nil
		})
	flag.BoolFunc("bp",
		"create README_GENERATED.md if README.md exists in the package directory instead of overwriting",
		func(_ string) error {
			config.WithPackageReadmeBackup = true
			return nil
		})

	flag.Parse()

	return &config, nil
}
