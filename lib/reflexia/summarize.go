package reflexia

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
    Path string
    Markdowns []string
}

func Summarize(repo_url string) {
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
    
    
    //packages
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
                pkgs[pkgName] = &Package{Name: pkgName, Files: []string{}, Markdowns: []string{}}
            }
            pkgs[pkgName].Files = append(pkgs[pkgName].Files, path)
            pkgs[pkgName].Path = path
        }
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking through files: %s", err)
    }

    for _, pkg := range pkgs {
        fmt.Printf("Package: %s\n", pkg.Name)
        pkgDir := filepath.Dir(pkg.Path)
        //readmePath := filepath.Join(pkgDir, "README.md")
        //readme_file, err := os.Create(fmt.Sprintf("./temp/%s.md", pkg.Name))
        //readme_file, err := os.Create(fmt.Sprintf("./temp/%s/%s.md", pkgDir, pkg.Name))
        readme_file, err := os.Create(fmt.Sprintf("%s/README.md", pkgDir))
        if err != nil {
            fmt.Printf("Error creating .md file for package %s: %s\n", pkg.Name, err)
            return
        }
        defer readme_file.Close()

        _, err = readme_file.WriteString(fmt.Sprintf("# %s\n\n", pkg.Name))
        if err != nil {
            fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
            return
        }

    
        // Process files and write to package .md file
        for _, file := range pkg.Files {
            fmt.Printf("  File: %s\n", file)
            content, err := os.ReadFile(file)
            if err != nil {
                fmt.Println("error reading file", err)
            }
            /*
            if filepath.Base(file) == "main.go" {
                doc_content := ai.CreateDoc(string(content))
                _, err = readme_file.WriteString(fmt.Sprintf("## Documentation\n\n%s\n\n", doc_content))
                if err != nil {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
            } else {
            */
                summary_content := ai.TestGenerateContent(string(content))
                _, err = readme_file.WriteString(fmt.Sprintf("## Summary for %s\n\n%s\n\n",filepath.Base(file), summary_content))
                if err != nil {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
                pkg.Markdowns = append(pkg.Markdowns, summary_content)
            

            //pkg.Markdowns = append(pkg.Markdowns, summary_content)
        } 

        // Generate summary for the whole package
        packageSummary := ai.TestGenerateContent(strings.Join(pkg.Markdowns, "\n"))
        _, err = readme_file.WriteString(fmt.Sprintf("## Package Summary\n\n%s\n\n", packageSummary))
        if err != nil {
            fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
            return
        }
    
    }

    //os.RemoveAll("temp")
}


func ParseContent(data []byte, path string) (response []string,err error){

    //var commentary string
    matches, err := getFunctions(string(data))
    if err != nil {
        log.Fatalf("Could not get functions in %s: %v", path, err)
        return nil,err
    }
    for _, match := range matches {
        fmt.Printf("Function found in %s:\n", path)
        fmt.Println(match, "\n")
        commentary := ai.GenerateCommentForFunction(match)
        response = append(response, commentary)
    }
    
    return response, nil
}

func getFunctions(input string) ([]string, error) {
    re, err := regexp.Compile(`(?s)func\s+.+?{.*?}`)
    if err != nil {
        return nil, err
    }
    matches := re.FindAllString(input, -1)
    return matches, nil
}