# Reflexia: Code Summarizer

Reflexia is a tool designed to generate code summaries for a given project directory. It utilizes a summarizer service to process the code and create a comprehensive summary of the project.

## Configuration

The project configuration is loaded from environment variables and a project configuration file. You can set the environment variables as needed, and the project configuration file should be placed in the same directory as the executable.

## Running Reflexia

1. Make sure you have the necessary dependencies installed.
2. Set the environment variables as required.
3. Run the Reflexia executable: `reflexia`

This will process the code in the current directory, generate summaries for each file, and create a final summary of all summaries in a file named "SUMMARY.md". Additionally, it will generate a README file named "README_GENERATED.md" or "README.md" if a README.md file already exists.

