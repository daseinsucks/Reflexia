# Reflexia

Reflexia is a code summarizer that generates summaries of code files, packages, and the entire project. It uses a summarizer service, which can be configured using environment variables, flags, and command-line arguments.

## Configuration and Run Instructions

1. Set up the summarizer service by configuring the HELPER_URL, MODEL, and API_TOKEN environment variables.

2. Run the project using the following command-line arguments:

   - `-g, --github-link`: valid link for GitHub repository
   - `-u, --github-username`: GitHub username for SSH auth
   - `-t, --github-token`: GitHub token for SSH auth
   - `-c`: do not check project root folder specific files such as go.mod or package.json
   - `-s`: do not create SUMMARY.md and README.md, just print the file summaries
   - `-r`: do not create README.md
   - `-p`: do not create README.md for every package in the project
   - `-br`: overwrite README.md for the root project directory instead of README_GENERATED.md creation
   - `-f`: Save individual file summary intermediate result to the FILES.md
   - `-bp`: create README_GENERATED.md if README.md exists in the package directory instead of overwriting

3. The project will generate summaries of code files, packages, and the entire project, and write them to the specified output files.

