FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/http

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]