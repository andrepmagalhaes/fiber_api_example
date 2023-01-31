FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . . 

RUN go mod tidy

RUN go build -o main

EXPOSE 8080 8080

CMD ["/app/main"]