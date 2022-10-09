# Rest API
## Запуск

1. docker-compose up
3. migrate -path migrations/ -database "postgres://postgres:qwe123@localhost:5432/restapi_gopher?sslmode=disable" up
4. make
5. ./apiserver

## Тестирование
1. docker-compose up
2. create restapi_test_gopher table
3. migrate -path migrations/ -database "postgres://postgres:qwe123@localhost:5432/restapi_test_gopher?sslmode=disable" up
4. make test