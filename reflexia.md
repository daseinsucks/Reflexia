# reflexia

## Summary



## Summary

This Go code is a package named "reflexia" that performs a series of tasks related to code analysis and documentation generation. It uses various libraries and packages including "fmt", "go/parser", "go/token", "log", "os", "path/filepath", "github.com/JackBekket/Reflexia/lib/ai", "github.com/JackBekket/Reflexia/lib/github", "github.com/go-git/go-git/v5", and "github.com/joho/godotenv".

The main function, `Reflexate(repo_url string)`, takes a repository URL as an argument. It first loads environment variables from a .env file and checks if the "REPO_LINK" environment variable is set. If not, it logs a fatal error.

Next, it creates a temporary directory named "temp" and clones the repository from the "REPO_LINK" environment variable into this directory.

The code then walks through the files in the "temp" directory, parses each Go file, and creates a map of packages. For each package, it creates a Markdown file and writes the package name to it.

For each file in a package, the code checks if the file is "main.go". If it is, it uses the "ai.CreateDoc" function to generate documentation for the file and writes this to the Markdown file. If the file is not "main.go", it uses the "ai.TestGenerateContent" function to generate a summary of the file and writes this to the Markdown file.

Finally, the code removes the "temp" directory.

The package also defines a `Package` struct that holds the name of a package and a slice of file names.

This code seems to be part of a larger system that uses AI to generate documentation and summaries of Go code.


