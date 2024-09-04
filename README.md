# Reflexia

Reflexia is a code summarizer that takes a Go project as input and generates summaries for each file, package, and the entire project. It uses a combination of environment variables, command-line arguments, and configuration files to determine the behavior and output of the summarizer.

## Configuration and Run Instructions

1. Set up environment variables:
   - HELPER_URL: helper URL
   - MODEL: model
   - API_TOKEN: API token

2. Configure the summarizer using a .toml file in the project directory. The file should contain the following fields:
   - file filters
   - project root filters
   - module match
   - stop words
   - prompts for code, summary, package, readme, and root path

3. Run the summarizer using the following command:
   - `reflexia -g <github-link> -u <github-username> -t <github-token>`

4. Optional arguments:
   - `-c`: do not check project root folder specific files
   - `-s`: do not create SUMMARY.md and README.md, just print the file summaries
   - `-r`: do not create README.md
   - `-p`: do not create README.md for every package in the project
   - `-br`: overwrite README.md for the root project directory instead of README_GENERATED.md creation
   - `-f`: Save individual file summary intermediate result to the FILES.md
   - `-bp`: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

5. The summarizer will generate summaries for each file, package, and the entire project, and write them to the appropriate README.md files.

