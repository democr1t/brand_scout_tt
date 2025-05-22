# Варианты запуска

## Docker

1. Если установлен **docker-compose**, выполните команду:

    ```bash
    docker-compose -f docker-compose.yaml up -d --build
    ```

2. Если установлен **make**, выполните:

    ```bash
    make dockerup
    ```

## Ручной запуск

1. Необходимо создать в корне файл `.env` и указать порт и строку подключения к БД. Пример:

    ```env
    PORT=9999
    DSN="host=localhost user=postgres password=postgres dbname=quote-api port=8888 sslmode=disable"
    ```

2. Создать базу данных `quote-api`.

3. Запустить как обычное Go-приложение.
