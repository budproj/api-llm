package service

import (
	"go-backend/model"
	"math/rand"
)

type GenerateService struct {
}

func (s GenerateService) Generate() model.Generate {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	randomNumber := rand.Intn(len(answers))

	phrase := answers[randomNumber]

	return model.Generate{
		Id:   randomNumber,
		Text: phrase,
	}
}
