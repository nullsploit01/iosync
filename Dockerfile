FROM golang:1.22.2

WORKDIR /go/src/app

COPY . .

RUN go build -o main ./cmd/api

CMD ["./main"]