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

// FetchImageData получает данные изображения через API и возвращает их в виде []byte.
// prompt: текст запроса для генерации изображения.
func (cc *ContentCreator) FetchImageData(prompt string) ([]byte, error) {
	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
		Model:          openai.CreateImageModelDallE3,
	}

	respBase64, err := cc.client.CreateImage(context.Background(), reqBase64)
	if err != nil {
		return nil, fmt.Errorf("Image creation error: %v", err)
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		return nil, fmt.Errorf("Base64 decode error: %v", err)
	}

	return imgBytes, nil
}

// SaveImageToFile сохраняет данные изображения в файл PNG.
// imgData: данные изображения в виде []byte.
// filePath: путь к файлу для сохранения изображения.
func SaveImageToFile(imgData []byte, filePath string) error {
	r := bytes.NewReader(imgData)
	img, err := png.Decode(r)
	if err != nil {
		return fmt.Errorf("PNG decode error: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("File creation error: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("PNG encode error: %v", err)
	}

	return nil
}
