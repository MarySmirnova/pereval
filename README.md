# API для проекта Pereval

### API доступен по адресу:
`178.154.220.204:8080`

### Swagger документация:
http://178.154.220.204:8080/swagger/index.html

## Как опубликовать на хостинг:

1. Установить Docker и Docker-compose

2. Создать каталог для хранения данных: <br>
`mkdir -p $HOME/docker/volumes/postgres`

3. В директории проекта создать файл `.env`, где необходимо задать следующие переменные окружения (определить не заданные):

>> FSTR_DB_HOST=postgresql <br>
>> FSTR_DB_PORT=5432 <br>
>> FSTR_DB_LOGIN= <br>
>> FSTR_DB_PASS= <br>
>> FSTR_DB_DATABASE= <br>

>> LOG_LEVEL=INFO <br>

>> REST_HOST= <br>
>> REST_LISTEN=:8080 <br>
>> REST_READ_TIMEOUT=30s <br>
>> REST_WRITE_TIMEOUT=30s <br>

4. Запуск приложения: <br>
`sudo docker-compose up -d`

