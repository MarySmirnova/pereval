# API для проекта Pereval

### API доступен по адресу:
`178.154.220.204:8080`

### Swagger документация:
`http://178.154.220.204:8080/swagger/index.html`

## Как опубликовать на хостинг:

1. Установить Docker и Docker-compose

2. Создать каталог для хранения данных
`mkdir -p $HOME/docker/volumes/postgres`

3. В директории проекта создать файл .env, где необходимо задать следующие переменные окружения (определить не заданные):

>> FSTR_DB_HOST=postgresql
>> FSTR_DB_PORT=5432
>> FSTR_DB_LOGIN=
>> FSTR_DB_PASS=
>> FSTR_DB_DATABASE=

>> LOG_LEVEL=INFO

>> REST_HOST=
>> REST_LISTEN=:8080
>> REST_READ_TIMEOUT=30s
>> REST_WRITE_TIMEOUT=30s

4. Запуск приложения:
`sudo docker-compose up -d`

