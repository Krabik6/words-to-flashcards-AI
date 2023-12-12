package contentcreator

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
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

	log.Printf("Sending OpenAI request with word: %s\n", word)

	resp, err := cc.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println(resp)
		log.Printf("Error sending OpenAI request: %v\n", err)
		return "", err
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.Content) == 0 {
		log.Printf("No response from OpenAI for word: %s\n", word)
		return "", fmt.Errorf("no response from OpenAI")
	}

	log.Printf("Received OpenAI response: %s\n", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil
}
