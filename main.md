# main

## Documentation


This is a Go program that uses the Reflexia library to analyze a Git repository. It uses the go-git library for interacting with Git repositories. It also uses the godotenv library to load environment variables from a .env file.

The main function starts by loading the .env file using the `godotenv.Load()` function. If the .env file cannot be loaded, the program logs a fatal error and exits.

Next, the program checks if the REPO_LINK environment variable is set. If it is not set, the program logs a fatal error and exits.

Finally, the program calls the `reflexia.Reflexate(repoLink)` function, which is likely a function from the Reflexia library to analyze the Git repository at the URL specified by the REPO_LINK environment variable.

The `Package` struct is defined to hold information about a package, including its name, a list of file names, and a pointer to a git.Repository object.

The imports for the program include the log package for logging errors, os for interacting with the operating system, the Reflexia library for analyzing Git repositories, the go-git library for interacting with Git repositories, and the godotenv library for loading environment variables from a .env file.

Please note that the import path for the Reflexia library is not correct. It should be "github.com/JackBekket/Reflexia/lib/reflexia".


