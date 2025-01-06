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

# Нейронные сети
### Установка ultralytics
> conda install -c conda-forge ultralytics

# Запуск проекта
### Запуск тестов
> go test -v services/go_nn/src/cmd/main_test.go
### Запуск с флагами
> go run services/go_nn/src/cmd/main.go -h 0.0.0.0 -p 8080 -m http
