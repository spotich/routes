FROM golang:latest

WORKDIR /app

COPY ./cmd/*.go go.mod ./

RUN go mod download

EXPOSE 80

CMD ["go", "run", "main.go"]