package service

import (
	"context"
	"fmt"

	"github.com/gofor-little/env"
	openai "github.com/sashabaranov/go-openai"
)

const generateOKRSummaryTemplate = "Resuma este Key-Result baseado na metodologia OKR: %s"

func GenerateOKRSummary(jsonData string) string {
	LLM_API_KEY := env.Get("LLM_API_KEY", "")
	client := openai.NewClient(LLM_API_KEY)
	stringToSendToOpenAi := fmt.Sprintf(generateOKRSummaryTemplate, jsonData)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: stringToSendToOpenAi,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
