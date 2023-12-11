package contentcreator

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"os"
)

func (cc *ContentCreator) GenerateImageUrl(prompt string) (string, error) {
	respUrl, err := cc.client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			Size:           openai.CreateImageSize512x512,
			ResponseFormat: openai.CreateImageResponseFormatB64JSON,
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

// GenerateImagePNG
// filePath - путь к папке для хранения файла. Указывать без '/'
func (cc *ContentCreator) GenerateImagePNG(prompt, filePath, fileName string) error {
	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := cc.client.CreateImage(context.Background(), reqBase64)
	if err != nil {
		return fmt.Errorf("Image creation error: %v", err)
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		return fmt.Errorf("Base64 decode error: %v", err)
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		return fmt.Errorf("PNG decode error: %v", err)
	}

	fullPath := filePath + "/" + fileName
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("File creation error: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		return fmt.Errorf("PNG encode error: %v", err)
	}

	return nil
}
