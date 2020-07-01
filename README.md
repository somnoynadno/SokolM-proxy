# SokolM-proxy

Простейший клиент для почасового сбора метеорологических данных через API sokolmeteo.com

## Запуск

1. Создаём файл .env с переменными окружения ниже
2. ``` $ sudo docker-compose up --build -d ```

## Environment

- db_user
- db_pass
- db_name
- db_host
- db_port
- login_payload

## БД

Работает с СУБД MySQL

Список существующих устройств пишется в файл db/devices.go
