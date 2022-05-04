FROM golang:1.18-alpine

WORKDIR /app

EXPOSE 8200

COPY ./srcs ./

RUN go mod download

ENTRYPOINT [ "go", "run", "main.go" ]