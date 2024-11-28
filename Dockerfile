FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o message-app

RUN chmod +x message-app

EXPOSE 9000

EXPOSE 8080

CMD ["./message-app"]