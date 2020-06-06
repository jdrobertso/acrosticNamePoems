FROM golang:1.13-alpine

WORKDIR /go/src/positive-words

COPY . .

RUN go build main.go

CMD go run main.go