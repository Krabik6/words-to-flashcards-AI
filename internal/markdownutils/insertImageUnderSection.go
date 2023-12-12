package markdownutils

import (
	"fmt"
	"log"
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
		log.Printf("Раздел %s не найден, изображение не вставлено.\n", section)
		return content // Раздел не найден, возвращаем исходный контент
	}

	// Создаем Markdown-разметку для изображения
	imageMarkdown := fmt.Sprintf("\n![image | %d](%s)\n\n", size, imagePath)

	insertionPoint := index + len(section)
	resultContent := content[:insertionPoint] + imageMarkdown + content[insertionPoint:]
	log.Printf("Изображение успешно вставлено в раздел %s.\n", section)
	return resultContent
}
