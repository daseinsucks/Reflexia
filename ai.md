# ai

## Summary


```go
package ai

import (
	"log"
	"os"

	helper "github.com/JackBekket/hellper/lib/langchain"
	"github.com/joho/godotenv"
	llm "github.com/tmc/langchaingo/llms"
)

// This package provides functions for generating function comments, summaries, and documentation for a given code.
// It uses an AI model to generate these elements based on the provided code.

// GenerateFunctionPromt generates a prompt for generating function comments.
func GenerateFunctionPromt(code string) string {
	return "Make a commentary for this function: " + code
}

// GenerateSummaryPromt generates a prompt for generating code summaries.
func GenerateSummaryPromt(code string) string {
	return "Make a summary for this code: " + code
}

// GenerateDocumentationPromt generates a prompt for generating code documentation.
func GenerateDocumentationPromt(code string) string {
	return "Make a documentation for this code: " + code
}

// GenerateContent generates content using an AI model.
func GenerateContent(baseURL, promt, modelName, apiToken, network string, options ...llm.CallOption) string {
	content, err := helper.GenerateContentInstruction(baseURL, promt, modelName, apiToken, network, options...)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// TestGenerateContent tests the GenerateContent function.
func TestGenerateContent(initialPromt string) string {
	getEnv()
	aiLink := os.Getenv("AI_URL")
	model := os.Getenv("MODEL")
	apiToken := os.Getenv("API_TOKEN")
	code := initialPromt
	request := GenerateSummaryPromt(code)
	network := "local"
	summary := GenerateContent(aiLink, request, model, apiToken, network)
	return summary
}

// GetSummaryPackage generates a summary for a package of code.
func GetSummaryPackage(initialPromt string) string {
	return TestGenerateContent(initialPromt)
}

// CreateDoc generates documentation for a piece of code.
func CreateDoc(code string) string {
	getEnv()
	aiLink := os.Getenv("AI_URL")
	model := os.Getenv("MODEL")
	apiToken := os.Getenv("API_TOKEN")
	request := GenerateDocumentationPromt(code)
	network := "local"
	documentation := GenerateContent(aiLink, request, model, apiToken, network)
	return documentation
}

// GenerateCommentForFunction generates comments for a function in a piece of code.
func GenerateCommentForFunction(code string) string {
	getEnv()
	aiLink := os.Getenv("AI_URL")
	model := os.Getenv("MODEL")
	apiToken := os.Getenv("API_TOKEN")
	network := "local"
	req := GenerateFunctionPromt(code)
	response := GenerateContent(aiLink, req, model, apiToken, network)
	return response
}

// getEnv loads environment variables from a .env file.
func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
```

This code package provides a set of functions for generating function comments, summaries, and documentation for a given code. It uses an AI model to generate these elements based on the provided code. The functions are designed to be used in a Go environment.

The package imports several packages, including "log", "os", "github.com/JackBekket/hellper/lib/langchain", "github.com/joho/godotenv", and "github.com/tmc/langchaingo/llms". The "log" package is used for logging errors, and the "os" package is used for interacting with the operating system. The "github.com/joho/godotenv" package is used for loading environment variables from a .env file. The "github.com/tmc/langchaingo/llms" package is used for interacting with an AI model.

The package provides several functions for generating prompts for generating function comments, summaries, and documentation. It also provides a function for generating content using an AI model. The "GenerateContent" function takes a base URL, a prompt, a model name, an API token, a network, and options for the AI model. The "TestGenerateContent" function tests the "GenerateContent" function. The "GetSummaryPackage" and "CreateDoc" functions use the "TestGenerateContent" function to generate summaries and documentation for a package of code, respectively. The "GenerateCommentForFunction" function generates comments for a function in a piece of code.

The package also provides a "getEnv" function for loading environment variables from a .env file. This function is called at the start of the "GenerateContent" function in the "TestGenerateContent" function.

The package is designed to be used in a Go environment.


