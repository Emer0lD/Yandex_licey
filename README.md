#Веб-сервис для вычисления арифметических выражений

## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Структура проекта

- `cmd/` — точка входа приложения.
- `configs/` — конфигурационные файлы проекта.
- `internal/` — внутренняя логика и модули приложения.
- `pkg/` — вспомогательные пакеты и утилиты.
- `build/` — скрипты для сборки и тестирования.
- `test/` — интеграционные и модульные тесты.
- `docs/` — документация проекта.

## Запуск сервиса

1. Установите [Go](https://go.dev/dl/).
2. Склонируйте проект с GitHub:
    ```bash
    git clone https://github.com/your-username/calc_service.git
    ```
3. Перейдите в папку проекта и запустите сервер:
    ```bash
    go run ./cmd/calc_service/...
    ```
4. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

### Альтернативный запуск
Вы можете использовать скрипты для сборки и запуска:
- **Для Linux/MacOS:**
    ```bash
    ./build/build.sh
    ```
- **Для Windows:**
    ```powershell
    .\build\build.bat
    ```

## Эндпоинты

### `POST /api/v1/calculate`

#### Описание
Эндпоинт принимает JSON с математическим выражением.

#### Пример запроса с использованием PowerShell

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
Пример успешного ответа

{
  "result": "6.000000"
}
Пример ошибки 500
Если выражение содержит некорректный символ $, сервер вернёт ошибку 500:

Пример запроса
powershell

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "1+$2"}'
Пример ответа
json

{
  "error": "Некорректное выражение"
}
Тестирование
Для запуска тестов выполните:

bash

go test ./...
Примечания
Для работы API требуется установленный Go (версии 1.18 и выше).
Все зависимости проекта управляются через go mod. Убедитесь, что в корне проекта находятся go.mod и go.sum.
