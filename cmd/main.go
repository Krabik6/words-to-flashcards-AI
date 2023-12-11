package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// Загрузка переменных из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение значения API ключа
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in .env")
	}

	client := openai.NewClient(apiKey)

	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()

		// Generate chat completion
		req := createChatCompletionRequest(word)
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		content := resp.Choices[0].Message.Content
		voice := "echo"

		// Generate and insert audio
		audioPath := fmt.Sprintf("audio/" + word + ".mp3")
		err = generateSpeechAudio(client, word, voice, audioPath)
		if err != nil {
			fmt.Printf("generating speech audio %v", err)
			continue
		}
		content = insertLocalAudioUnderSection(content, "## Audio", fmt.Sprintf(audioPath))

		// Generate and insert image
		imageURL, err := createDalleImage(client, word+" illustration")
		if err != nil {
			fmt.Printf("Image creation error: %v\n", err)
			continue
		}
		if imageURL == "" {
			fmt.Println("No image URL returned")
			continue
		}
		log.Println("Generated Image URL: ", imageURL)
		content = insertImageUnderSection(content, "## Illustration", fmt.Sprintf("![%s | %d](%s)", word, 500, imageURL))

		// Create markdown file
		err = createMarkdownFile(word+".MD", content)
		if err != nil {
			fmt.Printf("Error while creating markdown file: %v\n", err)
		}

		fmt.Printf("%s\n\n", content)
		fmt.Print("> ")
	}
}

func createMarkdownFile(fileName string, content string) error {
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	return nil
}

// createDalleImage generates an image using DALL-E and returns the URL.
func createDalleImage(client *openai.Client, prompt string) (string, error) {
	respUrl, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			Size:           openai.CreateImageSize512x512,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
			Style:          openai.CreateImageStyleNatural,
		},
	)
	if err != nil {
		return "", err
	}

	if len(respUrl.Data) == 0 {
		return "", fmt.Errorf("no image data returned")
	}

	return respUrl.Data[0].URL, nil
}

func createChatCompletionRequest(word string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Я создаю приложение для превращения английского слова в флэш-карту. \nТ.е я буду создавать связь \nСлово  - описание слова и примеры использования слова в речи.\nДля этого я буду использовать тебя.\nНе должно быть пропусков между строками. Это очень важно. Никаких пропусков между строками. Ни одной пустой строки.\nНА любое слово или фразу отвечай как дальше\n\n Примечание: не должно быть ни одной пустой строки, никаких пропусков, тк это сломает приложение\nПример запроса:\nanger. Пример ответа: Anger\n—\n## Description of Anger:\n<A very simple explanation of what the meaning of the word>\n## Examples of Usage:\n1. <Use of a word in a phrase>\n2. <Use of a word in a phrase>\n...\n## Audio\n## Illustration"},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: word,
			},
		},
	}
}

func insertImageUnderSection(content, section, imageMarkdown string) string {
	index := strings.Index(content, section)
	if index == -1 {
		return content // Section not found, return original content
	}

	insertionPoint := index + len(section)
	// Добавляем перевод строки перед и после imageMarkdown
	return content[:insertionPoint] + "\n" + imageMarkdown + "\n\n" + content[insertionPoint:]
}

func generateSpeechAudio(client *openai.Client, text, voice, fileName string) error {
	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1, // Use the default speech model
		Input:          text,
		Voice:          openai.SpeechVoice(voice), // Choose a voice
		ResponseFormat: openai.SpeechResponseFormatMp3,
		Speed:          1.0, // Normal speed
	}

	ctx := context.Background()
	response, err := client.CreateSpeech(ctx, request)
	if err != nil {
		return err
	}
	defer response.Close()

	// Create a file to save the audio
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the audio data to the file
	_, err = io.Copy(file, response)
	return err
}

func insertLocalAudioUnderSection(content, section, localAudioPath string) string {
	index := strings.Index(content, section)
	if index == -1 {
		return content // Section not found, return original content
	}

	insertionPoint := index + len(section)
	// Добавляем перевод строки перед и после ссылки на локальный аудиофайл
	audioMarkdown := fmt.Sprintf("\n![[%s]]", localAudioPath)
	return content[:insertionPoint] + audioMarkdown + content[insertionPoint:]
}
