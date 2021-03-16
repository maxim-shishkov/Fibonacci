### Фибоначчи

Реализовать сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.

Сервис должен отвечать на запросы и возвращать ответ.
В ответе должны быть перечислены все числа,
последовательности Фибоначчи с порядковыми номерами от x до y.

Требования:
1. Требуется реализовать два протокола: HTTP REST и GRPC
2. Код должен быть выложен в репозиторий с возможность предоставления доступа (например github.com, bitbucker.org, gitlab.com).
   Решение предоставить ссылкой на этот репозиторий.
3. Необходимо продумать и описать в readme развертку сервиса на другом компьютер


# Setup

### Установка зависимостей

> go get github.com/gorilla/mux
> 
> go get -u google.golang.org/grpc
> 
> go get github.com/golang/protobuf/protoc-gen-go

Установить подходящий релиз Protocol Buffers
> github.com/protocolbuffers/protobuf/releases

##HTTP REST
Запуск сервеса
`go run apiREST/main.go`

Любым способом передать GET запрос с параметрами X и Y.

Например http://localhost:8080/fibonacci/?x=0&y=10

##GRPC
Запуск сервеса
`go run apiGRPC/main.go`

Передать запрос GET, например через Evans ( https://github.com/ktr0731/evans ) с параметрами X и Y.
