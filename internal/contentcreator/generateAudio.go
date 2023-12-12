package contentcreator

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

func (cc *ContentCreator) FetchAudio(word, voice string) (io.ReadCloser, error) {
	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          word,
		Voice:          openai.SpeechVoice(voice),
		ResponseFormat: openai.SpeechResponseFormatMp3,
		Speed:          1.0,
	}

	log.Printf("Sending speech request for word: %s\n", word)

	response, err := cc.client.CreateSpeech(context.Background(), request)
	if err != nil {
		log.Printf("Speech request failed for word: %s, error: %v\n", word, err)
		return nil, err
	}

	log.Printf("Speech request successful for word: %s\n", word)

	return response, nil
}

func SaveAudioToFile(audioData io.ReadCloser, fileName string) error {
	defer audioData.Close()

	log.Printf("Saving audio to file: %s\n", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating audio file: %v\n", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, audioData)
	if err != nil {
		log.Printf("Error copying audio data to file: %v\n", err)
		return err
	}

	log.Printf("Audio saved to file: %s\n", fileName)

	return nil
}
