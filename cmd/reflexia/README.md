# reflexia

This package provides a tool for summarizing code and generating README files for projects. It uses a helper service to analyze the code and generate summaries, and it can be configured to include or exclude specific files and directories. The package also supports cloning repositories from GitHub and writing summaries to multiple output files.

## Project Package Structure

```
cmd/reflexia/reflexia.go
```

## Code Summary

### Main Function

The main function initializes the configuration, gets the working directory, and calls various functions to summarize the code and generate the output files. It first initializes the configuration using the `initConfig()` function, which loads environment variables and parses command-line arguments. Then, it calls the `processWorkingDirectory()` function to determine the working directory based on the provided GitHub link or the first command-line argument.

Next, the main function calls the `summarizeService.SummarizeCode()` function to get a map of files and their summaries. If the `WithFileSummary` flag is set, it calls the `writeFile()` function to write this map to a file named `FILES.md`.

The main function then calls the `summarizeService.SummarizeRequest()` function multiple times to generate summaries for each package, the root project, and the individual files. It writes these summaries to the appropriate README files using the `writeFile()` function.

### initConfig Function

The `initConfig()` function loads environment variables using the `godotenv.Load()` function and initializes the configuration struct. It sets default values for the configuration fields and parses command-line arguments using the `flag.Parse()` function.

### processWorkingDirectory Function

The `processWorkingDirectory()` function gets the working directory from the `PWD` environment variable. If a GitHub link is provided, it clones the repository to the working directory. Otherwise, it uses the first command-line argument as the working directory.

### fileMapToString and fileMapToMd Functions

These functions convert the file map to a string and a markdown string, respectively.

### getReadmePath Function

This function returns the path to the README.md file.

### writeFile Function

The `writeFile()` function writes the provided content to the specified path.

### Edge Cases

The package supports various edge cases, such as cloning repositories from GitHub, handling different working directories, and generating summaries for individual files or packages. It also allows users to customize the output by specifying which files and directories to include or exclude.

### Environment Variables, Flags, and Configuration Files

The package uses the following environment variables, flags, and configuration files:

- Environment variables: `PWD`, `HELPER_URL`, `MODEL`, `API_TOKEN`
- Flags: `-g`, `--github-link`, `-u`, `--github-username`, `-t`, `--github-token`, `-c`, `-s`, `-r`, `-p`, `-br`, `-f`, `-bp`
- Configuration files: `.env`

