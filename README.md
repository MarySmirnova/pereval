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

# Данные для проверки API

`Важно! Тестировать через swagger можно не всё, так как на данном этапе не все методы в нем отрабатывают корректно без ошибок сертификатов. Через Postman все работает ок.`

### JSON для проверки

    {
        "id": "12211",
        "beautyTitle": "пер. ",
        "title": "Пхия",
        "other_titles": "Триев",
        "connect": "", 
   
        "add_time": "2021-09-22 13:18:13",
        "user": {
            "email": "mar@ex.ru",
            "phone": "678909876",
            "name": "luka"
        }, 
   
        "coords":{
            "latitude": "45.3842",
            "longitude": "7.1525",
            "height": "1200"
        },
   
        "type": "pass", 
   
        "level":{
            "winter": "", 
            "summer": "1A",
            "autumn": "1A",
            "spring": ""
        },

        "images": {
            "sedlo": [
                        {
                        "url": "http://podvignaroda.ru/photos/natalia_160x160.jpg", 
                        "title": "PHOTO1 Podiom"
                        },
                        {
                        "url": "http://podvignaroda.ru/photos/pobeda_160x160.jpg",
                        "title": "GRRRRRR DRDRDR"
                        }                
                    ],
            "nord": [
                        {
                        "url": "http://chechnya.gov.ru/wp-content/uploads/2022/03/PSX_20220305_121506-300x200.jpg", 
                        "title": "ZLOST"
                        },
                        {
                        "url": "http://chechnya.gov.ru/wp-content/uploads/2022/02/PSX_20220216_102915-1-300x200.jpg",
                        "title":"ETO FOTO"
                        }
                    ],
            "west": null,
            "south": [
                        {
                        "url": "http://chechnya.gov.ru/wp-content/uploads/2017/12/RK_12_17-300x200.jpg",
                        "title":"truba"
                        }
                    ],
            "east": []
            }
    }