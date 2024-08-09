package reflexia

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	ai "github.com/JackBekket/Reflexia/lib/ai"
	github "github.com/JackBekket/Reflexia/lib/github"
	"github.com/go-git/go-git/v5"

	//github "./lib"
	"github.com/joho/godotenv"
)

type Package struct {
    Name  string
    Files []string
    Repo *git.Repository
}

func Reflexate(repo_url string) {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    repoLink := os.Getenv("REPO_LINK")
    if repoLink == "" {
        log.Fatal("REPO_LINK is not set in .env")
    }

    perm := os.FileMode(0755) // 0755 is the common permission for directories

    err = os.Mkdir("temp", perm)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        //return
    }
	r, err := github.Clone(repoLink,"./temp/")
    if err != nil {
        log.Fatal("error", err)
    }
    log.Println("repo :",r)
    
    /*
    cmd := exec.Command("git", "clone", repoLink, "temp")
    err = cmd.Run()
    if err != nil {
        log.Fatalf("Error cloning repository: %s", err)
    }
    */
    pkgs := make(map[string]*Package)

    err = filepath.Walk("temp", func(path string, info os.FileInfo, err error) error {
        if filepath.Ext(path) == ".go" {
            fset := token.NewFileSet()
            f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
            if err != nil {
                return err
            }
            pkgName := f.Name.Name
            if _, ok := pkgs[pkgName]; !ok {
                pkgs[pkgName] = &Package{Name: pkgName, Files: []string{}}
            }
            pkgs[pkgName].Files = append(pkgs[pkgName].Files, path)
        }
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking through files: %s", err)
    }

    for _, pkg := range pkgs {
		fmt.Printf("Package: %s\n", pkg.Name)
	
		for _, file := range pkg.Files {
			fmt.Printf("  File: %s\n", file)
	
			content, err := os.ReadFile(file)
	
			if err != nil {
				log.Fatalf("Error reading file: %s", err)
			}
	
			if filepath.Base(file) == "main.go" {
				ai.CreateDoc(string(content))
			} else {
				ai.TestGenerateContent(string(content))
				//ai.CreateDoc(string(content))
			}
		}
	}
    
    os.RemoveAll("temp")
}