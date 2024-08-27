package ai

import (
	"log"
	"os"

	helper "github.com/JackBekket/hellper/lib/langchain"
	"github.com/joho/godotenv"
	llm "github.com/tmc/langchaingo/llms"
)


 var sum_instruction_pkg = "Make a summary for this code package: "
 var sum_escape = "```"
 var escape_string = "`"

 var doc_instruction = "Make a documentation for this code: "
 var doc_func_instruction = "Add commentaries to this function: "




func GenerateFunctionPromt(code string) (string){
	result := doc_func_instruction + sum_escape + "\n" + code + "\n" + sum_escape
	return result
}

func GenerateSummaryPromt(code string) (string) {
	result := sum_instruction_pkg + escape_string + "\n" + code + "\n" + escape_string
	return result
}

func GenerateDocumentationPromt(code string) (string) {
	result := doc_instruction + "\n" + sum_escape + code + sum_escape
	return result
}


func GenerateContent (base_url string, promt string, model_name string, api_token string, network string,) (string){    //options ...llms.CallOption



	   //llm_stuff
	   // try repetition penalty to workaround moments when ai stuck with same phrase over and over again
	   rp := llm.WithRepetitionPenalty(0.6)
	   //model_name
	   //var stop_words_array []string
	   //stop_words_array = append(stop_words_array, "Follow this instruction and write appropriate response:")
	   //stop_words_array = append(stop_words_array, "Response:")

	   //stop_words := llm.WithStopWords(stop_words_array)
	   //max_length := llm.WithMaxLength(7000)
	   max_tokens := llm.WithMaxTokens(7000)

	   
	   options := []llm.CallOption{
		rp,
		max_tokens,
		//stop_words,
		//max_length,
	   }
	   


	
	content, err := helper.GenerateContentInstruction(base_url, promt, model_name, api_token, network,options...)
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