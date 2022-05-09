FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go mod download --only=production

RUN go get -u ./

RUN go build main.go

FROM golang:1.18-alpine as app

WORKDIR /app

COPY --from=build ./main .

ENTRYPOINT [ "./main" ]

