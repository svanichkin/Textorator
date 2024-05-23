package open

import (
	"context"
	"main/conf"

	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init() error {

	client = openai.NewClient(conf.Config.Openai)
	return nil

}

func DoTransform(text string) (string, error) {

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Provide only the answer itself, without adding comments. Do only what is asked and nothing more.",
			},
		},
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}

func DoGenerate(text string) (string, error) {

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Provide only the answer itself, without adding comments. Do only what is asked and nothing more.",
			},
		},
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}
