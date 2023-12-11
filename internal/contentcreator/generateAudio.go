package contentcreator

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
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

	response, err := cc.client.CreateSpeech(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func SaveAudioToFile(audioData io.ReadCloser, fileName string) error {
	defer audioData.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, audioData)
	return err
}
