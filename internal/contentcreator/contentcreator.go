package contentcreator

import "github.com/sashabaranov/go-openai"

type ContentCreator struct {
	client *openai.Client
}

// NewContentCreator создает новый экземпляр ContentCreator.
func NewContentCreator(apiKey string) *ContentCreator {
	return &ContentCreator{
		client: openai.NewClient(apiKey),
	}
}
