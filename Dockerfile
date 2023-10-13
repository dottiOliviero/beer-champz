FROM golang:1.19

WORKDIR /go/src/app

COPY . .

RUN go build -o main cmd/main.go

CMD ["./main"]
