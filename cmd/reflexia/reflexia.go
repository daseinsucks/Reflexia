package main

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"

	"reflexia/pkg/project"
	"reflexia/pkg/summarizer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ghLink := flag.String("g", "", "valid link for github repository")
	ghUsername := flag.String("u", "", "github username for ssh auth")
	ghToken := os.Getenv("GH_TOKEN")
	flag.StringVar(&ghToken, "t", ghToken, "github token for ssh auth")
	lightCheck := false
	noSummary := false
	noReadme := false
	flag.BoolFunc("c",
		"do not check project root folder specific files such as go.mod or package.json",
		func(_ string) error {
			lightCheck = true
			return nil
		})
	flag.BoolFunc("s",
		"do not create SUMMARY.md and README.md, just print the file summaries",
		func(_ string) error {
			noSummary = true
			return nil
		})
	flag.BoolFunc("r",
		"do not create README.md",
		func(_ string) error {
			noReadme = true
			return nil
		})
	flag.Parse()

	dirPath, err := processWorkingDirectory(*ghLink, *ghUsername, ghToken)
	if err != nil {
		log.Fatal(err)
	}

	summarizerService := &summarizer.SummarizerService{
		HelperURL: loadEnv("HELPER_URL"),
		Model:     loadEnv("MODEL"),
		ApiToken:  loadEnv("API_TOKEN"),
		Network:   "local",
		LlmOptions: []llms.CallOption{
			llms.WithStopWords(
				[]string{
					"<end_of_summary>",
				},
			),
			llms.WithRepetitionPenalty(0.7),
		},
	}
	projectConfig := project.GetProjectConfig(dirPath, lightCheck)

	fileMap, err := summarizerService.SummarizeProjectFiles(projectConfig)
	if err != nil {
		log.Fatal(err)
	}
	if !noSummary {
		// TODO: make summary for each package (maybe directory separated)
		summaryContent, err := summarizerService.SummarizeProject(projectConfig, fileMap)
		if err != nil {
			log.Fatal(err)
		}

		summaryFile, err := os.Create(filepath.Join(dirPath, "SUMMARY.md"))
		if err != nil {
			log.Fatal(err)
		}
		defer summaryFile.Close()

		if _, err = summaryFile.WriteString(summaryContent); err != nil {
			log.Fatal(err)
		}

		if !noReadme {
			readmeContent, err := summarizerService.SummarizeReadme(projectConfig, summaryContent)
			if err != nil {
				log.Fatal(err)
			}

			readmeFilename := "README_GENERATED.md"
			if _, err := os.Stat(filepath.Join(dirPath, "README.md")); err != nil {
				if os.IsNotExist(err) {
					readmeFilename = "README.md"
				} else {
					log.Fatal(err)
				}
			}
			readmeFile, err := os.Create(filepath.Join(dirPath, readmeFilename))
			if err != nil {
				log.Fatal(err)
			}
			defer readmeFile.Close()
			if _, err = readmeFile.WriteString(readmeContent); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func processWorkingDirectory(ghLink, ghUsername, ghToken string) (string, error) {
	dirPath := loadEnv("PWD")

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

func loadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("empty environment key %s", key)
	}
	return value
}
