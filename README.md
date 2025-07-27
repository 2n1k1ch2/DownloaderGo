### DownloaderGo
Приложение, которое позволяет создавать задачи на скачивание файлов (.pdf и .jpeg) по ссылкам из интернета
, архивировать их в .zip и скачивать результат.

### Запуск
```bash
go build -o downloader
./downloader
```
### API
## Создание задачи
POST /create-task

cURL пример:

curl -X POST http://localhost:8080/create-task

## Добавление ссылки к задаче
POST /add-link?id={task_id}&url={file_url}

Ограничения:

Поддерживаются только .pdf, .jpeg, .jpg

Не более 3 ссылок в одной задаче

cURL пример:

curl -X POST "http://localhost:8080/add-link?id=TASK_ID&url=https://example.com/sample.pdf"

## Получение статуса задачи
GET /get-task?id={task_id}

cURL пример:

curl "http://localhost:8080/get-task?id=TASK_ID"

##  Скачивание архива
GET /Download?id={task_id}

Важно: используйте флаг -OJ в curl:

-O сохраняет архив под оригинальным именем

-J говорит curl использовать имя из заголовка Content-Disposition или аналог флага в других приложениях

cURL пример:curl -OJ "http://localhost:8080/Download?id=TASK_ID"

### Ограничения
Не более 3 ссылок в задаче

Не более 3 задач одновременно

Файлы должны быть доступны без авторизации

Поддерживаемые типы: .pdf, .jpeg

