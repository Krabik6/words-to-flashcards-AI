# Flashcarder

Это приложение предназначено для создания обучающих флэш-карт по запросу.


## Запуск приложения в Docker

1. Убедитесь, что у вас установлен Docker и Docker Compose.

2. Склонируйте репозиторий:

    ```
    git clone https://github.com/Krabik6/words-to-flashcards-AI
    cd wordsToFlashCards
    ```

3. Создайте файл `.env` в корне проекта и установите переменные окружения:

    ```
    OPENAI_API_KEY=your_openai_api_key
    CONTENT_PATH=your_content_path 
    ```
`CONTENT_PATH` можно оставить пустым или `./`

4. Соберите и запустите контейнер:

    ```
    docker-compose up --build
    ```

## Выполнение запроса на создание флэш-карты

Чтобы запросить создание флэш-карты для определенного слова, выполните GET-запрос по следующему адресу:

```
localhost:8081/generateFlashcard/:word
```

Замените `:word` на слово, для которого хотите создать флэш-карту.

### Пример запроса:

```
localhost:8081/generateFlashcard/laptop
```

### Пример curl запроса:
```bash
curl -X GET localhost:8081/generateFlashcard/laptop
```
---

Пожалуйста, замените `<repository_url>`, `your_openai_api_key` и `your_content_path` на соответствующие значения в вашем проекте.