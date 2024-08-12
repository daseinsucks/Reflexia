package reflexia

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

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
                pkgs[pkgName] = &Package{Name: pkgName, Files: []string{}}
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

        readme_file, err := os.Create(fmt.Sprintf("%s.md", pkg.Name))
        if err != nil  {
            fmt.Printf("Error creating .md file for package %s: %s\n", pkg.Name, err)
            return
         }
        defer readme_file.Close()

        _, err = readme_file.WriteString(fmt.Sprintf("# %s\n\n", pkg.Name))
        if err != nil  {
            fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
            return
        }

        var doc_content string
        var summary_content string
	

                        /*
                out_code,err := ParseContent(content,pkg.Path)
                if err != nil {
                    fmt.Println("Error parsing content of file for matching functions", pkg.Name, pkg.Path,err)
                }
				//ai.CreateDoc(string(content))
                _, err = readme_file.WriteString(fmt.Sprintf("## Functions\n\n%s\n\n", out_code))
                */

        /*
		for _, file := range pkg.Files {
			fmt.Printf("  File: %s\n", file)
	
			content, err := os.ReadFile(file)
	
			if err != nil {
				log.Fatalf("Error reading file: %s", err)
			}
	
			if filepath.Base(file) == "main.go" {
                doc_content = ai.CreateDoc(string(content))

                _, err = readme_file.WriteString(fmt.Sprintf("## Documentation\n\n%s\n\n", doc_content))
                if err != nil  {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
			} else {
				//summary_content = ai.TestGenerateContent(string(content))

                

                summary_content = ai.TestGenerateContent(string(content))
                _, err = readme_file.WriteString(fmt.Sprintf("## Summary\n\n%s\n\n", summary_content))
                if err != nil  {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
			}

		}
        */

        for _, file := range pkg.Files {
            fmt.Printf("  File: %s\n", file)
          
            fset := token.NewFileSet()
            f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
            if err != nil {
              log.Fatalf("Error parsing file: %s", err)
            }
          
            // This code is commenting functions
            // Inspect the AST and extract function declarations
            ast.Inspect(f, func(n ast.Node) bool {
              switch x := n.(type) {
              case *ast.FuncDecl:
                // This is a function declaration
                // Get the function declaration
                start := fset.Position(x.Pos()).Offset
                end := fset.Position(x.End()).Offset
          
                // Read the file content
                content, err := ioutil.ReadFile(file)
                if err != nil {
                  log.Fatalf("Error reading file: %s", err)
                }
          
                // Extract the function declaration
                funcDecl := string(content[start:end])
          
                // Call ai.GenerateCommentForFunction with the function declaration
                comment := ai.GenerateCommentForFunction(funcDecl)
                fmt.Println("Comment:", comment)
                _, err = readme_file.WriteString(fmt.Sprintf("## Summary\n\n%s\n\n", comment))
                if err != nil  {
                    fmt.Printf("Error writing to .md file for function %s: %s\n", pkg.Name, err)
                }
                pkgs[pkg.Name].Markdowns = append(pkgs[pkg.Name].Markdowns, comment)
              }
              return true
            })
            // end walking through functions

            content, err := os.ReadFile(file)   // walking through file in general
            if err != nil {
                fmt.Println("error reading file", err)
            }

            if filepath.Base(file) == "main.go" {
                doc_content = ai.CreateDoc(string(content))

                _, err = readme_file.WriteString(fmt.Sprintf("## Documentation\n\n%s\n\n", doc_content))
                if err != nil  {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
			} else {
				//summary_content = ai.TestGenerateContent(string(content))

                

                summary_content = ai.TestGenerateContent(string(content))
                _, err = readme_file.WriteString(fmt.Sprintf("## Summary\n\n%s\n\n", summary_content))
                if err != nil  {
                    fmt.Printf("Error writing to .md file for package %s: %s\n", pkg.Name, err)
                    return
                }
			}

          } //end walking through package files
          





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