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
	content = AddHeaderToFront(content)
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	return nil
}

// AddHeaderToFront добавляет заданный заголовок в начало Markdown-контента.
func AddHeaderToFront(content string) string {
	header := `---
type: flashcard
tags:
  - 🃏
---
`
	return header + content
}
