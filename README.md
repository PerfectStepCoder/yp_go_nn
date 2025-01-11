# Комплекс для запуска нейронных сетей

## Статический анализ кода
### Проверка стиля 
> gofmt -s -w .
-s simplifies the code
-w writes results directly
> gofmt -l -s .
> goimports -l .
> /Users/dmitrii/go/bin/goimports -l .
### Проверка стуктурных тегов
> go vet -structtag

## API
### REST HTTP
#### Создание документации Swagger
> go install github.com/swaggo/swag/cmd/swag@latest
> export PATH="$PATH:$(go env GOPATH)/bin"  // чтобы видна была утилита
> swag init -g cmd/main.go -o api/irest/docs // создаем документацию

### GRPC
#### Кодогенерация
##### Go доступность в консоле
> export PATH="$PATH:$(go env GOPATH)/bin"
> cd ../internal/proto
> protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative models.proto server.proto
##### Python
Установка
> pip install --upgrade grpcio grpcio-tools
> python -m grpc_tools.protoc --version
Генерация
> python -m grpc_tools.protoc -I . --python_out=./gen_py --pyi_out=./gen_py --grpc_python_out=./gen_py ./models.proto ./server.proto
-I .: Указывает путь к директории, где находятся файлы .proto. Текущая директория обозначена как .
--python_out=.: Указывает путь для генерации файлов Python, содержащих структуры данных из .proto
--grpc_python_out=.: Указывает путь для генерации файлов Python, содержащих определения gRPC-сервисов 
example.proto: Имя файла .proto, который вы хотите обработать.
#### Запуск сервиса
> cd ../cmd/
> go run main.go -h 0.0.0.0 -p 3001 -m grpc
#### Запуск утилиты создание отчетов
> cd ../cmd/test_reporter
> go run main.go reporter.go settings.go

# Нейронные сети
### Установка ultralytics
> conda install -c conda-forge ultralytics

# Запуск проекта
### Запуск с флагами
> go run services/go_nn/src/cmd/main.go -h 0.0.0.0 -p 8080 -m http
