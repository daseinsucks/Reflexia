package summarizer

import (
	helper "github.com/JackBekket/hellper/lib/langchain"
	llm "github.com/tmc/langchaingo/llms"
)

type SummarizerService struct {
	HelperURL string
	Model     string
	ApiToken  string
	Network   string
}

var llmOptions = []llm.CallOption{
	llm.WithStopWords(
		[]string{
			"<end_of_summary>",
		},
	),
	llm.WithRepetitionPenalty(0.7),
}

func (s *SummarizerService) CodeSummaryRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"```"+content+"```",
		s.Model, s.ApiToken, s.Network,
		llmOptions...,
	)

	return response, err
}

func (s *SummarizerService) SummarizeRequest(prompt, content string) (string, error) {
	response, err := helper.GenerateContentInstruction(s.HelperURL,
		prompt+"\n\n"+content,
		s.Model, s.ApiToken, s.Network,
		llmOptions...,
	)

	return response, err
}
