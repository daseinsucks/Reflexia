package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"reflexia/internal/project"
	"reflexia/internal/summarizer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dirPath := loadEnv("PWD")
	summarizerService := &summarizer.SummarizerService{
		HelperURL: loadEnv("HELPER_URL"),
		Model:     loadEnv("MODEL"),
		ApiToken:  loadEnv("API_TOKEN"),
		Network:   "local",
	}
	projectConfig := project.GetProjectConfig(dirPath)
	fileMap := map[string]string{}

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, filter := range projectConfig.FileFilter {
			if strings.HasSuffix(d.Name(), filter) {
				content, err := os.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}
				relPath, err := filepath.Rel(dirPath, path)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(relPath)
				response, err := summarizerService.CodeSummaryRequest(
					projectConfig.CodePrompt, string(content))
				fmt.Printf("\n")
				fileMap[relPath] = response
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	content := ""
	for file, summary := range fileMap {
		content += file + "\n" + summary + "\n\n"
	}
	fmt.Println("Summary:")
	summary, err := summarizerService.SummarizeRequest(projectConfig.SummaryPrompt, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")

	fmt.Println("Readme:")
	readme, err := summarizerService.SummarizeRequest(projectConfig.ReadmePrompt, summary)
	if err != nil {
		log.Fatal(err)
	}

	readmeFilename := "README_GENERATED.md"
	if _, err := os.Stat(filepath.Join(dirPath, "README.md")); err != nil {
		readmeFilename = "README.md"
	}
	readmeFile, err := os.Create(readmeFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer readmeFile.Close()

	summaryFile, err := os.Create("SUMMARY.md")
	if err != nil {
		log.Fatal(err)
	}
	defer summaryFile.Close()

	if _, err = readmeFile.WriteString(readme); err != nil {
		log.Fatal(err)
	}
	if _, err = summaryFile.WriteString(summary); err != nil {
		log.Fatal(err)
	}
}

func loadEnv(key string) string {
	if value := os.Getenv(key); value == "" {
		log.Fatalf("empty environment key %s", key)
	} else {
		return value
	}

	return ""
}
