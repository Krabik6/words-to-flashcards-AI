package contentcreator

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

//TODO - температура и токены

func (cc *ContentCreator) GenerateText(systemMessage, word string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemMessage},
			{Role: openai.ChatMessageRoleUser, Content: word},
		},
	}

	resp, err := cc.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.Content) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
