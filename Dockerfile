###
#build image
###
FROM golang:1.18-alpine As build

WORKDIR /app

# copy all the project
COPY . .

#download all dependencies
RUN go mod download

# get missing repository
RUN go get -u ./

#build the go project
RUN go build main.go

###
#production image
###
FROM golang:1.18-alpine as app

WORKDIR /app

#copy go.mod from build
COPY --from=build /app/go.mod .

#copy go.sum from build
COPY --from=build /app/go.sum .

#copy .env from build
COPY --from=build /app/.env .

#copy the builded application from build
COPY --from=build /app/main .

#prefix command
ENTRYPOINT [ "./main" ]