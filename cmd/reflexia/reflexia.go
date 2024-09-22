package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"

	util "reflexia/internal"
	"reflexia/pkg/project"
	"reflexia/pkg/summarize"
)

type Config struct {
	GithubLink              *string
	GithubUsername          *string
	GithubToken             *string
	WithConfigFile          *string
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

	workdir, err := processWorkingDirectory(
		*config.GithubLink, *config.GithubUsername, *config.GithubToken)
	if err != nil {
		log.Fatal(err)
	}

	projectConfig, err := project.GetProjectConfig(
		workdir, *config.WithConfigFile, config.LightCheck,
	)
	if err != nil {
		log.Fatal(err)
	}
	summarizeService := &summarize.SummarizerService{
		HelperURL: loadEnv("HELPER_URL"),
		Model:     loadEnv("MODEL"),
		ApiToken:  loadEnv("API_TOKEN"),
		Network:   "local",
		LlmOptions: []llms.CallOption{
			llms.WithStopWords(
				projectConfig.StopWords,
			),
			llms.WithRepetitionPenalty(0.7),
		},
	}

	// Generates summary for a file content
	fileMap, err := summarizeService.SummarizeCode(projectConfig)
	if err != nil {
		log.Fatal(err)
	}
	if config.WithFileSummary {
		if err := writeFile(
			filepath.Join(workdir, "FILES.md"),
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
							pkgDir = filepath.Join(workdir, filepath.Dir(file))
						}
					}
				}

				// Generate Summary for a package (summarizing and group file summarization by package name)
				pkgSummaryContent, err := summarizeService.SummarizeRequest(
					projectConfig.PackagePrompt,
					fileMapToString(pkgFileMap),
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

		projectStructure, err := getProjectStructure(workdir)
		if err != nil {
			log.Fatal(err)
		}

		// Generate summary for a whole project (this will be deprecated)
		summaryContent, err := summarizeService.SummarizeRequest(
			projectConfig.SummaryPrompt,
			fileMapToString(fileMap)+"\n\n"+projectStructure,
		)
		if err != nil {
			log.Fatal(err)
		}
		if err := writeFile(
			filepath.Join(workdir, "SUMMARY.md"), summaryContent); err != nil {
			log.Fatal(err)
		}

		if !config.NoReadme {
			fmt.Println("\nReadme: ")

			// Generating README based on summary content (this will be reworked)
			readmeContent, err := summarizeService.SummarizeRequest(
				projectConfig.ReadmePrompt,
				summaryContent,
			)
			if err != nil {
				log.Fatal(err)
			}

			readmeFilename := "README.md"
			if !config.NoBackupRootReadme {
				readmeFilename, err = getReadmePath(workdir)
				if err != nil {
					log.Fatal(err)
				}
			}
			if err := writeFile(
				filepath.Join(workdir, readmeFilename), readmeContent); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func loadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("empty environment key %s", key)
	}
	return value
}

func getProjectStructure(workdir string) (string, error) {
	content := "project file structure:\n"
	if err := util.WalkDirIgnored(workdir, filepath.Join(workdir, ".gitignore"), func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(workdir, path)
		if err != nil {
			return err
		}
		content += "- " + relPath + "\n"
		return nil
	}); err != nil {
		return "", err
	}
	return content, nil
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
		content += "# " + file + "\n" + summary + "\n\n"
	}
	return strings.ReplaceAll(content, "\n", "  \n")
}

func getReadmePath(workdir string) (string, error) {
	readmeFilename := "README_GENERATED.md"
	if _, err := os.Stat(filepath.Join(workdir, "README.md")); err != nil {
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

func processWorkingDirectory(githubLink, githubUsername, githubToken string) (string, error) {
	workdir := loadEnv("PWD")

	if githubLink != "" {
		u, err := url.ParseRequestURI(githubLink)
		if err != nil {
			return "", err
		}

		sPath := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
		if len(sPath) != 2 {
			return "", errors.New("github repository url does not have two path elements")
		}

		tempDirEl := []string{workdir, "temp"}
		tempDirEl = append(tempDirEl, sPath...)
		tempDir := filepath.Join(tempDirEl...)

		workdir = tempDir

		if _, err := os.Stat(workdir); err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(workdir, os.FileMode(0755)); err != nil {
					return "", err
				}

				cloneOptions := git.CloneOptions{
					URL:               githubLink,
					RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
					Depth:             1,
				}
				if githubUsername != "" && githubToken != "" {
					cloneOptions.Auth = &http.BasicAuth{
						Username: githubUsername,
						Password: githubToken,
					}
				}

				if _, err := git.PlainClone(workdir, false, &cloneOptions); err != nil {
					if err := os.RemoveAll(workdir); err != nil {
						return "", err
					}
					return "", err
				}

			} else {
				return "", err
			}
		}
	} else if len(flag.Args()) > 0 {
		workdir = flag.Arg(0)
		if _, err := os.Stat(workdir); err != nil {
			return "", err
		}
	}

	return workdir, nil
}

func initConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := Config{}

	config.GithubLink = flag.String("g", "", "valid link for github repository")
	config.GithubUsername = flag.String("u", "", "github username for ssh auth")

	githubToken := os.Getenv("GH_TOKEN")
	config.GithubToken = &githubToken
	flag.StringVar(config.GithubToken, "t", *config.GithubToken, "github token for ssh auth")

	config.WithConfigFile = flag.String("l", "", "config filename in project_config to use")

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
