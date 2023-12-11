package contentcreator

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
)

func (cc *ContentCreator) GenerateAudio(word, voice, fileName string) error {
	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          word,
		Voice:          openai.SpeechVoice(voice),
		ResponseFormat: openai.SpeechResponseFormatMp3,
		Speed:          1.0,
	}

	response, err := cc.client.CreateSpeech(context.Background(), request)
	if err != nil {
		return err
	}
	defer response.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response)
	return err
}
