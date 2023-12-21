FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 80

CMD ["make", "run"]