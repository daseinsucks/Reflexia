package util

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	ignore "github.com/crackcomm/go-gitignore"
	"github.com/joho/godotenv"
)

type WalkDirIgnoredFunction func(path string, d fs.DirEntry) error

func WalkDirIgnored(dirPath string, f WalkDirIgnoredFunction) error {
	ignoreFile, err := ignore.CompileIgnoreFile(
		filepath.Join(dirPath, ".gitignore"))
	if err != nil {
		log.Printf(
			"failed to load .gitignore file for %s, %e",
			dirPath, err)
	}
	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if ignoreFile.MatchesPath(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		return f(path, d)
	})
	return err
}

func LoadEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file, %e", err)
	}
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("empty environment key %s", key)
	}
	return value
}
