package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"wordsToFlashCards/internal/contentcreator"
	"wordsToFlashCards/internal/markdownutils"
)

const (
	voice         = "echo"
	systemMessage = "Я создаю приложение для превращения английского слова в флэш-карту. \nТ.е я буду создавать связь \nСлово  - описание слова и примеры использования слова в речи.\nДля этого я буду использовать тебя.\nНе должно быть пропусков между строками. Это очень важно. Никаких пропусков между строками. Ни одной пустой строки.\nНА любое слово или фразу отвечай как дальше\n\n Примечание: не должно быть ни одной пустой строки, никаких пропусков, тк это сломает приложение\nПример запроса:\nanger. Пример ответа: Anger\n—\n## Description of Anger:\n<A very simple explanation of what the meaning of the word>\n## Examples of Usage:\n1. <Use of a word in a phrase>\n2. <Use of a word in a phrase>\n...\n## Audio\n## Illustration"
	//contentPath  = "/Master Vault/content/"
	contentPath = ""
)

const size = 700

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in .env")
	}

	contentPath := os.Getenv("CONTENT_PATH")
	if contentPath == "" {
		log.Fatal("CONTENT_PATH is not set in .env")
	}
	//contentPath += contentPath
	log.Println("contentPath: ", contentPath)
	contentCreators := contentcreator.NewContentCreator(apiKey)

	r := gin.Default()

	r.GET("/generateFlashcard/:word", func(c *gin.Context) {
		word := c.Param("word")

		content, err := contentCreators.GenerateText(systemMessage, word)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		audioData, err := contentCreators.FetchAudio(word, voice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		audioPath := fmt.Sprintf(contentPath + "flashcards/" + "audio/" + word + ".mp3")
		log.Println("audioPath: ", audioPath)
		err = contentcreator.SaveAudioToFile(audioData, audioPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		content = markdownutils.InsertLocalAudioUnderSection(content, "## Audio", "audio/"+word+".mp3")

		imgData, err := contentCreators.FetchImageData(word)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		imagePath := fmt.Sprintf(contentPath + "flashcards/" + "images/" + word + ".png")
		err = contentcreator.SaveImageToFile(imgData, imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		content = markdownutils.InsertImageUnderSection(content, "## Illustration", "images/"+word+".png", size)

		markdownFilePath := fmt.Sprintf(contentPath + "flashcards/" + word + ".md")
		err = markdownutils.CreateMarkdownFile(markdownFilePath, content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Flashcard generated successfully"})
	})

	r.Run(":8081")
}
