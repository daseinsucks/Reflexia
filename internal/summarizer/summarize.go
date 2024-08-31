package summarizer

import (
	helper "github.com/JackBekket/hellper/lib/langchain"
)

type SummarizerService struct {
	HelperURL string
	Model     string
	ApiToken  string
	Network   string
}

func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
	)

	return response, err
}

func (s *SummarizerService) SummarizeRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"\n\n"+content,
		s.Model, s.ApiToken, s.Network,
	)

	return response, err
}
