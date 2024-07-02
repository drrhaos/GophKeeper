# Менеджер паролей GophKeeper

GophKeeper представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

## Сервер

Для запуска сервера создайте в корне проекта файл .env и задайте параметры.
```
PORT=8080
PORT_REST=8081
POSTGRES_DB=gophkeeper
POSTGRES_USER=postgres
POSTGRES_PASSWORD=example
PGADMIN_DEFAULT_EMAIL=postgres@eample.com
PGADMIN_DEFAULT_PASSWORD=example
DATABASE_DSN="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@db:5432/gophkeeper?sslmode=disable"

```

Далее запустите сервер коммандой.

```
docker compose -f docker-compose.yml build
docker compose -f docker-compose.yml up -d
```

Возможные параметры запуска сервера 
```
Usage of ./cmd/server/server:
  -d string
        Сетевой адрес базя данных postgres://postgres:postgres@postgres:5432/praktikum?sslmode=disable
  -g string
        Сетевой порт grpc (default "8080")
  -r string
        Сетевой порт rest (default "8081")
  -s string
        Путь до файлов статики  (default "../../swagger-ui/")
  -w string
        Путь до рабочей дирректории (default "./data")
```

## Клиент

Для сборки и запуска клиента выполните команду.

```
./tools/build_client.sh
./cmd/client/client
```

Возможные параметры запуска клиента 

```
Usage of ./cmd/client/client:
  -g string
        Сетевой адрес grpc host:port (default "127.0.0.1:8080")
  -s string
        Ключ шифрования (default "test")
  -w string
        Путь до рабочей дирректории  (default "/home/user/.gophkeeper")
```