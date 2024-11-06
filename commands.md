Stop process
```
sudo lsof -i :15672
sudo lsof -i :5432
sudo lsof -i :3000
sudo lsof -i :80
sudo kill PID_number
```

Инициализация зависимостей
```
go mod init github.com/EvgeniyBudaev/tgdating-go/app
```

Сборка
```
go build -v ./cmd/
```

Удаление неиспользуемых зависимостей
```
go mod tidy -v
```

ENV Config
https://github.com/kelseyhightower/envconfig
```
go get -u github.com/kelseyhightower/envconfig
```

Логирование
https://pkg.go.dev/go.uber.org/zap
```
go get -u go.uber.org/zap
```

Валидатор
https://github.com/go-playground/validator
```
go get github.com/go-playground/validator/v10
```

Errors
```
go get -u github.com/pkg/errors
```

Подключение к БД
Драйвер для Postgres
```
go get -u github.com/lib/pq
```

PostGIS
```
pg_config --version // PostgreSQL 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
sudo apt-get update
sudo apt install postgis postgresql-14-postgis-3
sudo -u postgres psql -c "CREATE EXTENSION postgis;" tgbot
sudo systemctl restart postgresql
```

Fiber
https://github.com/gofiber/fiber
```
go get -u github.com/gofiber/fiber/v2
```

CORS
https://github.com/gorilla/handlers
```
go get -u github.com/gorilla/handlers
```

go-webp Сжатие изображений
https://github.com/h2non/bimg
```
sudo apt-get update
sudo apt install libvips-dev
go get -u github.com/h2non/bimg
```

Миграции
https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
https://www.appsloveworld.com/go/83/golang-migrate-installation-failing-on-ubuntu-22-04-with-the-following-gpg-error
```
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
sudo sh -c 'echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list'
sudo apt-get update
sudo apt-get install -y golang-migrate
go get -u github.com/golang-migrate/migrate
```

Если ошибка E: Указаны конфликтующие значения параметра Signed-By из источника
https://packagecloud.io/golang-migrate/migrate/ubuntu/
jammy: /etc/apt/keyrings/golang-migrate_migrate-archive-keyring.gpg !=
```
cd /etc/apt/sources.list.d
ls
sudo rm migrate.list
```

Создание миграционного репозитория
```
migrate create -ext sql -dir migrations initSchema
```

Создание up sql файлов
```
migrate -path migrations -database "postgres://localhost:5432/tgbot?sslmode=disable&user=postgres&password=root" up
```

Создание down sql файлов
```
migrate -path migrations -database "postgres://localhost:5432/tgbot?sslmode=disable&user=postgres&password=root" down
```

Если ошибка Dirty database version 1. Fix and force version
```
migrate create -ext sql -dir migrations initSchema force 20240410053939
```

Create Let’s Encrypt Wildcard Certificates in NGINX
https://www.webhi.com/how-to/generate-lets-encrypt-wildcard-certificates-nginx/

Keycloak
```
./kc.sh start-dev --http-port=8181
```

Make
```
sudo apt update
sudo apt-get install build-essential
make --version
make build_server
```

S3
https://github.com/aws/aws-sdk-go
https://pkg.go.dev/github.com/aws/aws-sdk-go/service/s3#hdr-Using_the_Client
```
go get -u github.com/aws/aws-sdk-go
```

Telegram Bot API
```
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
```

Telegram Init data
https://github.com/telegram-mini-apps/init-data-golang
```
go get github.com/telegram-mini-apps/init-data-golang
```

Crypto
https://github.com/Luzifer/go-openssl
```
go get github.com/Luzifer/go-openssl/v4
```

Protocol Buffer Compiler
https://grpc.io/docs/protoc-installation/
```
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

После этого установите утилиты, которые отвечают за кодогенерацию go-файлов:
```
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf
```
Разработка gRPC-сервера
После того как были реализованы все необходимые интерфейсы, можно приступать к созданию функции main.
Она запустит gRPC-сервер.
Вот алгоритм по шагам:
1. При вызове net.Listen указать порт, который будет прослушивать сервер.
2. Создать экземпляр gRPC-сервера функцией grpc.NewServer().
3. Зарегистрировать созданный сервис ProfileService на сервере gRPC.
4. Вызвать Serve() для начала работы сервера. Он будет слушать указанный порт, пока процесс не прекратит работу.

Разработка gRPC-клиента
Соединение с сервером устанавливается при вызове функции grpc.Dial(). В первом параметре указывается адрес сервера,
далее перечисляются опциональные параметры.
Функция pb.NewProfileClient(conn) возвращает переменную интерфейсного типа UsersClient, для которого сгенерированы
методы с соответствующими запросами из proto-файла.

Вызовите утилиту protoc для генерации соответствующих go-файлов. Для этого выполните команду:
```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  contracts/proto/profiles/profile.proto
```
В --go-out запишется файл с кодом для Protobuf-сериализации.
В --go-grpc_out сохранится файл с gRPC-интерфейсами и методами.
Так как вы указали параметр paths=source_relative, сгенерированные файлы создадутся в поддиректории ./proto.
Если бы указали параметр paths=import, то сгенерированные файлы создались бы в директории,
указанной в директиве go_package, то есть ./demo/proto.