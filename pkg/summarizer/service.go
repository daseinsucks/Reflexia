package summarizer

import (
	helper "github.com/JackBekket/hellper/lib/langchain"
	"github.com/tmc/langchaingo/llms"
)

type SummarizerService struct {
	HelperURL  string
	Model      string
	ApiToken   string
	Network    string
	LlmOptions []llms.CallOption
}

func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
		s.LlmOptions...,
	)

	return response, err
}

func (s *SummarizerService) SummarizeRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"\n\n"+content,
		s.Model, s.ApiToken, s.Network,
		s.LlmOptions...,
	)

	return response, err
}
