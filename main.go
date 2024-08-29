package main

import (
	"log"
	"os"

	reflexia "github.com/JackBekket/Reflexia/lib/reflexia"
	"github.com/go-git/go-git/v5"

	//github "./lib"
	"github.com/joho/godotenv"
)

type Package struct {
    Name  string
    Files []string
    Repo *git.Repository
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    repoLink := os.Getenv("REPO_LINK")
    if repoLink == "" {
        log.Fatal("REPO_LINK is not set in .env")
    }

    reflexia.Summarize(repoLink)
    
}