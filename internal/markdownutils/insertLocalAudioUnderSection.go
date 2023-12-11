package markdownutils

import (
	"fmt"
	"strings"
)

// InsertLocalAudioUnderSection вставляет ссылку на локальный аудиофайл в указанный раздел.
// content: исходный Markdown-контент.
// section: заголовок раздела, под которым нужно вставить аудио.
// localAudioPath: путь к локальному аудиофайлу.
// Возвращает обновленный Markdown-контент.
func InsertLocalAudioUnderSection(content, section, localAudioPath string) string {
	index := strings.Index(content, section)
	if index == -1 {
		return content // Раздел не найден, возвращаем исходный контент
	}

	insertionPoint := index + len(section)
	audioMarkdown := fmt.Sprintf("\n![[%s]]", localAudioPath)
	return content[:insertionPoint] + audioMarkdown + content[insertionPoint:]
}
