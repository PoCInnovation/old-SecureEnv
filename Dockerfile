FROM golang:1.18-alpine As build

WORKDIR /app

COPY . .

RUN go mod download

RUN go get -u ./

RUN go build main.go

FROM golang:1.18-alpine as app

WORKDIR /app

COPY --from=build /app/go.mod .

COPY --from=build /app/go.sum .

COPY --from=build /app/.env .

COPY --from=build /app/main .


ENTRYPOINT [ "./main" ]