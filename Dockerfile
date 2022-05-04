FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go get -u ./

RUN go build main.go

ENTRYPOINT [ "./main" ]