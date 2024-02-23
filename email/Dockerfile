FROM golang:latest

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .

RUN go build -o /run-email-service

EXPOSE 8080

CMD ["/run-email-service"]
