FROM golang:1.19-bullseye as base

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=base /main .
COPY views/ /views/
COPY dist/ /dist/

CMD ["./main"]
