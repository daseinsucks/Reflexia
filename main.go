package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	github "github.com/jackbekket/Reflexia/localgithub"
	//github "./lib"
	"github.com/joho/godotenv"
)

type Package struct {
    Name  string
    Files []string
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

    os.Mkdir("temp", os.ModePerm)
	github.Clone(repoLink,"/temp/")
    cmd := exec.Command("git", "clone", repoLink, "temp")
    err = cmd.Run()
    if err != nil {
        log.Fatalf("Error cloning repository: %s", err)
    }

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
        }
    }
    os.RemoveAll("temp")
}