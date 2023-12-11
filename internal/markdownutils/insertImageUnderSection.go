package markdownutils

import (
	"fmt"
	"strings"
)

// InsertImageUnderSection вставляет Markdown-разметку изображения, указанного путем imagePath, в указанный раздел content под заголовком section.
// content: исходный Markdown-контент.
// section: заголовок раздела, под которым нужно вставить изображение.
// imagePath: путь к изображению, которое нужно вставить.
// Возвращает обновленный Markdown-контент.
func InsertImageUnderSection(content, section, imagePath string, size uint) string {
	index := strings.Index(content, section)
	if index == -1 {
		return content // Раздел не найден, возвращаем исходный контент
	}

	// Создаем Markdown-разметку для изображения
	imageMarkdown := fmt.Sprintf("\n![image | %d](%s)\n\n", size, imagePath)

	insertionPoint := index + len(section)
	return content[:insertionPoint] + imageMarkdown + content[insertionPoint:]
}
