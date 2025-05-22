Варианты запуска
    Docker: 
        1) Если установлен docker-compose, то надо выполнить комаду
            docker-compose -f docker-compose.yaml up  -d --build
        2) Если установлен make, то выполнить
            make dockerup
    Ручной запуск:
        1) необходимо создать в корне файл .env и указать порт и строку подключения к бд - пример ниже.
            PORT=9999
            DSN="host=localhost user=postgres password=postgres dbname=quote-api port=8888 sslmode=disable"
        2)  Создать бд quote-api
        3)  Запустить как другое golang приложение