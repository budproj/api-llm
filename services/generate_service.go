package service

import (
	"context"
	"encoding/json"
	"fmt"
	model "go-backend/services/entities"

	"github.com/gofor-little/env"
	openai "github.com/sashabaranov/go-openai"
)

type GenerateService struct {
}

func (s GenerateService) Generate(Skri model.SummarizeKeyResultInput) string {
	LLM_API_KEY := env.Get("LLM_API_KEY", "")
	client := openai.NewClient(LLM_API_KEY)
	jsonData, err := json.Marshal(Skri)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
	}
	fmt.Println(Skri)
	stringToSendToOpenAi := fmt.Sprintf("Resuma este Key-Result baseado na metodologia OKR: %s", jsonData)
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

	return resp.Choices[0].Message.Content + `<p class="mt-3 text-sm italic text-pink-500">
        Todas as informações acima foram geradas por Inteligencia Artificial e são baseada nas interações que foram realizadas até o momento
      </p>
      <div class="flex items-center mt-2 space-x-2">
        <p class="text-gray-600">Gostou do conteúdo?</p>
        <button class="px-2 py-1 text-pink-500 hover:bg-pink-500 hover:text-white"
                x-on:click="feedback === 1 ? feedback = 0 : feedback = 1">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
               x-bind:class="{ 'opacity-50': feedback === -1 }">
            <!-- Thumb up icon SVG -->
          </svg>
        </button>
        <button class="px-2 py-1 text-pink-500 hover:bg-pink-500 hover:text-white"
                x-on:click="feedback === -1 ? feedback = 0 : feedback = -1">
          <svg class="w-6 h-6 transform rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24"
               x-bind:class="{ 'opacity-50': feedback === 1 }">
            <!-- Thumb down icon SVG -->
          </svg>
        </button>
					`
}
