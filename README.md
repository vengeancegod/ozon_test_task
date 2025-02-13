Репозиторий с тестовым заданием url-shortener. Реализованы два эндпоинта ost/sto для сокращения ссылки и получении оригинальной ссылки. 
В реализации сервиса старался придерживаться clean arch. Помимо rest api реализовал grpc.
Запуск сервиса cd ozon_bank_test -> make up-postgres/up-inmemory в зависимости от того, какое хранилище хотим использовать.
Чтобы выбрать, какой api использовать (rest/grpc), нужно поменять строку в Dockerfile (RUN CGO_ENABLED=0 go build -o main ./cmd/grpc (или http)
