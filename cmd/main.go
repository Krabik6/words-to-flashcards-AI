package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"wordsToFlashCards/internal/contentcreator"
	"wordsToFlashCards/internal/markdownutils"
)

const (
	voice         = "echo"
	systemMessage = "Я создаю приложение для превращения английского слова в флэш-карту. \nТ.е я буду создавать связь \nСлово  - описание слова и примеры использования слова в речи.\nДля этого я буду использовать тебя.\nНе должно быть пропусков между строками. Это очень важно. Никаких пропусков между строками. Ни одной пустой строки.\nНА любое слово или фразу отвечай как дальше\n\n Примечание: не должно быть ни одной пустой строки, никаких пропусков, тк это сломает приложение\nПример запроса:\nanger. Пример ответа: Anger\n—\n## Description of Anger:\n<A very simple explanation of what the meaning of the word>\n## Examples of Usage:\n1. <Use of a word in a phrase>\n2. <Use of a word in a phrase>\n...\n## Audio\n## Illustration"
)

const size = 700

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

	contentPath := os.Getenv("CONTENT_PATH")
	if contentPath == "" {
		log.Fatal("CONTENT_PATH is not set in .env")
	}

	contentCreators := contentcreator.NewContentCreator(apiKey)

	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()

		// Generate chat completion
		content, err := contentCreators.GenerateText(systemMessage, word)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		// Generate audio
		audioData, err := contentCreators.FetchAudio(word, voice)
		if err != nil {
			log.Fatal(err)
		}

		// Save audio to file
		audioPath := fmt.Sprintf(contentPath + "flashcards/" + "audio/" + word + ".mp3")
		err = contentcreator.SaveAudioToFile(audioData, audioPath)
		if err != nil {
			log.Fatal(err)
		}

		// Insert audio under section "Audio"
		content = markdownutils.InsertLocalAudioUnderSection(content, "## Audio", "audio/"+word+".mp3")

		// Generate image
		imgData, err := contentCreators.FetchImageData(word)
		if err != nil {
			log.Fatal(err)
		}

		// Save image to file
		imagePath := fmt.Sprintf(contentPath + "flashcards/" + "images/" + word + ".png")

		err = contentcreator.SaveImageToFile(imgData, imagePath)
		if err != nil {
			log.Fatal(err)
		}

		// Insert image under section "Illustration"
		content = markdownutils.InsertImageUnderSection(content, "## Illustration", "images/"+word+".png", size)

		markdownFilePath := fmt.Sprintf(contentPath + "flashcards/" + word + ".md")
		// Create markdown file
		err = markdownutils.CreateMarkdownFile(markdownFilePath, content)
		if err != nil {
			fmt.Printf("Error while creating markdown file: %v\n", err)
		}

		fmt.Printf("%s\n\n", content)
		fmt.Print("> ")
	}
}
