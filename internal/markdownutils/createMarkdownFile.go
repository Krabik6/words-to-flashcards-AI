package markdownutils

import (
	"fmt"
	"os"
)

// CreateMarkdownFile сохраняет переданный Markdown-контент в файл.
// fileName: путь и имя файла, куда будет сохранен контент.
// content: Markdown-контент для сохранения.
// Возвращает ошибку в случае неудачи.
func CreateMarkdownFile(fileName, content string) error {
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	return nil
}
