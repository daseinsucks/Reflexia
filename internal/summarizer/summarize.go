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

func (s *SummarizerService) SummarizeRequest(content string, prompt string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
	)

	return response, err
}
