package ai

import (
	"log"
	"os"

	helper "github.com/JackBekket/hellper/lib/langchain"
	"github.com/joho/godotenv"
	llm "github.com/tmc/langchaingo/llms"
)


 var sum_instruction_pkg = "Make a summary for this code: "
 var sum_escape = "```"
 var escape_string = "`"

 var doc_instruction = "Make a documentation for this code: "
 var doc_func_instruction = "Refactor this code, add commentaries about this function algorithm: "




func GenerateFunctionPromt(code string) (string){
	result := doc_func_instruction + sum_escape + code + "\n" + sum_escape
	return result
}

func GenerateSummaryPromt(code string) (string) {
	result := sum_instruction_pkg + sum_escape + code + "\n" + sum_escape
	return result
}

func GenerateDocumentationPromt(code string) (string) {
	result := doc_instruction + code + sum_escape
	return result
}


func GenerateContent (base_url string, promt string, model_name string, api_token string, network string,) (string){    //options ...llms.CallOption

	/*
	base_url := "http://example.com"
    prompt := "Hello, world"
    model_name := "test_model"
    api_token := "your_api_token"
    network := "mainnet"
	*/

	   //llm_stuff
	   options := llm.WithRepetitionPenalty(0.8)  // try repetition penalty to workaround moments when ai stuck with same phrase over and over again



	content, err := helper.GenerateContentInstruction(base_url, promt, model_name, api_token, network,options)
    if err != nil {
        log.Fatal(err)
    }

    //println(content)

	return content
}

func TestGenerateContent(initial_promt string) (string){
	aiLink := os.Getenv("AI_URL")
    if aiLink == "" {
        log.Fatal("AI_LINK is not set in .env")
    }

	model := os.Getenv("MODEL")
	if model =="" {
		log.Fatal("MODEL IS NOT SET")
	}

	api_token := os.Getenv("API_TOKEN")
	if api_token == "" {
		log.Fatal("API_TOKEN is not set")
	}

	code := initial_promt
	request := GenerateSummaryPromt(code)

	network := "local"

	summary := GenerateContent(aiLink,request,model,api_token,network)


	return summary
}

func GetSummaryPackage(initial_promt string) (string) {
	res := TestGenerateContent(initial_promt)
	return res
}


func CreateDoc(code string) (string) {
	aiLink := os.Getenv("AI_URL")
    if aiLink == "" {
        log.Fatal("AI_LINK is not set in .env")
    }

	model := os.Getenv("MODEL")
	if model =="" {
		log.Fatal("MODEL IS NOT SET")
	}

	api_token := os.Getenv("API_TOKEN")
	if api_token == "" {
		log.Fatal("API_TOKEN is not set")
	}

	request := GenerateDocumentationPromt(code)

	network := "local"

	documentation := GenerateContent(aiLink,request,model,api_token,network)


	return documentation
}

func GenerateCommentForFunction(code string) (string) {
	aiLink := os.Getenv("AI_URL")
    if aiLink == "" {
        log.Fatal("AI_LINK is not set in .env")
    }

	model := os.Getenv("MODEL")
	if model =="" {
		log.Fatal("MODEL IS NOT SET")
	}

	api_token := os.Getenv("API_TOKEN")
	if api_token == "" {
		log.Fatal("API_TOKEN is not set")
	}
	network := "local"

	req := GenerateFunctionPromt(code)
	response := GenerateContent(aiLink,req,model,api_token,network)
	return response
}



func getEnv() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}